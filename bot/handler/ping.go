package handler

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/navythenerd/navycord/bot/discord"
)

func Ping() discord.Handler {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		log.Println("Command handler for ping called")
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Pong!",
			},
		})
	}
}
