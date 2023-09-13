package discord

type Handler interface{}

func (s *Service) registerHandler() {
	s.Session().AddHandlerOnce(s.discordReady())
	s.Session().AddHandlerOnce(s.discordRulesHandler())
	s.Session().AddHandler(s.applicationCommandHandler())
}
