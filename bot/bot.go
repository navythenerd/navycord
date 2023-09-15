package bot

import (
	"log"

	"github.com/navythenerd/nerdguardian/bot/discord"
	"github.com/navythenerd/nerdguardian/bot/storage"
	"github.com/navythenerd/nerdguardian/bot/twitch"
	"github.com/navythenerd/nerdguardian/bot/web"
)

type Bot struct {
	config         *Config
	discordService *discord.Service
	ttvChatService *twitch.ChatService
	webService     *web.Service
	storageService *storage.Service
}

func New(cfg *Config) (*Bot, error) {
	bot := &Bot{
		config: cfg,
	}

	// setup storage
	storageService, err := storage.New(&cfg.Storage)

	if err != nil {
		log.Fatal(err)
	}

	bot.storageService = storageService

	// setup discord connection
	log.Println("Setting up discord service")
	discordService, err := discord.New(&cfg.Discord, storageService)

	if err != nil {
		return nil, err
	}

	bot.discordService = discordService

	err = discordService.Connect()

	if err != nil {
		return nil, err
	}

	// setup twitch chat connection
	log.Println("Setting up twitch chat service")
	bot.ttvChatService = twitch.NewChatService(&cfg.Twitch, discordService, storageService)

	bot.ttvChatService.Connect()

	// setup local web service
	log.Println("Setting up web service")
	webService := web.New(&cfg.Web, discordService, storageService)
	bot.webService = webService
	webService.Start()

	// return bot instance
	return bot, nil
}

func (b *Bot) Shutdown() {
	_, err := b.discordService.Session().ChannelMessageSend(b.config.Discord.LogChannelId, "I'm going down!")
	log.Println(err)

	b.webService.Shutdown()
	b.ttvChatService.Shutdown()
	b.discordService.Shutdown()
	b.storageService.Shutdown()
}
