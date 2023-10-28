package discord

type Handler interface{}

func (s *Service) registerHandler() {
	s.session.AddHandlerOnce(s.discordReady())
	s.session.AddHandler(s.applicationCommandHandler())
	s.session.AddHandler(agreeRulesReactionHandler(s.config.RulesMessageId, s.config.AgreeRulesEmoteReaction, s.config.VerifiedRoleId))

	if s.config.InviteService {
		s.web.Mux().Get(s.config.InviteHandle, s.inviteHandler())
	}
}
