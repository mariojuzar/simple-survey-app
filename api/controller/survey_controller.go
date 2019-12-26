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

func getSurveyService() service.SurveyService  {
	return service.NewSurveyService()
}

var surveyService = getSurveyService()

func CreateSurvey(c *gin.Context)  {
	var survey model.Survey
	util.GetRequestBodySurvey(c, &survey)

	var response = &baseResponse.BaseResponse{
		ServerTime:	time.Now(),
	}

	er := survey.Validate("create")

	if er != nil {
		response.Code = http.StatusBadRequest
		response.Message = er.Error()

		c.JSON(http.StatusBadRequest, response)
	} else {
		newSurvey, err := surveyService.CreateSurvey(survey.Name)

		if err != nil {
			response.Code = http.StatusBadRequest
			response.Message = err.Error()

			c.JSON(http.StatusBadRequest, response)
		} else {
			response.Code = http.StatusCreated
			response.Message = http.StatusText(http.StatusCreated)
			response.Data = newSurvey

			c.JSON(http.StatusCreated, response)
		}
	}
}

func UpdateSurvey(c *gin.Context) {
	var survey model.Survey
	util.GetRequestBodySurvey(c, &survey)

	er := survey.Validate("update")

	var response = &baseResponse.BaseResponse{
		ServerTime:	time.Now(),
	}

	if er != nil {
		response.Code = http.StatusBadRequest
		response.Message = er.Error()

		c.JSON(http.StatusBadRequest, response)
	} else {
		updatedSurvey, err := surveyService.UpdateSurvey(survey)

		if err != nil {
			response.Code = http.StatusBadRequest
			response.Message = err.Error()

			c.JSON(http.StatusBadRequest, response)
		} else {
			response.Code = http.StatusOK
			response.Message = http.StatusText(http.StatusOK)
			response.Data = updatedSurvey

			c.JSON(http.StatusOK, response)
		}
	}
}

func DeleteSurvey(c *gin.Context)  {
	id := util.StrToUint(c.Params.ByName("id"))

	var response = &baseResponse.BaseResponse{
		ServerTime:	time.Now(),
	}

	if id == 0 {
		response.Code = http.StatusBadRequest
		response.Message = "Required id"

		c.JSON(http.StatusBadRequest, response)
	} else {
		survey, err := surveyService.DeleteSurvey(id)

		if err != nil {
			response.Code = http.StatusBadRequest
			response.Message = err.Error()

			c.JSON(http.StatusBadRequest, response)
		} else {
			response.Code = http.StatusOK
			response.Message = http.StatusText(http.StatusOK)
			response.Data = survey

			c.JSON(http.StatusOK, response)
		}
	}
}

func GetSurveyById(c *gin.Context)  {
	id := util.StrToUint(c.Params.ByName("id"))

	var response = &baseResponse.BaseResponse{
		ServerTime:	time.Now(),
	}

	if id == 0 {
		response.Code = http.StatusBadRequest
		response.Message = "Required id"

		c.JSON(http.StatusBadRequest, response)
	} else {
		survey, err := surveyService.GetSurveyById(id)

		if err != nil {
			response.Code = http.StatusBadRequest
			response.Message = err.Error()

			c.JSON(http.StatusBadRequest, response)
		} else {
			response.Code = http.StatusOK
			response.Message = http.StatusText(http.StatusOK)
			response.Data = survey

			c.JSON(http.StatusOK, response)
		}
	}
}

func GetAllSurvey(c *gin.Context)  {
	survey, err := surveyService.GetAllSurvey()

	var response = &baseResponse.BaseResponse{
		ServerTime:	time.Now(),
	}

	if err != nil {
		response.Code = http.StatusBadRequest
		response.Message = err.Error()

		c.JSON(http.StatusBadRequest, response)
	} else {
		response.Code = http.StatusOK
		response.Message = http.StatusText(http.StatusOK)
		response.Data = survey

		c.JSON(http.StatusOK, response)
	}
}