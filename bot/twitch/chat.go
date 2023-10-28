package twitch

import (
	"fmt"
	"log"

	ttvirc "github.com/gempir/go-twitch-irc/v4"
	"github.com/navythenerd/nerdguardian/bot/storage"
	"github.com/navythenerd/nerdguardian/bot/web"
	"github.com/nicklaw5/helix/v2"
)

type ChatService struct {
	irc      *ttvirc.Client
	helix    *helix.Client
	storage  *storage.Service
	web      *web.Service
	config   *Config
	commands map[string]commandHandler
	timers   map[string]*intervalTimer
}

func NewChatService(cfg *Config, storage *storage.Service, web *web.Service) *ChatService {
	srv := &ChatService{
		config:   cfg,
		storage:  storage,
		web:      web,
		commands: make(map[string]commandHandler),
		timers:   make(map[string]*intervalTimer),
	}

	srv.irc = ttvirc.NewClient(cfg.User, fmt.Sprintf("oauth:%s", cfg.Token))

	srv.irc.OnConnect(func() {
		log.Printf("Twitch Bot joined channel: %s\n", cfg.Channel)
		srv.irc.Say(cfg.Channel, cfg.JoinMessage)
	})

	srv.registerCommands()

	srv.irc.OnPrivateMessage(srv.privateMessageHandler)

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
	for _, v := range s.timers {
		v.stop()
	}

	s.irc.Say(s.config.Channel, s.config.PartMessage)
}

func (s *ChatService) privateMessageHandler(message ttvirc.PrivateMessage) {
	if message.Message[0] == '!' {
		s.executeCommand(message)
		return
	}
}
