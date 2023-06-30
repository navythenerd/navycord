package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

func Invite(session *discordgo.Session, channelId string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		invite, err := session.ChannelInviteCreate(channelId, discordgo.Invite{
			MaxAge:  600,
			MaxUses: 1,
		})

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "Error generating invite link")
			return
		}

		http.Redirect(w, r, fmt.Sprintf("https://discord.com/invite/%s", invite.Code), http.StatusSeeOther)
	})
}
