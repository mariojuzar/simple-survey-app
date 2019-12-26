package controller

import (
	"github.com/gin-gonic/gin"
	baseResponse "github.com/mariotj/survey-app/api/entity/response"
	"github.com/mariotj/survey-app/api/libraries/util"
	"github.com/mariotj/survey-app/api/service"
	"net/http"
	"time"
)

func getReportService() service.ReportService {
	return service.NewReportService()
}

var reportService = getReportService()

func GetReportBySurveyId(c *gin.Context) {
	id := util.StrToUint(c.Params.ByName("id"))
	report, err := reportService.GetAllAnswerBySurveyId(id)

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
		response.Data = report

		c.JSON(http.StatusOK, response)
	}
}