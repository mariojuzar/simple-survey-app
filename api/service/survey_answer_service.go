package service

import (
	"github.com/mariotj/survey-app/api/entity/model"
	"github.com/mariotj/survey-app/api/libraries/exception"
	"time"
)

type SurveyAnswerService interface {
	AddAnswerByQuestionId(answer, name string, questionId uint32) (model.SurveyAnswer, error)
	UpdateAnswerById(answer string, answerId uint32) (model.SurveyAnswer, error)
}

type surveyAnswerService struct {

}

func (s surveyAnswerService) AddAnswerByQuestionId(answer, name string, questionId uint32) (model.SurveyAnswer, error) {
	var question model.SurveyQuestion
	databaseService.DB.Model(&model.SurveyQuestion{}).Find(&question, "id = ?", questionId)

	surveyAnswer := model.SurveyAnswer{
		SurveyQuestion:   question,
		SurveyQuestionId: question.ID,
		Username:         name,
		Answer:           answer,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	databaseService.DB.Create(&surveyAnswer)

	if err := databaseService.DB.GetErrors(); len(err) > 0 {
		return model.SurveyAnswer{}, err[0]
	}

	return surveyAnswer, nil
}

func (s surveyAnswerService) UpdateAnswerById(answer string, answerId uint32) (model.SurveyAnswer, error) {
	var surveyAnswer model.SurveyAnswer
	databaseService.DB.First(&surveyAnswer, "id = ?", answerId)

	if surveyAnswer.ID == answerId {
		surveyAnswer.Answer = answer
		databaseService.DB.Model(&model.SurveyAnswer{}).Where("id = ?", surveyAnswer.ID).Updates(&surveyAnswer)

		if err := databaseService.DB.GetErrors(); len(err) > 0 {
			return model.SurveyAnswer{}, err[0]
		}

		return surveyAnswer, nil
	} else {
		return model.SurveyAnswer{}, exception.NotFoundException()
	}
}

func NewSurveyAnswerService() SurveyAnswerService {
	return surveyAnswerService{}
}