package twitch

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/gempir/go-twitch-irc/v4"
)

func (s *ChatService) defaultCommandHandler(response string, permissionLevel uint) commandHandler {
	return func(message twitch.PrivateMessage) {
		if hasPermission(getPermissionMask(message.User.Badges), permissionLevel) {
			s.irc.Say(s.config.Channel, response)
			return
		}
	}
}

func (s *ChatService) discordInviteHandler(message twitch.PrivateMessage) {
	invite, err := s.discordService.Session().ChannelInviteCreate(s.discordService.Config().InviteChannelId, discordgo.Invite{
		MaxAge:  600,
		MaxUses: 0,
	})

	if err != nil {
		log.Println(err)
		s.irc.Say(s.config.Channel, "Error generating invite link")
		return
	}

	s.irc.Say(s.config.Channel, fmt.Sprintf("https://discord.com/invite/%s", invite.Code))
}
