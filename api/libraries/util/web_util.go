package util

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/mariotj/survey-app/api/entity/model"
)

func GetRequestBodySurvey(c *gin.Context, survey *model.Survey)  {
	decoder := json.NewDecoder(c.Request.Body)

	_ = decoder.Decode(survey)
}

func GetRequestBodySurveyQuestion(c *gin.Context, question *model.SurveyQuestion)  {
	decoder := json.NewDecoder(c.Request.Body)

	_ = decoder.Decode(question)
}

func GetRequestBodySurveyAnswer(c *gin.Context, answer *model.SurveyAnswer)  {
	decoder := json.NewDecoder(c.Request.Body)

	_ = decoder.Decode(answer)
}