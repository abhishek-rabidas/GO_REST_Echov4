package Server

import (
	"AIDS_Trigger/Message"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

//------------//
//  Handlers //
//----------//

func Alert(c echo.Context) error {
	logrus.Debug("Alert Controller Called")
	msg := new(Message.DetectedMessage)
	err := c.Bind(&msg)
	logrus.Debugf("Body: [%+v]", msg)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	return c.JSON(http.StatusOK, "Success")
}
