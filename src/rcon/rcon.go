package rcon

import (
	"fmt"
	"log"
	"time"

	mcrcon "github.com/Kelwing/mc-rcon"
	. "github.com/rrune/mcbot/util"
)

type Rcon struct{}

func New() Rcon {
	return Rcon{}
}

func (r Rcon) GetConnection() mcrcon.MCConn {
	conn := new(mcrcon.MCConn)
	err := conn.Open("mc:25575", "minecraft")

	for err != nil {
		err = conn.Open("mc:25575", "minecraft")
		time.Sleep(5 * time.Second)
	}

	err = conn.Authenticate()
	if err != nil {
		log.Fatalln("Auth failed", err)
	}

	return *conn
}

func (r Rcon) AddRecovery(username string) {
	conn := r.GetConnection()
	defer conn.Close()
	_, err :=
		conn.SendCommand(fmt.Sprintf("lp user %s group add recover ", username))
	Check(err, "Adding to group")
}

func (r Rcon) RemoveRecovery(username string) {
	conn := r.GetConnection()
	defer conn.Close()
	_, err := conn.SendCommand(fmt.Sprintf("lp user %s group remove recover ", username))
	Check(err, "Adding to group")
}
