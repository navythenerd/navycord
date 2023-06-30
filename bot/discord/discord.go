package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type Service struct {
	session *discordgo.Session
	config  *Config
}

func New(cfg *Config) (*Service, error) {
	// create discord session
	connectString := fmt.Sprintf("Bot %s", cfg.Token)
	s, err := discordgo.New(connectString)

	if err != nil {
		return nil, err
	}

	// set all intents so we can do any action
	s.Identify.Intents = discordgo.IntentsAll

	d := &Service{
		session: s,
		config:  cfg,
	}

	return d, err
}

func (s *Service) Connect() error {
	return s.session.Open()
}

func (s *Service) Shutdown() {
	if s.session != nil {
		s.session.Close()
	}
}

func (s *Service) Session() *discordgo.Session {
	return s.session
}
