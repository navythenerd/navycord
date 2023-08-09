package bot

import (
	"log"

	"github.com/navythenerd/navycord/bot/handler"
)

func (bot *Bot) registerWebHandler() {
	if bot.config.Web.InviteService {
		log.Println("Webservice invites enabled")
		bot.webService.Mux().Get(bot.config.Web.InviteHandle, handler.Invite(bot.discordService.Session(), bot.config.Discord.InviteChannelId))
	}
}
