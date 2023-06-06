package Server

import (
	"AIDS_Trigger/Message"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
)

//------------//
//  Handlers //
//----------//

type ServerConfig struct {
	Db     *gorm.DB
	Buffer Buffer
}

func Init(db *gorm.DB) *ServerConfig {
	return &ServerConfig{db, Buffer{Active: false}}
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

	s.Buffer = Buffer{
		LastActivityTime: msg.Timestamp,
		Active:           true,
		DetectionTag:     msg.Class,
		Location:         "GIR NATIONAL PARK",
	}

	fmt.Printf("[%s detected at %s]", msg.Class, msg.Timestamp)
	return c.JSON(http.StatusOK, "Success")
}

func (s *ServerConfig) CheckActivity(c echo.Context) error {
	if s.Buffer.Active {
		return c.JSON(200, s.Buffer)
	} else {
		return c.JSON(955, "No activity")
	}
}
