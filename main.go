package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	Token  = flag.String("t", "", "Bot Token")
	Prefix = "b."
)

func init() {
	flag.Parse()
}

func main() {
	dg, err := discordgo.New("Bot " + *Token)
	if err != nil {
		fmt.Println("[ERROR] Creating Discord session:")
		fmt.Println(err)
		return
	}

	dg.AddHandler(ready)
	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.IntentGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("[ERROR] Opening connection:")
		fmt.Println(err)
		return
	}
	defer dg.Close()

	fmt.Println("[INFO] Bot is running, press CTRL-C to exit")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop
}

func ready(s *discordgo.Session, r *discordgo.Ready) {
	fmt.Printf("[INFO] Bot %s is ready\n", r.User.Username)
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, Prefix) {
		s.ChannelMessageSend(m.ChannelID, "👁️")
	}
}
