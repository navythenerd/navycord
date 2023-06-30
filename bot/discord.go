package bot

import (
	"github.com/navythenerd/navycord/bot/handler"
)

func (b *Bot) registerDiscordHandler() {
	b.discordService.Session().AddHandlerOnce(handler.DiscordReady("1124026687080902726"))
	b.discordService.Session().AddHandlerOnce(handler.DiscordRules(b.storageService, b.config.Discord.Rules, b.config.Discord.RulesChannelId))
}
