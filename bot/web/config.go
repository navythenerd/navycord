package web

type Config struct {
	Address       string `json:"address"`
	Port          uint   `json:"port"`
	InviteService bool   `json:"inviteService"`
	InviteHandle  string `json:"inviteHandle"`
}
