package discord

type Config struct {
	AppId           string `json:"appId"`
	Token           string `json:"token"`
	GuildId         string `json:"guildId"`
	LogChannelId    string `json:"logChannelId"`
	InviteChannelId string `json:"inviteChannelId"`
	RulesMessageId  string `json:"rulesMessageId"`
	VerifiedRoleId  string `json:"verifiedRoleId"`
	Rules           string `json:"rules"`
	RulesChannelId  string `json:"rulesChannelId"`
}
