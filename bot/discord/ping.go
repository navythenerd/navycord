package discord

import (
	"github.com/bwmarrin/discordgo"
)

var ping *discordgo.ApplicationCommand = &discordgo.ApplicationCommand{
	Name:        "ping",
	Description: "Command for pinging the bot",
}

func pingHandler() Handler {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Pong!",
			},
		})
	}
}
