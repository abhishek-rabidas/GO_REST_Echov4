package main

import (
	"AIDS_Trigger/Server"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	if os.Args[1] == "-debug" {
		logrus.SetLevel(logrus.DebugLevel)
	}

	//Database
	serverConfig := Server.Init(Server.InitializeDB())

	e := echo.New()

	//Routes
	e.POST("/alert", func(c echo.Context) error {
		return serverConfig.Alert(c)
	})

	logrus.Fatal(e.Start(":55555"))

}
