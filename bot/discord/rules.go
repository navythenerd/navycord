package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func agreeRulesReactionHandler(rulesMessageId string, agreeRulesEmoteReaction string, verfiedRoleId string) interface{} {
	return func(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
		if m.MessageID == rulesMessageId && m.Emoji.Name == agreeRulesEmoteReaction {
			err := s.GuildMemberRoleAdd(m.GuildID, m.UserID, verfiedRoleId)

			log.Printf("Adding verified role for user %s (with reason: accepted rules)\n", m.UserID)

			if err != nil {
				log.Print(err)
			}
		}
	}
}
