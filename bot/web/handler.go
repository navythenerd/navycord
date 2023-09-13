package web

import "log"

func (s *Service) registerHandler() {
	if s.config.InviteService {
		log.Println("Webservice invites enabled")
		s.Mux().Get(s.config.InviteHandle, s.inviteHandler())
	}
}
