package main

import (
	"AIDS_Trigger/Server"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {

	logrus.Info(os.Args)
	if os.Args[1] == "-debug" {
		logrus.SetLevel(logrus.DebugLevel)
	}

	e := echo.New()

	e.POST("/alert", func(c echo.Context) error {
		return Server.Alert(c)
	})

	logrus.Fatal(e.Start(":55555"))

}
