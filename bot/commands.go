package bot

import (
	"log"

	"github.com/navythenerd/navycord/bot/commands"
	"github.com/navythenerd/navycord/bot/handler"
)

func (b *Bot) registerCommands() {
	log.Println("Registering commands...")
	cmd, err := b.discordService.Session().ApplicationCommandCreate(b.discordService.Session().State.User.ID, "", commands.Ping)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(cmd.Name + " registered")

	b.storageService.AddApplicationCommand(cmd, handler.Ping(), "")
}

func (b *Bot) unregisterCommands() {
	commandStorage := b.storageService.GetApplicationCommands()

	for _, cmd := range commandStorage {
		if cmd.RegisteredGuildId != "" {
			err := b.discordService.Session().ApplicationCommandDelete(b.discordService.Session().State.User.ID, cmd.RegisteredGuildId, cmd.ID)

			if err != nil {
				log.Printf("Error while deleting command: %s", err)
			}
		}
	}
}
