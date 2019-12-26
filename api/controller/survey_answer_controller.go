package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mariotj/survey-app/api/entity/model"
	baseResponse "github.com/mariotj/survey-app/api/entity/response"
	"github.com/mariotj/survey-app/api/libraries/util"
	"github.com/mariotj/survey-app/api/service"
	"net/http"
	"time"
)

func getSurveyAnswerService() service.SurveyAnswerService {
	return service.NewSurveyAnswerService()
}

var surveyAnswerService = getSurveyAnswerService()

func AddAnswerByQuestionId(c *gin.Context)  {
	var surveyAnswer  model.SurveyAnswer
	util.GetRequestBodySurveyAnswer(c, &surveyAnswer)

	var response = &baseResponse.BaseResponse{
		ServerTime:	time.Now(),
	}

	er := surveyAnswer.Validate("create")

	if er != nil {
		response.Code = http.StatusBadRequest
		response.Message = er.Error()

		c.JSON(http.StatusBadRequest, response)
	} else {
		newSurveyAnswer, err := surveyAnswerService.AddAnswerByQuestionId(surveyAnswer.Answer, surveyAnswer.Username, surveyAnswer.SurveyQuestionId)

		if err != nil {
			response.Code = http.StatusBadRequest
			response.Message = err.Error()

			c.JSON(http.StatusBadRequest, response)
		} else {
			response.Code = http.StatusCreated
			response.Message = http.StatusText(http.StatusCreated)
			response.Data = newSurveyAnswer

			c.JSON(http.StatusCreated, response)
		}
	}
}

func UpdateAnswerById(c *gin.Context)  {
	var surveyAnswer  model.SurveyAnswer
	util.GetRequestBodySurveyAnswer(c, &surveyAnswer)

	var response = &baseResponse.BaseResponse{
		ServerTime:	time.Now(),
	}

	er := surveyAnswer.Validate("update")

	if er != nil {
		response.Code = http.StatusBadRequest
		response.Message = er.Error()

		c.JSON(http.StatusBadRequest, response)
	} else {
		updatedSurveyAnswer, err := surveyAnswerService.UpdateAnswerById(surveyAnswer.Answer, surveyAnswer.SurveyQuestionId)

		if err != nil {
			response.Code = http.StatusBadRequest
			response.Message = err.Error()

			c.JSON(http.StatusBadRequest, response)
		} else {
			response.Code = http.StatusOK
			response.Message = http.StatusText(http.StatusOK)
			response.Data = updatedSurveyAnswer

			c.JSON(http.StatusOK, response)
		}
	}
}
