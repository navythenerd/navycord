package bot

func (bot *Bot) initRoutes() {
	bot.router.Get("/invite", NoCache(bot.inviteHandler()))
}
