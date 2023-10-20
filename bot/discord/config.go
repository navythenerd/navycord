package discord

type Config struct {
	AppId                   string `json:"appId"`
	Token                   string `json:"token"`
	GuildId                 string `json:"guildId"`
	LogChannelId            string `json:"logChannelId"`
	InviteChannelId         string `json:"inviteChannelId"`
	RulesMessageId          string `json:"rulesMessageId"`
	VerifiedRoleId          string `json:"verifiedRoleId"`
	RulesChannelId          string `json:"rulesChannelId"`
	AgreeRulesEmoteReaction string `json:"agreeRulesEmoteReaction"`
}

func (s *Service) Config() *Config {
	return s.config
}
