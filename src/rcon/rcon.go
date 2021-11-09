package rcon

import (
	"fmt"
	"log"
	"time"

	mcrcon "github.com/Kelwing/mc-rcon"
	. "github.com/rrune/mcbot/util"
)

type Rcon struct {
	conn *mcrcon.MCConn
}

func New() Rcon {
	conn := new(mcrcon.MCConn)
	err := conn.Open("mc:25575", "minecraft")
	/* if err != nil {
		log.Fatalln("Open failed", err)
	} */
	for err != nil {
		err = conn.Open("mc:25575", "minecraft")
		time.Sleep(5 * time.Second)
	}
	defer conn.Close()

	err = conn.Authenticate()
	if err != nil {
		log.Fatalln("Auth failed", err)
	}

	return Rcon{
		conn: conn,
	}
}

func (r Rcon) AddRecovery(username string) {
	_, err := r.conn.SendCommand(fmt.Sprintf("lp user %s group add recovery", username))
	Check(err, "Adding to group")
	return
}

func (r Rcon) RemoveRecovery(username string) {
	_, err := r.conn.SendCommand(fmt.Sprintf("lp user %s group remove recovery", username))
	Check(err, "Adding to group")
	return
}
