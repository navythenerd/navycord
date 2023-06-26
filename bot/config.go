package bot

import (
	"encoding/json"
	"os"
)

type Config struct {
	AppId           string `json:"appId"`
	Token           string `json:"token"`
	GuildId         string `json:"guildId"`
	InviteChannelId string `json:"inviteChannelId"`
}

func ReadConfig(cfg *Config, file string) error {
	rawFile, err := os.ReadFile(file)

	if err != nil {
		return err
	}

	err = json.Unmarshal(rawFile, cfg)
	return err
}
