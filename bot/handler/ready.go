package handler

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/navythenerd/nerdguardian/bot/discord"
)

func DiscordReady(statusChannelId string) discord.Handler {
	return func(s *discordgo.Session, r *discordgo.Ready) {
		_, err := s.ChannelMessageSend(statusChannelId, "Ready to serve!")

		if err != nil {
			log.Fatal(err)
		}
	}
}
