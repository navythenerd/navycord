package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/navythenerd/navycord/bot/discord"
	"github.com/navythenerd/navycord/bot/storage"
)

func Command(storage *storage.Service) discord.Handler {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		handler := storage.GetApplicationCommandHandler(i.ApplicationCommandData().Name).(func(*discordgo.Session, *discordgo.InteractionCreate))
		handler(s, i)
	}
}
