package handler

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/navythenerd/navycord/bot/discord"
	"github.com/navythenerd/navycord/bot/storage"
)

func DiscordRules(storageService *storage.Service, rules string, channel string) discord.Handler {
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
			m, err = discordSession.ChannelMessageEdit(channel, rulesMessages[0].ID, string(mdRules))

			if err != nil {
				log.Println(err)
				return
			}
		}

		ruleMessage := storage.Message{
			ID:         m.ID,
			User:       m.Author.ID,
			Content:    m.Content,
			Annotation: "rules",
		}

		storageService.DB().Save(&ruleMessage)
	}
}
