package model

import (
	"github.com/mariotj/survey-app/api/libraries/exception"
	"strings"
	"time"
)

type SurveyAnswer struct {
	ID					uint32			`gorm:"primary_key;auto_increment" json:"id"`
	SurveyQuestion 		SurveyQuestion	`json:"surveyQuestion"`
	SurveyQuestionId 	uint32			`gorm:"not null" json:"survey_question_id"`
	Username 			string			`gorm:"size:255;not null" json:"username"`
	Answer 				string			`gorm:"size:255;not null" json:"answer"`
	CreatedAt 			time.Time		`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt 			time.Time 		`gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (s *SurveyAnswer) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if s.Answer == "" {
			return exception.RequiredFieldException("answer")
		}
		if s.ID == 0 {
			return exception.RequiredFieldException("ID")
		}

	case "create":
		if s.Answer == "" {
			return exception.RequiredFieldException("answer")
		}
		if s.Username == "" {
			return exception.RequiredFieldException("username")
		}
		if s.SurveyQuestionId == 0 {
			return exception.RequiredFieldException("survey question id")
		}
	}
	return nil
}