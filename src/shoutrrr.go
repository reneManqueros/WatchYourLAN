package main

import (
	"github.com/containrrr/shoutrrr"
	"log"
	"watchyourlan/models"
)

func shoutr_notify(message string) {
	if models.AppConfig.ShoutUrl != "" {
		err := shoutrrr.Send(models.AppConfig.ShoutUrl, message)
		if err != nil {
			log.Println("ERROR: Notification failed (shoutrrr):", err)
		}
	}
}
