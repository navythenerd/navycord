package twitch

type Config struct {
	User                 string `json:"user"`
	Token                string `json:"token"`
	Channel              string `json:"channel"`
	JoinMessage          string `json:"joinMessage"`
	PartMessage          string `json:"partMessage"`
	Commands             string `json:"commands"`
	DiscordInviteMaxUses uint   `json:"discordInviteMaxUses"`
	DiscordInviteMaxAge  uint   `json:"discordInviteMaxAge"`
}
