package bot

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/navythenerd/lionrouter"
)

type Bot struct {
	router  *lionrouter.Router
	session *discordgo.Session
	server  *http.Server
	config  *Config
}

func New(cfg *Config) (*Bot, error) {
	bot := &Bot{
		router: lionrouter.New(),
		config: cfg,
	}

	session, err := discordgo.New(fmt.Sprintf("Bot %s", bot.config.Token))

	if err != nil {
		return nil, err
	}

	bot.session = session
	session.Identify.Intents = discordgo.IntentsAll

	err = session.Open()

	if err != nil {
		return nil, err
	}

	bot.initRoutes()
	bot.startServer()

	return bot, nil
}

func (bot *Bot) startServer() {
	bot.server = &http.Server{
		Addr:    ":8000",
		Handler: bot.router,
	}

	go func() {
		log.Print(bot.server.ListenAndServe())
	}()
}

func (bot *Bot) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	bot.server.Shutdown(ctx)
	bot.session.Close()
}
