package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/mmm-development/bulker-discord/internal/clog"
	discord "github.com/mmm-development/bulker-discord/internal/discord_frontend"
)

var (
	Token = flag.String("t", "", "Bot Token")
)

func init() {
	flag.Parse()

	clog.L.Register(os.Stdout, clog.INFO)
}

func main() {
	dg, err := discordgo.New("Bot " + *Token)
	if err != nil {
		clog.L.Fatal("Creating Discord Session:\n%v", err)
	}

	dg.AddHandler(ready)
	dg.AddHandler(interactionCreate)

	dg.Identify.Intents = discordgo.IntentGuildMessages

	err = dg.Open()
	if err != nil {
		clog.L.Fatal("Opening connection:\n%v", err)
	}
	defer dg.Close()

	clog.L.Info("Registering commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(discord.Commands))
	for i, v := range discord.Commands {
		cmd, err := dg.ApplicationCommandCreate(dg.State.User.ID, "", v)
		if err != nil {
			clog.L.Fatal("Creating '%v' command:\n%v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}
	defer func() {
		clog.L.Info("Cleaning up registered commands...")
		for _, v := range registeredCommands {
			err := dg.ApplicationCommandDelete(dg.State.User.ID, "", v.ID)
			if err != nil {
				clog.L.Fatal("Removing '%v' command:\n%v", v.Name, err)
			}
		}
	}()

	clog.L.Info("Bot is running, press CTRL-C to exit")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop
}

func ready(s *discordgo.Session, r *discordgo.Ready) {
	clog.L.Info("Bot %s is ready", r.User.Username)
}

func interactionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		if h, ok := discord.CommandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	case discordgo.InteractionMessageComponent:
		if h, ok := discord.ComponentHandlers[i.MessageComponentData().CustomID]; ok {
			h(s, i)
		}
	}
}
