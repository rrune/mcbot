package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rrune/mcbot/rcon"
)

var command = &discordgo.ApplicationCommand{
	Name:        "recovery",
	Description: "Recovery mode",
	Options: []*discordgo.ApplicationCommandOption{

		{
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Name:        "add",
			Description: "Enter recovery mode",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "username",
					Description: "Username to add recovery mode to",
					Required:    true,
				},
			},
		},
		{
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Name:        "remove",
			Description: "Exit recovery mode",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "username",
					Description: "Username to remove recovery mode from",
					Required:    true,
				},
			},
		},
	},
}

type handler struct {
	rcon     rcon.Rcon
	commands *discordgo.ApplicationCommand
}

func New() handler {
	rcon := rcon.New()

	return handler{
		rcon:     rcon,
		commands: command,
	}
}

func (h handler) GetCommands() *discordgo.ApplicationCommand {
	return h.commands
}

func (h handler) CommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	content := ""

	switch i.ApplicationCommandData().Options[0].Name {
	case "add":
		username := i.ApplicationCommandData().Options[0].Options[0].StringValue()

		h.rcon.AddRecovery(username)

		content = "Set " + username + " into recovery mode"
	case "remove":
		username := i.ApplicationCommandData().Options[0].Options[0].StringValue()

		h.rcon.AddRecovery(username)

		content = "Removed recovery mode from " + username
	}
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	})
}
