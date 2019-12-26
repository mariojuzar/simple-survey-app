package model

import (
	"github.com/mariotj/survey-app/api/libraries/exception"
	"strings"
	"time"
)

type SurveyQuestion struct {
	ID			uint32		`gorm:"primary_key;auto_increment" json:"id"`
	Survey		Survey		`json:"survey"`
	SurveyId	uint32		`gorm:"not null" json:"survey_id"`
	Question 	string		`gorm:"size:255;not null" json:"question"`
	CreatedAt 	time.Time	`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt 	time.Time 	`gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (s *SurveyQuestion) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if s.Question == "" {
			return exception.RequiredFieldException("question")
		}
		if s.ID == 0 {
			return exception.RequiredFieldException("ID")
		}
	default:
		if s.Question == "" {
			return exception.RequiredFieldException("question")
		}
		if s.SurveyId == 0 {
			return exception.RequiredFieldException("survey id")
		}
	}
	return nil
}