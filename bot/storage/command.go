package storage

import (
	"github.com/bwmarrin/discordgo"
	"github.com/navythenerd/navycord/bot/discord"
)

type ApplicationCommand struct {
	*discordgo.ApplicationCommand
	Handler           discord.Handler
	RegisteredGuildId string
}

func (s *Service) GetApplicationCommands() map[string]*ApplicationCommand {
	return s.commands
}

func (s *Service) GetApplicationCommandHandler(name string) discord.Handler {
	return s.commands[name].Handler
}

func (s *Service) AddApplicationCommand(command *discordgo.ApplicationCommand, handler discord.Handler, guildId string) {
	s.commands[command.Name] = &ApplicationCommand{
		ApplicationCommand: command,
		Handler:            handler,
		RegisteredGuildId:  guildId,
	}
}
