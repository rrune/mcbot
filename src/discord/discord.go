package discord

import (
	"fmt"

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
		Name:        "checkmods",
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
		"checkmods": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			data := h.modcheck.GetCache()
			fields := []*discordgo.MessageEmbedField{}

			for _, mod := range data {
				updated := ":x:"
				necessary := ""
				if mod.Updated {
					updated = ":white_check_mark:"
				}
				if !mod.OnCurse {
					updated = "_unknown_"
				}
				if mod.Necessary {
					necessary = "\n_Necessary_"
				}
				field := &discordgo.MessageEmbedField{
					Name:   mod.Name,
					Value:  fmt.Sprintf("Updated: %s\n[Link](%s)%s", updated, mod.Link, necessary),
					Inline: true,
				}
				fields = append(fields, field)
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						{
							Type:        "rich",
							Title:       "Mods",
							Description: "See which mods are updated to 1.18",
							Color:       0x5b8731,
							Fields:      fields,
						},
					},
				},
			})
		},
	}
	return handlers
}
