package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/navythenerd/navycord/bot"
)

func main() {
	log.Println("Starting NavyCord")
	var cfg bot.Config
	bot.ReadConfig(&cfg, "config.json")

	bot, err := bot.New(&cfg)

	if err != nil {
		log.Fatalf("Cannot create bot instance: %s", err)
		return
	}

	defer bot.Shutdown()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-stop
}
