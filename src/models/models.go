package models

type Config struct {
	Token     string `yaml:"token"`
	GuildID   string `yaml:"guildID"`
	ChannelID string `yaml:"channelID"`
	//Token     string `yaml:"tokenTest"`
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
	LastestFiles []struct {
		SortableGameVersion []struct {
			GameVersion string `json:"gameVersion"`
		} `json:"sortableGameVersion"`
	} `json:"latestFiles"`
}
