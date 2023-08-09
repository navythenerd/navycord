package commands

import "github.com/bwmarrin/discordgo"

var Ping *discordgo.ApplicationCommand = &discordgo.ApplicationCommand{
	Name:        "ping",
	Description: "Command for pinging the bot",
}
