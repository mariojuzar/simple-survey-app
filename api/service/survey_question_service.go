package service

import (
	"github.com/mariotj/survey-app/api/entity/model"
	"github.com/mariotj/survey-app/api/libraries/exception"
	"time"
)

type SurveyQuestionService interface {
	AddQuestionBySurveyId(question string, surveyId uint32) (model.SurveyQuestion, error)
	UpdateQuestionById(question string, questionId uint32) (model.SurveyQuestion, error)
	GetAllQuestionBySurveyId(surveyId uint32) ([]model.SurveyQuestion, error)
	DeleteQuestionById(questionId uint32) (model.SurveyQuestion, error)
	DeleteAllQuestionBySurveyId(surveyId uint32) ([]model.SurveyQuestion, error)
}

type surveyQuestionService struct {

}

func (s surveyQuestionService) AddQuestionBySurveyId(question string, surveyId uint32) (model.SurveyQuestion, error) {
	var survey model.Survey
	databaseService.DB.Model(&model.Survey{}).Find(&survey, "id = ?", surveyId)

	surveyQuestion := model.SurveyQuestion{
		Survey:    survey,
		SurveyId:  survey.ID,
		Question:  question,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	databaseService.DB.Create(&surveyQuestion)

	if err := databaseService.DB.GetErrors(); len(err) > 0 {
		return model.SurveyQuestion{}, err[0]
	}

	return surveyQuestion, nil
}

func (s surveyQuestionService) UpdateQuestionById(question string, questionId uint32) (model.SurveyQuestion, error) {
	var surveyQuestionGet model.SurveyQuestion
	databaseService.DB.First(&surveyQuestionGet, "id = ?", questionId)

	if surveyQuestionGet.ID == questionId {
		surveyQuestionGet.Question = question
		surveyQuestionGet.UpdatedAt = time.Now()
		databaseService.DB.Model(&model.SurveyQuestion{}).Where("id = ?", surveyQuestionGet.ID).Updates(surveyQuestionGet)

		return surveyQuestionGet, nil
	} else {
		return model.SurveyQuestion{}, exception.NotFoundException()
	}
}

func (s surveyQuestionService) GetAllQuestionBySurveyId(surveyId uint32) ([]model.SurveyQuestion, error) {
	var surveyQuestions []model.SurveyQuestion
	databaseService.DB.Find(&surveyQuestions, "survey_id = ?", surveyId)

	if err := databaseService.DB.GetErrors(); len(err) > 0 {
		return []model.SurveyQuestion{}, err[0]
	}

	return surveyQuestions, nil
}

func (s surveyQuestionService) DeleteQuestionById(questionId uint32) (model.SurveyQuestion, error) {
	var surveyQuestionGet model.SurveyQuestion
	databaseService.DB.First(&surveyQuestionGet, "id = ?", questionId)

	if surveyQuestionGet.ID == questionId {
		databaseService.DB.Delete(&surveyQuestionGet)

		return surveyQuestionGet, nil
	} else {
		return model.SurveyQuestion{}, exception.NotFoundException()
	}
}

func (s surveyQuestionService) DeleteAllQuestionBySurveyId(surveyId uint32) ([]model.SurveyQuestion, error) {
	var surveyQuestions []model.SurveyQuestion
	databaseService.DB.Find(&surveyQuestions, "survey_id = ?", surveyId)

	databaseService.DB.Delete(&surveyQuestions)

	if err := databaseService.DB.GetErrors(); len(err) > 0 {
		return []model.SurveyQuestion{}, err[0]
	}

	return surveyQuestions, nil
}

func NewSurveyQuestionService() SurveyQuestionService {
	return surveyQuestionService{}
}
