package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type ApplicationCommand struct {
	*discordgo.ApplicationCommand
	Handler           Handler
	RegisteredGuildId string
}

func (s *Service) registerApplicationCommands() {
	// ping command
	cmd, err := s.Session().ApplicationCommandCreate(s.Session().State.User.ID, "", ping)

	if err != nil {
		log.Fatal(err)
	}

	s.addApplicationCommand(cmd, pingHandler(), "")
}

func (s *Service) unregisterApplicationCommands() {
	for _, cmd := range s.commands {
		if cmd.RegisteredGuildId != "" {
			err := s.Session().ApplicationCommandDelete(s.Session().State.User.ID, cmd.RegisteredGuildId, cmd.ID)

			if err != nil {
				log.Printf("Error while deleting command: %s", err)
			}
		}
	}
}

func (s *Service) addApplicationCommand(command *discordgo.ApplicationCommand, handler Handler, guildId string) {
	s.commands[command.Name] = &ApplicationCommand{
		ApplicationCommand: command,
		Handler:            handler,
		RegisteredGuildId:  guildId,
	}
}

func (discord *Service) applicationCommandHandler() Handler {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if handler := discord.commands[i.ApplicationCommandData().Name].Handler.(func(*discordgo.Session, *discordgo.InteractionCreate)); handler != nil {
			handler(s, i)
			return
		}

		log.Printf("No handler for command: %s", i.ApplicationCommandData().Name)
	}
}
