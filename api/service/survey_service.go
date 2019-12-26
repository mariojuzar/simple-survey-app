package service

import (
	"github.com/mariotj/survey-app/api/entity/model"
	"github.com/mariotj/survey-app/api/libraries/exception"
	"time"
)

type SurveyService interface {
	CreateSurvey(name string) (model.Survey, error)
	UpdateSurvey(survey model.Survey) (model.Survey, error)
	DeleteSurvey(surveyId uint32) (model.Survey, error)
	GetSurveyById(surveyId uint32) (model.Survey, error)
	GetAllSurvey() ([]model.Survey, error)
}

type surveyService struct {

}

func (s surveyService) CreateSurvey(name string) (model.Survey, error) {
	survey := model.Survey{
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	databaseService.DB.Create(&survey)

	if err := databaseService.DB.GetErrors(); len(err) > 0 {
		return model.Survey{}, err[0]
	}

	return survey, nil
}

func (s surveyService) UpdateSurvey(survey model.Survey) (model.Survey, error) {
	var surveyGet model.Survey
	databaseService.DB.First(&surveyGet, "id = ?", survey.ID)

	if surveyGet.ID == survey.ID {
		surveyGet.Name = survey.Name
		surveyGet.UpdatedAt = time.Now()
		databaseService.DB.Model(model.Survey{}).Where("id = ?", surveyGet.ID).Updates(surveyGet)

		if err := databaseService.DB.GetErrors(); len(err) > 0 {
			return model.Survey{}, err[0]
		}

		return survey, nil
	} else {
		return model.Survey{}, exception.NotFoundException()
	}
}

func (s surveyService) DeleteSurvey(surveyId uint32) (model.Survey, error) {
	var survey model.Survey
	databaseService.DB.Model(model.Survey{}).Find(&survey, "id = ?", surveyId)

	if survey.ID == surveyId {
		databaseService.DB.Delete(&survey)

		if err := databaseService.DB.GetErrors(); len(err) > 0 {
			return model.Survey{}, err[0]
		}
		return survey, nil
	} else {
		return model.Survey{}, exception.NotFoundException()
	}

}

func (s surveyService) GetSurveyById(surveyId uint32) (model.Survey, error) {
	var survey model.Survey
	databaseService.DB.Model(model.Survey{}).Find(&survey, "id = ?", surveyId)

	if survey.ID == surveyId {
		if err := databaseService.DB.GetErrors(); len(err) > 0 {
			return model.Survey{}, err[0]
		}

		return survey, nil
	} else {
		return model.Survey{}, exception.NotFoundException()
	}
}

func (s surveyService) GetAllSurvey() ([]model.Survey, error) {
	var surveys []model.Survey
	databaseService.DB.Find(&surveys)

	if err := databaseService.DB.GetErrors(); len(err) > 0 {
		return []model.Survey{}, err[0]
	}

	return surveys, nil
}

func NewSurveyService()	SurveyService {
	return surveyService{}
}