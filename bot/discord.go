package bot

import (
	"github.com/navythenerd/navycord/bot/handler"
)

func (bot *Bot) registerDiscordHandler() {
	bot.discord.Session().AddHandlerOnce(handler.DiscordReady("1124026687080902726"))
}
