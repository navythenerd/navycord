package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/navythenerd/nerdguardian/bot/discord"
)

func Ping() discord.Handler {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Pong!",
			},
		})
	}
}
