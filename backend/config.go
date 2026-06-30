package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const ConfigDir = "./config"
const PanelConfigFile = "./config/panel.json"

func LoadPanelConfig() (*PanelConfig, error) {

	cfg := &PanelConfig{}

	if _, err := os.Stat(PanelConfigFile); os.IsNotExist(err) {

		cfg.AutoFailover = true
		cfg.Servers = []VLESSServer{}
		cfg.Routing = []RoutingRule{}

		if err := SavePanelConfig(cfg); err != nil {
			return nil, err
		}

		return cfg, nil
	}

	data, err := os.ReadFile(PanelConfigFile)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, cfg)

	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func SavePanelConfig(cfg *PanelConfig) error {

	if err := os.MkdirAll(ConfigDir, 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(cfg, "", "    ")

	if err != nil {
		return err
	}

	return os.WriteFile(
		filepath.Clean(PanelConfigFile),
		data,
		0644,
	)
}

func AddServer(server VLESSServer) error {

	cfg, err := LoadPanelConfig()

	if err != nil {
		return err
	}

	cfg.Servers = append(cfg.Servers, server)

	return SavePanelConfig(cfg)
}

func GetServers() ([]VLESSServer, error) {

	cfg, err := LoadPanelConfig()

	if err != nil {
		return nil, err
	}

	return cfg.Servers, nil
}
