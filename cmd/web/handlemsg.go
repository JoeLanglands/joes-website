package main

import (
	"fmt"

	"github.com/JoeLanglands/joes-website/internal/config"
)

func listenForMessages(cfg *config.SiteConfig) {
	go func() {
		for {
			msg := <-cfg.Msg
			handleMsg(msg)
		}
	}()
}

func handleMsg(msg []byte) {
	fmt.Println(string(msg))
}
