package cron

import (
	"log"
	"time"

	"github.com/robfig/cron"
	"github.com/varandrew/gin-product/app/models"
)

func Setup() {
	log.Println("Starting jobs...")

	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllTag")
		models.CleanAllTag()
	})
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllArticle...")
		models.CleanAllArticle()
	})

	c.Start()

	timer := time.NewTimer(time.Second * 10)

	for {
		select {
		case <-timer.C:
			timer.Reset(time.Second * 10)
		}
	}
}
