package util

import (
	"fmt"
	"log"
	"os"
)

func Check(err error, msg string) {
	if err != nil {
		f, err2 := os.OpenFile("../err.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		defer f.Close()
		if err2 != nil {
			log.Fatal(err, err2)
		}
		_, err2 = f.Write([]byte(msg + ", error:" + (fmt.Sprint(err))))
		if err2 != nil {
			log.Fatal(err, err2)
		}
		log.Fatal(err)
	}
}
