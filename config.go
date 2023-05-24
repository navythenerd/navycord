package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	Token   string `json:"token"`
	GuildId string `json:"guildId"`
}

func readConfig(cfg *Config, file string) error {
	rawFile, err := os.ReadFile(file)

	if err != nil {
		return err
	}

	err = json.Unmarshal(rawFile, cfg)
	return err
}
