package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mariotj/survey-app/api"
	"github.com/mariotj/survey-app/api/entity/path"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddAnswerByQuestionId(t *testing.T) {
	err := refreshAllTable()

	if err != nil {
		log.Fatal(err)
	}

	surveys, err := SeedSurveys()
	_, err = SeedQuestions(surveys[0].ID)
	_, err = SeedQuestions(surveys[1].ID)

	samples := []struct{
		inputJSON	string
		statusCode  int
		message		string
	} {
		{
			inputJSON: 	`{"username" : "test", "answer" : "test answer", "survey_question_id" : 1}`,
			statusCode: http.StatusCreated,
			message:	http.StatusText(http.StatusCreated),
		},
		{
			inputJSON: 	`{"username" : "test", "survey_question_id" : 1}`,
			statusCode: http.StatusBadRequest,
			message:	"Required answer",
		},
		{
			inputJSON: 	`{"answer" : "test answer", "survey_question_id" : 1}`,
			statusCode: http.StatusBadRequest,
			message:	"Required username",
		},
		{
			inputJSON: 	`{"answer" : "test answer", "username" : "test"}`,
			statusCode: http.StatusBadRequest,
			message:	"Required survey question id",
		},
		{
			inputJSON: 	`{}`,
			statusCode: http.StatusBadRequest,
			message:	http.StatusText(http.StatusBadRequest),
		},
	}

	router := api.Run()


	for _, sample := range samples {
		req, err := http.NewRequest("POST", path.BaseUrl + path.SurveyAnswer, bytes.NewBufferString(sample.inputJSON))
		if err != nil {
			t.Errorf("this is the error: %v", err)
		}
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal([]byte(rec.Body.String()), &responseMap)
		if err != nil {
			fmt.Printf("Cannot convert to json: %v", err)
		}

		assert.Equal(t, rec.Code, sample.statusCode)
	}

}

func TestUpdateAnswerById(t *testing.T) {
	err := refreshAllTable()

	if err != nil {
		log.Fatal(err)
	}

	surveys, err := SeedSurveys()
	question1, err := SeedQuestions(surveys[0].ID)
	_, err = SeedQuestions(surveys[1].ID)
	_, err = SeedAnswers(question1[0].ID)

	samples := []struct{
		inputJSON	string
		id 			int
		statusCode  int
		message		string
	} {
		{
			inputJSON: 	`{"id": 1, "username" : "test update", "answer" : "test answer", "survey_question_id" : 1}`,
			statusCode: http.StatusOK,
			message:	http.StatusText(http.StatusOK),
		},
		{
			inputJSON: 	`{"username" : "test", "survey_question_id" : 1}`,
			statusCode: http.StatusBadRequest,
			message:	http.StatusText(http.StatusBadRequest),
		},
	}

	router := api.Run()


	for _, sample := range samples {
		req, err := http.NewRequest("PUT", path.BaseUrl + path.SurveyAnswer, bytes.NewBufferString(sample.inputJSON))
		if err != nil {
			t.Errorf("this is the error: %v", err)
		}
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal([]byte(rec.Body.String()), &responseMap)
		if err != nil {
			fmt.Printf("Cannot convert to json: %v", err)
		}

		assert.Equal(t, rec.Code, sample.statusCode)
	}
}