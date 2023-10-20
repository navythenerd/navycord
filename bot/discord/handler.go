package discord

type Handler interface{}

func (s *Service) registerHandler() {
	s.Session().AddHandlerOnce(s.discordReady())
	s.Session().AddHandler(s.applicationCommandHandler())
	s.Session().AddHandler(agreeRulesReactionHandler(s.config.RulesMessageId, s.config.AgreeRulesEmoteReaction, s.config.VerifiedRoleId))
}
