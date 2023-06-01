package Message

import "gorm.io/gorm"

type DetectedMessage struct {
	gorm.Model
	Timestamp string `json:"timestamp"`
	Class     string `json:"class"`
}
