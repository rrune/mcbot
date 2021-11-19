package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/go-yaml/yaml"
	"github.com/rrune/mcbot/discord"
	"github.com/rrune/mcbot/models"
	. "github.com/rrune/mcbot/util"
)

func main() {
	config := models.Config{}
	f, err := os.ReadFile("./config.yml")
	Check(err, "Error while reading config.yml")
	yaml.Unmarshal(f, &config)

	dg, err := discordgo.New("Bot " + config.Token)
	Check(err, "error creating Discord session")

	h := discord.New()

	for _, handler := range h.GetHandlers() {
		dg.AddHandler(handler)
	}

	err = dg.Open()
	Check(err, "error opening connection")

	for _, command := range h.GetCommands() {
		_, err = dg.ApplicationCommandCreate(dg.State.User.ID, "", command)
		Check(err, fmt.Sprintf("Cannot create '%v' command: %v", command.Name, err))
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}
