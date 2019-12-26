package path

var BaseUrl = "/api/v1"
var ID = "/:id"

// survey path
var Survey = "/survey"
var SurveyById = Survey + ID

// survey question path
var SurveyQuestion =  "/question"
var SurveyQuestionById = SurveyQuestion + ID
var SurveyQuestionQuestionById = SurveyQuestion + SurveyQuestion + ID
var SurveyQuestionBySurveyId = SurveyQuestion + Survey + ID

// survey answer path
var SurveyAnswer = Survey + "/answer"
var SurveyAnswerById = SurveyAnswer + ID

// report
var Report = "/report" + SurveyById