package twitch

import (
	"log"

	"github.com/gempir/go-twitch-irc/v4"
)

func (s *ChatService) defaultCommandHandler(response string, permissionLevel uint) commandHandler {
	return func(message twitch.PrivateMessage) {
		if hasPermissions(getPermissionsMask(message.User.Badges), permissionLevel) {
			s.irc.Say(s.config.Channel, response)
			return
		}
	}
}

func (s *ChatService) reloadCommandsHandler(message twitch.PrivateMessage) {
	if hasPermissions(getPermissionsMask(message.User.Badges), permissionBroadcaster) {
		log.Println("Reloading commands")
		s.commands = make(map[string]commandHandler)

		for _, timer := range s.timers {
			timer.stop()
		}
		s.timers = make(map[string]*intervalTimer)

		s.registerCommands()
	}
}
