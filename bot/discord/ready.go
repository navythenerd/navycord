package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func (discordService *Service) discordReady() Handler {
	return func(s *discordgo.Session, r *discordgo.Ready) {
		_, err := s.ChannelMessageSend(discordService.config.LogChannelId, "Ready to serve!")

		if err != nil {
			log.Fatal(err)
		}
	}
}
