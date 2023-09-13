package web

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/navythenerd/lionrouter"
	"github.com/navythenerd/nerdguardian/bot/discord"
	"github.com/navythenerd/nerdguardian/bot/storage"
)

type Service struct {
	config         *Config
	server         *http.Server
	router         *lionrouter.Router
	discordService *discord.Service
	storageService *storage.Service
}

func New(cfg *Config, discordService *discord.Service, storageService *storage.Service) *Service {
	s := &Service{
		config:         cfg,
		discordService: discordService,
		storageService: storageService,
	}

	s.router = lionrouter.New()

	s.server = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.Address, cfg.Port),
		Handler: s.router,
	}

	return s
}

func (s *Service) Start() {
	log.Println("Starting web service")
	s.registerHandler()

	go func() {
		log.Printf("Listening on %s:%d", s.config.Address, s.config.Port)
		log.Print(s.server.ListenAndServe())
	}()
}

func (s *Service) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()

		if err := s.server.Shutdown(ctx); err != nil {
			log.Println(err)
		}
	}()

	wg.Wait()
}

func (s *Service) Mux() *lionrouter.Router {
	return s.router
}
