package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/navythenerd/nerdguardian/bot/storage"
)

type Service struct {
	session  *discordgo.Session
	storage  *storage.Service
	commands map[string]*ApplicationCommand
	config   *Config
}

func New(cfg *Config, storage *storage.Service) (*Service, error) {
	// create discord session
	connectString := fmt.Sprintf("Bot %s", cfg.Token)
	s, err := discordgo.New(connectString)

	if err != nil {
		return nil, err
	}

	// set all intents so we can do any action
	s.Identify.Intents = discordgo.IntentsAll

	d := &Service{
		session:  s,
		storage:  storage,
		commands: make(map[string]*ApplicationCommand),
		config:   cfg,
	}

	return d, err
}

func (s *Service) Connect() error {
	s.registerHandler()

	err := s.session.Open()

	if err != nil {
		return err
	}

	s.registerApplicationCommands()
	return nil
}

func (s *Service) Shutdown() {
	if s.session != nil {
		s.unregisterApplicationCommands()
		s.session.Close()
	}
}

func (s *Service) Session() *discordgo.Session {
	return s.session
}
