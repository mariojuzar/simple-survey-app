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

func getSurveyQuestionService() service.SurveyQuestionService {
	return service.NewSurveyQuestionService()
}

var surveyQuestionService = getSurveyQuestionService()

func AddQuestionBySurveyId(c *gin.Context)  {
	var question model.SurveyQuestion
	util.GetRequestBodySurveyQuestion(c, &question)

	var response = &baseResponse.BaseResponse{
		ServerTime:	time.Now(),
	}

	er := question.Validate("create")

	if er != nil {
		response.Code = http.StatusBadRequest
		response.Message = er.Error()

		c.JSON(http.StatusBadRequest, response)
	} else {
		surveyQuestion, err := surveyQuestionService.AddQuestionBySurveyId(question.Question, question.SurveyId)

		if err != nil {
			response.Code = http.StatusBadRequest
			response.Message = err.Error()

			c.JSON(http.StatusBadRequest, response)
		} else {
			response.Code = http.StatusCreated
			response.Message = http.StatusText(http.StatusCreated)
			response.Data = surveyQuestion

			c.JSON(http.StatusCreated, response)
		}
	}
}

func UpdateQuestionById(c *gin.Context)  {
	var question model.SurveyQuestion
	util.GetRequestBodySurveyQuestion(c, &question)

	var response = &baseResponse.BaseResponse{
		ServerTime:	time.Now(),
	}

	er := question.Validate("update")

	if er != nil {
		response.Code = http.StatusBadRequest
		response.Message = er.Error()

		c.JSON(http.StatusBadRequest, response)
	} else {
		surveyQuestion, err := surveyQuestionService.UpdateQuestionById(question.Question, question.ID)

		if err != nil {
			response.Code = http.StatusBadRequest
			response.Message = err.Error()

			c.JSON(http.StatusBadRequest, response)
		} else {
			response.Code = http.StatusOK
			response.Message = http.StatusText(http.StatusOK)
			response.Data = surveyQuestion

			c.JSON(http.StatusOK, response)
		}
	}
}

func GetAllQuestionBySurveyId(c *gin.Context)  {
	id := util.StrToUint(c.Params.ByName("id"))

	var response = &baseResponse.BaseResponse{
		ServerTime:	time.Now(),
	}

	if id == 0 {
		response.Code = http.StatusBadRequest
		response.Message = "Required id"

		c.JSON(http.StatusBadRequest, response)
	} else {
		surveyQuestions, err := surveyQuestionService.GetAllQuestionBySurveyId(id)

		if err != nil {
			response.Code = http.StatusBadRequest
			response.Message = err.Error()

			c.JSON(http.StatusBadRequest, response)
		} else {
			response.Code = http.StatusOK
			response.Message = http.StatusText(http.StatusOK)
			response.Data = surveyQuestions

			c.JSON(http.StatusOK, response)
		}
	}
}

func DeleteQuestionById(c *gin.Context)  {
	id := util.StrToUint(c.Params.ByName("id"))

	var response = &baseResponse.BaseResponse{
		ServerTime:	time.Now(),
	}

	if id == 0 {
		response.Code = http.StatusBadRequest
		response.Message = "Required id"

		c.JSON(http.StatusBadRequest, response)
	} else {
		surveyQuestion, err := surveyQuestionService.DeleteQuestionById(id)

		if err != nil {
			response.Code = http.StatusBadRequest
			response.Message = err.Error()

			c.JSON(http.StatusBadRequest, response)
		} else {
			response.Code = http.StatusOK
			response.Message = http.StatusText(http.StatusOK)
			response.Data = surveyQuestion

			c.JSON(http.StatusOK, response)
		}
	}
}

func DeleteAllQuestionBySurveyId(c *gin.Context)  {
	id := util.StrToUint(c.Params.ByName("id"))

	var response = &baseResponse.BaseResponse{
		ServerTime:	time.Now(),
	}

	if id == 0 {
		response.Code = http.StatusBadRequest
		response.Message = "Required id"

		c.JSON(http.StatusBadRequest, response)
	} else {
		surveyQuestions, err := surveyQuestionService.DeleteAllQuestionBySurveyId(id)

		if err != nil {
			response.Code = http.StatusBadRequest
			response.Message = err.Error()

			c.JSON(http.StatusBadRequest, response)
		} else {
			response.Code = http.StatusOK
			response.Message = http.StatusText(http.StatusOK)
			response.Data = surveyQuestions

			c.JSON(http.StatusOK, response)
		}
	}
}