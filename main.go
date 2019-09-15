package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	token := os.Getenv("DISCORD_APP_TOKEN")
	discord, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Fatal("Error creating Discord session,", err)
		return
	}

	defer discord.Close()

	err = discord.Open()
	if err != nil {
		fmt.Println(err)
	}

	log.Print("Bot initialized")

	discord.AddHandler(messageCreateEvent)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGSEGV, syscall.SIGHUP)
	<-sc
}
