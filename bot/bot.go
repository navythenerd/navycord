package bot

import (
	"log"

	"github.com/navythenerd/navycord/bot/discord"
	"github.com/navythenerd/navycord/bot/storage"
	"github.com/navythenerd/navycord/bot/web"
)

type Bot struct {
	config         *Config
	discordService *discord.Service
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
	discordService, err := discord.New(&cfg.Discord)

	if err != nil {
		return nil, err
	}

	bot.discordService = discordService
	bot.registerDiscordHandler()

	err = discordService.Connect()

	if err != nil {
		return nil, err
	}

	// setup commands
	bot.registerCommands()

	// setup local web service
	webService := web.New(&cfg.Web)
	bot.webService = webService
	bot.registerWebHandler()
	webService.Start()

	// return bot instance
	return bot, nil
}

func (b *Bot) Shutdown() {
	_, err := b.discordService.Session().ChannelMessageSend(b.config.Discord.LogChannelId, "I'm going down!")
	log.Println(err)

	b.unregisterCommands()
	b.webService.Shutdown()
	b.discordService.Shutdown()
	b.storageService.Shutdown()
}
