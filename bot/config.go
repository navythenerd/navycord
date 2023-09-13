package bot

import (
	"encoding/json"
	"os"

	"github.com/navythenerd/nerdguardian/bot/discord"
	"github.com/navythenerd/nerdguardian/bot/storage"
	"github.com/navythenerd/nerdguardian/bot/twitch"
	"github.com/navythenerd/nerdguardian/bot/web"
)

type Config struct {
	Discord discord.Config `json:"discord"`
	Twitch  twitch.Config  `json:"twitch"`
	Storage storage.Config `json:"storage"`
	Web     web.Config     `json:"web"`
}

func ReadConfig(cfg *Config, file string) error {
	rawFile, err := os.ReadFile(file)

	if err != nil {
		return err
	}

	err = json.Unmarshal(rawFile, cfg)
	return err
}
