package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

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

func messageCreateEvent(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	log.Printf("%20s %20s %20s > %s\n", m.ChannelID, time.Now().Format(time.Stamp), m.Author.Username, m.Content)
}
