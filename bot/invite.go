package bot

import (
	"fmt"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

func (bot *Bot) inviteHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		invite, err := bot.session.ChannelInviteCreate(bot.config.InviteChannelId, discordgo.Invite{
			MaxAge:  600,
			MaxUses: 1,
		})

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error generating invite link"))
			return
		}

		http.Redirect(w, r, fmt.Sprintf("https://discord.com/invite/%s", invite.Code), http.StatusSeeOther)
	})
}
