package Server

import (
	"AIDS_Trigger/Message"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
)

//------------//
//  Handlers //
//----------//

type ServerConfig struct {
	Db *gorm.DB
}

func Init(db *gorm.DB) *ServerConfig {
	return &ServerConfig{db}
}

func (s *ServerConfig) Alert(c echo.Context) error {
	logrus.Debug("Alert Controller Called")
	msg := new(Message.DetectedMessage)
	err := c.Bind(&msg)
	logrus.Debugf("Body: [%+v]", msg)
	s.Db.Create(msg)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	return c.JSON(http.StatusOK, "Success")
}
