package Server

import (
	"AIDS_Trigger/Message"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	conn *echo.Echo
}

func StartServer() *Server {
	e := echo.New()

	e.POST("/alert", alertController)

	logrus.Fatal(e.Start(":55555"))

	return &Server{e}
}

func alertController(c echo.Context) error {
	logrus.Debug("Alert Controller Called")
	msg := new(Message.DetectedMessage)
	err := c.Bind(&msg)
	logrus.Debugf("Body: [%+v]", msg)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	return c.JSON(http.StatusOK, "Success")
}
