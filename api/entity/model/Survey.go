package model

import (
	"github.com/mariotj/survey-app/api/libraries/exception"
	"strings"
	"time"
)

type Survey struct {
	ID			uint32		`gorm:"primary_key;auto_increment" json:"id"`
	Name 		string		`gorm:"size:255;not null" json:"name"`
	CreatedAt 	time.Time	`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt 	time.Time 	`gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (s *Survey) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if s.Name == "" {
			return exception.RequiredFieldException("name")
		}
		if s.ID == 0 {
			return exception.RequiredFieldException("ID")
		}

	default:
		if s.Name == "" {
			return exception.RequiredFieldException("name")
		}
	}

	return nil
}