package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/navythenerd/goarg"
)

func main() {
	fmt.Println("NavyCord - Discord Bot")

	argParser := goarg.NewParser()

	tokenOption := goarg.NewOption("botToken", true,
		goarg.NewStringMatcher(goarg.PrefixDoubleDash, "token", true),
		goarg.NewStringMatcher(goarg.PrefixDash, "t", true),
	)

	argParser.AddOption(tokenOption)
	err := argParser.Parse(os.Args)

	if err != nil {
		fmt.Println(err)
		return
	}

	token, _ := argParser.Value("botToken")
	guildId, _ := argParser.Value("guildId")

	dcSession, err := discordgo.New(fmt.Sprintf("Bot %s", token))

	if err != nil {
		fmt.Println(err)
		return
	}

	dcSession.AddHandler(helloCommandHandler)

	err = dcSession.Open()

	if err != nil {
		fmt.Println(err)
		return
	}

	defer dcSession.Close()

	registeredHelloCommand, err := dcSession.ApplicationCommandCreate(dcSession.State.User.ID, "", &helloCommand)

	if err != nil {
		fmt.Println(err)
		return
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	err = dcSession.ApplicationCommandDelete(dcSession.State.User.ID, guildId, registeredHelloCommand.ID)

	if err != nil {
		fmt.Println(err)
		return
	}
}
