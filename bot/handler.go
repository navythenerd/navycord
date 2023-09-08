package bot

import (
	"github.com/navythenerd/navycord/bot/handler"
)

func (b *Bot) registerDiscordHandler() {
	b.discordService.Session().AddHandlerOnce(handler.DiscordReady(b.config.Discord.LogChannelId))
	b.discordService.Session().AddHandlerOnce(handler.DiscordRules(b.storageService, b.config.Discord.Rules, b.config.Discord.RulesChannelId, b.config.Discord.AgreeRulesEmoteReaction, b.config.Discord.VerifiedRoleId))
	b.discordService.Session().AddHandler(handler.Command(b.storageService))
}
