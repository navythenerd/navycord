package handler

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func AgreeRulesReaction(rulesMessageId string, verfiedRoleId string) interface{} {
	return func(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
		if m.MessageID == rulesMessageId {
			err := s.GuildMemberRoleAdd(m.GuildID, m.UserID, verfiedRoleId)

			log.Print(m.GuildID, m.UserID)

			if err != nil {
				log.Print(err)
			}
		}
	}
}
