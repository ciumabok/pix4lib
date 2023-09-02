package main

import (
	"log"
	"time"

	"github.com/Pix4Devs/pix4lib/socks"
)

func main() {
	client := socks.NewSocks5Client(time.Duration(time.Second * 15))
	err := client.Connect("2.56.119.93", 5074, true, func() (string, string) {
		return "rvugmczm", "3j296ogi3b86"
	}, "108.177.126.100", 443); if err != nil {
		log.Fatal(err)
	}
}