package main

import "github.com/bwmarrin/discordgo"

var helloCommand = discordgo.ApplicationCommand{
	Name:        "hello-command",
	Description: "Hello command",
}
