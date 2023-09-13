package handler

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/navythenerd/nerdguardian/bot/discord"
	"github.com/navythenerd/nerdguardian/bot/storage"
)

func DiscordRules(storageService *storage.Service, rules string, channel string, agreeRulesEmoteReaction string, verifiedRole string) discord.Handler {
	return func(discordSession *discordgo.Session, r *discordgo.Ready) {
		query := storageService.DB().Model(&storage.Message{})
		query = query.Where("annotation = ?", "rules").Limit(1)
		var rulesMessages []storage.Message
		query.Find(&rulesMessages)

		mdRules, err := os.ReadFile(rules)

		if err != nil {
			log.Println(err)
			return
		}

		var m *discordgo.Message

		if len(rulesMessages) != 1 {
			m, err = discordSession.ChannelMessageSend(channel, string(mdRules))

			if err != nil {
				log.Println(err)
				return
			}

		} else {
			msg := rulesMessages[0]
			// try to edit existing message
			m, err = discordSession.ChannelMessageEdit(channel, msg.ID, string(mdRules))

			if err != nil {
				storageService.DB().Delete(msg)
				log.Println(err)

				// message doesn't exist anymore try to create a new rules message
				m, err = discordSession.ChannelMessageSend(channel, string(mdRules))

				if err != nil {
					log.Println(err)
					return
				}
			}
		}

		ruleMessage := storage.Message{
			ID:         m.ID,
			User:       m.Author.ID,
			Content:    m.Content,
			Annotation: "rules",
		}

		discordSession.AddHandler(AgreeRulesReaction(m.ID, agreeRulesEmoteReaction, verifiedRole))

		storageService.DB().Save(&ruleMessage)
	}
}

func AgreeRulesReaction(rulesMessageId string, agreeRulesEmoteReaction string, verfiedRoleId string) interface{} {
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
