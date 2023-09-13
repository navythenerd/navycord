package web

import (
	"fmt"
	"io"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

func (s *Service) inviteHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		invite, err := s.discordService.Session().ChannelInviteCreate(s.discordService.Config().InviteChannelId, discordgo.Invite{
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
