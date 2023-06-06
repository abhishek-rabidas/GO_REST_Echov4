package Server

import "github.com/sirupsen/logrus"

type Buffer struct {
	LastActivityTime string
	Active           bool
	DetectionTag     string
	Location         string
}

func (s *ServerConfig) ClearBuffer() {
	logrus.Debug("Clear Buffer")
	s.Buffer = Buffer{
		LastActivityTime: "",
		Active:           false,
		DetectionTag:     "",
		Location:         "",
	}
}
