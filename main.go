package main

import (
	"AIDS_Trigger/Server"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func main() {
	if os.Args[1] == "-debug" {
		logrus.SetLevel(logrus.DebugLevel)
	}

	//Database
	serverConfig := Server.Init(Server.InitializeDB())

	e := echo.New()

	ticker := time.NewTicker(5 * time.Second)

	// Create a channel to receive the tick events
	tickChan := ticker.C

	go func() {
		for {
			<-tickChan
			serverConfig.ClearBuffer()
		}
	}()

	//Routes
	e.POST("/alert", func(c echo.Context) error {
		return serverConfig.Alert(c)
	})

	e.GET("/", func(c echo.Context) error {
		return serverConfig.CheckActivity(c)
	})

	logrus.Fatal(e.Start(":55555"))

}
