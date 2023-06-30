package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type Discord struct {
	session *discordgo.Session
	config  *Config
}

func New(cfg *Config) (*Discord, error) {
	// create discord session
	connectString := fmt.Sprintf("Bot %s", cfg.Token)
	s, err := discordgo.New(connectString)

	if err != nil {
		return nil, err
	}

	// set all intents so we can do any action
	s.Identify.Intents = discordgo.IntentsAll

	d := &Discord{
		session: s,
		config:  cfg,
	}

	return d, err
}

func (d *Discord) Connect() error {
	return d.session.Open()
}

func (d *Discord) Shutdown() {
	if d.session != nil {
		d.session.Close()
	}
}

func (d *Discord) Session() *discordgo.Session {
	return d.session
}
