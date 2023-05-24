package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

func main() {
	fmt.Println("NavyCord - Discord Bot")
	var cfg Config
	readConfig(&cfg, "config.json")

	dcSession, err := discordgo.New(fmt.Sprintf("Bot %s", cfg.Token))

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

	err = dcSession.ApplicationCommandDelete(dcSession.State.User.ID, "", registeredHelloCommand.ID)

	if err != nil {
		fmt.Println(err)
		return
	}
}
