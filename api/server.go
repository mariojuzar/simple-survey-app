package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mariotj/survey-app/api/controller"
	"github.com/mariotj/survey-app/api/entity/path"
	"github.com/mariotj/survey-app/api/entity/response"
	"github.com/mariotj/survey-app/api/service"
	"net/http"
	"time"
)

func Run() *gin.Engine {
	engine := gin.Default()
	engine.RedirectTrailingSlash = false

	service.Initialize()

	v1 := engine.Group(path.BaseUrl)
	{
		// routes for survey
		v1.POST(path.Survey, controller.CreateSurvey)
		v1.PUT(path.Survey, controller.UpdateSurvey)
		v1.DELETE(path.SurveyById, controller.DeleteSurvey)
		v1.GET(path.SurveyById, controller.GetSurveyById)
		v1.GET(path.Survey, controller.GetAllSurvey)

		// routes for question
		v1.POST(path.SurveyQuestion, controller.AddQuestionBySurveyId)
		v1.PUT(path.SurveyQuestion, controller.UpdateQuestionById)
		v1.GET(path.SurveyQuestionBySurveyId, controller.GetAllQuestionBySurveyId)
		v1.DELETE(path.SurveyQuestionQuestionById, controller.DeleteQuestionById)
		v1.DELETE(path.SurveyQuestionBySurveyId, controller.DeleteAllQuestionBySurveyId)

		// routes for answer
		v1.POST(path.SurveyAnswer, controller.AddAnswerByQuestionId)
		v1.PUT(path.SurveyAnswer, controller.UpdateAnswerById)

		// routes for report
		v1.GET(path.Report, controller.GetReportBySurveyId)
	}

	engine.NoRoute(func(context *gin.Context) {
		var resp = &response.BaseResponse{
			ServerTime:	time.Now(),
		}

		resp.Code = http.StatusNotFound
		resp.Message = "Route not found"

		context.JSON(http.StatusNotFound, resp)
	})

	return engine
}