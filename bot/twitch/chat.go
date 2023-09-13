package twitch

import (
	"fmt"
	"log"

	ttvirc "github.com/gempir/go-twitch-irc/v4"
)

type ChatService struct {
	irc    *ttvirc.Client
	config *Config
}

func NewChatService(cfg *Config) *ChatService {
	srv := &ChatService{
		config: cfg,
	}

	srv.irc = ttvirc.NewClient(cfg.User, fmt.Sprintf("oauth:%s", cfg.Token))

	srv.irc.OnConnect(func() {
		log.Printf("Twitch Bot joined channel: %s\n", cfg.Channel)
		srv.irc.Say(cfg.Channel, "NerdGuardian joined!")
	})

	srv.irc.Join(cfg.Channel)

	return srv
}

func (s *ChatService) Connect() {
	go func() {
		err := s.irc.Connect()

		if err != nil {
			log.Fatal(err)
		}
	}()
}

func (s *ChatService) Shutdown() {
	s.irc.Say(s.config.Channel, "NerdGuardian leaving!")
}
