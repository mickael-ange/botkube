package e2e

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v3"

	"github.com/kubeshop/botkube/internal/plugin"
)

const (
	executorBinaryPrefix = "executor_"
	indexFileEndpoint    = "/botkube.yaml"
)

type PluginsConfig struct {
	BinariesDirectory string `envconfig:"default=dist"`
	Server            struct {
		Host string `envconfig:"default=http://host.k3d.internal"`
		Port int    `envconfig:"default=3000"`
	}
}

func NewPluginServer(cfg PluginsConfig) (string, func() error) {
	fs := http.FileServer(http.Dir(cfg.BinariesDirectory))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	basePath := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	http.HandleFunc(indexFileEndpoint, func(w http.ResponseWriter, _ *http.Request) {
		idx, err := buildIndex(basePath, cfg.BinariesDirectory)
		if err != nil {
			log.Printf("Cannot build index file: %s", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		out, err := yaml.Marshal(idx)
		if err != nil {
			log.Printf("Cannot marshall index file: %s", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		_, err = w.Write(out)
		if err != nil {
			log.Printf("Cannot send marshalled index file: %s", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("Listening on %s...", addr)

	server := &http.Server{
		Addr:              addr,
		ReadHeaderTimeout: 3 * time.Second,
	}

	return basePath + indexFileEndpoint, func() error {
		return server.ListenAndServe()
	}
}

func buildIndex(urlBasePath string, dir string) (plugin.Index, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return plugin.Index{}, err
	}

	entries := map[string]plugin.IndexEntry{}
	for _, entry := range files {
		if entry.IsDir() {
			continue
		}
		if !strings.HasPrefix(entry.Name(), executorBinaryPrefix) {
			continue
		}

		parts := strings.Split(entry.Name(), "_")
		if len(parts) != 4 {
			return plugin.Index{}, fmt.Errorf("path %s doesn't follow required pattern <plugin_type>_<plugin_name>_<os>_<arch>", entry.Name())
		}

		name, os, arch := parts[1], parts[2], parts[3]
		item, found := entries[name]
		if !found {
			item = plugin.IndexEntry{
				Name:        name,
				Type:        plugin.TypeExecutor,
				Description: "Executor description",
				Version:     "v0.1.0",
			}
		}
		item.URLs = append(item.URLs, plugin.IndexURL{
			URL: fmt.Sprintf("%s/static/%s", urlBasePath, entry.Name()),
			Platform: plugin.IndexURLPlatform{
				OS:   os,
				Arch: arch,
			},
		})
		entries[name] = item
	}

	var out plugin.Index
	for _, item := range entries {
		out.Entries = append(out.Entries, item)
	}
	return out, nil
}
