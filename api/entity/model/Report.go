package model

import "time"

type Report struct {
	ID					uint32			`gorm:"primary_key;auto_increment" json:"id"`
	Question 			string			`gorm:"size:255;not null" json:"question"`
	SurveyQuestionId 	uint32			`gorm:"not null" json:"survey_question_id"`
	Username 			string			`gorm:"size:255;not null" json:"username"`
	Answer 				string			`gorm:"size:255;not null" json:"answer"`
	CreatedAt 			time.Time		`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt 			time.Time 		`gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
