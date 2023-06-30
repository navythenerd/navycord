package bot

import (
	"log"

	"github.com/navythenerd/navycord/bot/discord"
	"github.com/navythenerd/navycord/bot/storage"
	"github.com/navythenerd/navycord/bot/web"
)

type Bot struct {
	config  *Config
	discord *discord.Discord
	web     *web.Service
	storage *storage.Storage
}

func New(cfg *Config) (*Bot, error) {
	bot := &Bot{
		config: cfg,
	}

	// setup storage
	storage, err := storage.New(&cfg.Storage)

	if err != nil {
		log.Fatal(err)
	}

	bot.storage = storage

	// setup discord connection
	discord, err := discord.New(&cfg.Discord)

	if err != nil {
		return nil, err
	}

	bot.discord = discord
	bot.registerDiscordHandler()

	err = discord.Connect()

	if err != nil {
		return nil, err
	}

	// setup local web service
	service := web.New(&cfg.Web)
	bot.web = service
	bot.registerWebHandler()

	service.Start()

	// return bot instance
	return bot, nil
}

func (b *Bot) Shutdown() {
	_, err := b.discord.Session().ChannelMessageSend("1124026687080902726", "I'm going down!")
	log.Println(err)

	b.web.Shutdown()
	b.discord.Shutdown()
}
