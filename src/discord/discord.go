package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rrune/mcbot/modcheck"
	"github.com/rrune/mcbot/rcon"
)

var command = []*discordgo.ApplicationCommand{
	{
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
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "reason",
						Description: "Reason for using recovery mode",
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
	},
	{
		Name:        "CheckMods",
		Description: "Check which mods are updated to 1.18",
	},
}

type handler struct {
	rcon     rcon.Rcon
	modcheck modcheck.Modcheck
	commands []*discordgo.ApplicationCommand
}

func New() handler {
	rcon := rcon.New()
	modcheck := modcheck.Init()

	return handler{
		rcon:     rcon,
		modcheck: modcheck,
		commands: command,
	}
}

func (h handler) GetCommands() []*discordgo.ApplicationCommand {
	return h.commands
}

func (h handler) GetHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	handlers := map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"recovery": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			content := ""

			switch i.ApplicationCommandData().Options[0].Name {
			case "add":
				username := i.ApplicationCommandData().Options[0].Options[0].StringValue()
				reason := i.ApplicationCommandData().Options[0].Options[1].StringValue()

				h.rcon.AddRecovery(username)

				content = "Set " + username + " into recovery mode because: " + reason
			case "remove":
				username := i.ApplicationCommandData().Options[0].Options[0].StringValue()

				h.rcon.RemoveRecovery(username)

				content = "Removed recovery mode from " + username
			}
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: content,
				},
			})
		},
		"checkMods": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if i.ApplicationCommandData()
			data := h.modcheck.Check()
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: data[0].Name,
				},
			})
		},
	}
	return handlers
}
