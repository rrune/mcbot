package models

type Config struct {
	Token string `yaml:"token"`
	//Token     string `yaml:"tokenTest"`
	GuildID   string `yaml:"guildID"`
	ChannelID string `yaml:"channelID"`
}

type Mod struct {
	Name      string `json:"name"`
	Link      string `json:"link"`
	CurseID   string `json:"curseID"`
	Necessary bool   `json:"necessary"`
	OnCurse   bool   `json:"onCurse"`
}

type ResMod struct {
	Name      string
	Link      string
	Updated   bool
	Necessary bool
	OnCurse   bool
}

type Response struct {
	GameVersionLatestFiles []GameVersionLatestFile `json:"gameVersionLatestFiles"`
}

type GameVersionLatestFile struct {
	GameVersion string `json:"gameVersion"`
}
