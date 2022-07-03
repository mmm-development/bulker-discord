package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/mmm-development/bulker-discord/appcmd"
)

var (
	Token = flag.String("t", "", "Bot Token")
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
	dg.AddHandler(interactionCreate)

	dg.Identify.Intents = discordgo.IntentGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("[ERROR] Opening connection:")
		fmt.Println(err)
		return
	}
	defer dg.Close()

	fmt.Println("[INFO] Registering commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(appcmd.Commands))
	for i, v := range appcmd.Commands {
		cmd, err := dg.ApplicationCommandCreate(dg.State.User.ID, "", v)
		if err != nil {
			fmt.Printf("[ERROR] Creating '%v' command:\n", v.Name)
			fmt.Println(err)
			return
		}
		registeredCommands[i] = cmd
	}
	defer func() {
		fmt.Println("[INFO] Cleaning up registered commands...")
		for _, v := range registeredCommands {
			err := dg.ApplicationCommandDelete(dg.State.User.ID, "", v.ID)
			if err != nil {
				fmt.Printf("[ERROR] Removing '%v' command:\n", v.Name)
				fmt.Println(err)
				return
			}
		}
	}()

	fmt.Println("[INFO] Bot is running, press CTRL-C to exit")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop
}

func ready(s *discordgo.Session, r *discordgo.Ready) {
	fmt.Printf("[INFO] Bot %s is ready\n", r.User.Username)
}

func interactionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if h, ok := appcmd.CommandHandlers[i.ApplicationCommandData().Name]; ok {
		h(s, i)
	}
}
