package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/navythenerd/navycord/bot"
)

func main() {
	fmt.Println("NavyCord - Discord Bot")
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
	log.Println("Press Ctrl+C to exit")
	<-stop
}
