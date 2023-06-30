package bot

import (
	"encoding/json"
	"os"

	"github.com/navythenerd/navycord/bot/discord"
	"github.com/navythenerd/navycord/bot/storage"
	"github.com/navythenerd/navycord/bot/web"
)

type Config struct {
	Discord discord.Config `json:"discord"`
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
