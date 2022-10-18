package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/rama4zis/go-reload-data/controllers"
	"github.com/rama4zis/go-reload-data/controllers/webcontroller"

	"github.com/gin-gonic/gin"
	"github.com/rama4zis/go-reload-data/models"
)

func main() {
	fmt.Println("Test JSON")
	min := 0
	max := 15
	status := models.Status{
		Water: rand.Intn(max-min) + min,
		Wind:  rand.Intn(max-min) + min,
	}

	controllers.WriteDataJSON(status)

	ticker := time.NewTicker(15 * time.Second)
	go tick(ticker)

	r := gin.Default()
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("./web/*")

	// web routes
	r.GET("/", webcontroller.Index)

	// api routes
	r.GET("/api/index", controllers.Index)

	r.Run(":8080")
}

func tick(ticker *time.Ticker) {
	for t := range ticker.C {
		newStatus := models.Status{
			Water: rand.Intn(15-1) + 1,
			Wind:  rand.Intn(15-1) + 1,
		}
		controllers.WriteDataJSON(newStatus)
		fmt.Println("Renew at: ", t)
	}
}
