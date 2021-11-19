package modcheck

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/rrune/mcbot/models"
	. "github.com/rrune/mcbot/util"
)

//https://curseforgeapi.docs.apiary.io/#/reference/0/get-addon-info/get-addon-info/200?mc=reference%2F0%2Fget-addon-info%2Fget-addon-info%2F200

var client = &http.Client{}

type Modcheck struct {
	modlist []models.Mod
}

func Init() Modcheck {
	modlist := []models.Mod{}
	//f, err := os.ReadFile("./modcheck/modlist.json")
	f, err := os.ReadFile("./modlist.json")
	Check(err, "Error while reading the Modlist")
	err = json.Unmarshal(f, &modlist)
	Check(err, "Error while unmarshaling Modlist")

	modcheck := Modcheck{
		modlist: modlist,
	}
	return modcheck
}

func (m Modcheck) Check() (r []models.ResMod) {
	for _, mod := range m.modlist {
		isUpdated := m.checkMod(mod.CurseID)
		Res := models.ResMod{
			Name:      mod.Name,
			Link:      mod.Link,
			Updated:   isUpdated,
			Necessary: mod.Necessary,
		}
		r = append(r, Res)
	}

	return
}

func (m Modcheck) checkMod(id string) (r bool) {
	req, err := http.NewRequest("GET", "https://addons-ecs.forgesvc.net/api/v2/addon/"+id, nil)
	Check(err, "Error while creating Request")
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	Check(err, "Error while doing request")
	defer res.Body.Close()

	respStruct := models.Response{}
	err = json.NewDecoder(res.Body).Decode(&respStruct)
	Check(err, "Error while decoding JSON")

	version := respStruct.GameVersionLatestFiles[0].GameVersion
	if version == "1.18" {
		r = true
	}
	return
}
