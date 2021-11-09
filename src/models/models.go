package models

type Config struct {
	Token     string `yaml:"token"`
	GuildID   string `yaml:"guildID"`
	ChannelID string `yaml:"channelID"`
}
