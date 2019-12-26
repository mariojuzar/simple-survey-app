package service

import "github.com/mariotj/survey-app/api/entity/model"

type ReportService interface {
	GetAllAnswerBySurveyId(surveyId uint32) ([]model.Report, error)
}

type reportService struct {

}

func (r reportService) GetAllAnswerBySurveyId(surveyId uint32) ([]model.Report, error) {
	var answerReport []model.Report
	var questions []model.SurveyQuestion
	databaseService.DB.Find(&questions, "survey_id = ?", surveyId)

	questionIds := getListId(questions)

	databaseService.DB.Table("survey_answers").Select("*").Joins("left join survey_questions on survey_questions.id = survey_answers.survey_question_id where survey_questions.id IN (?)", questionIds).Scan(&answerReport)

	if err := databaseService.DB.GetErrors(); len(err) > 0 {
		return []model.Report{}, err[0]
	}

	return answerReport, nil
}

func NewReportService() ReportService {
	return reportService{}
}

func getListId(questions []model.SurveyQuestion) []uint32 {
	var result []uint32
	for i := range questions {
		result = append(result, questions[i].ID)
	}
	return result
}