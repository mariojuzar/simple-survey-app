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
	"strconv"
	"testing"
)

func TestAddQuestionBySurveyId(t *testing.T) {
	err := refreshAllTable()

	if err != nil {
		log.Fatal(err)
	}

	_, err = SeedSurveys()

	samples := []struct{
		inputJSON	string
		statusCode  int
		message		string
	} {
		{
			inputJSON: 	`{"question" : "test question", "survey_id" : 1}`,
			statusCode: http.StatusCreated,
			message:	http.StatusText(http.StatusCreated),
		},
		{
			inputJSON: 	`{}`,
			statusCode: http.StatusBadRequest,
			message:	"Required question",
		},
		{
			inputJSON: 	`{"question" : "test created"}`,
			statusCode: http.StatusBadRequest,
			message:	"Required survey id",
		},
	}

	router := api.Run()


	for _, sample := range samples {
		req, err := http.NewRequest("POST", path.BaseUrl + path.SurveyQuestion, bytes.NewBufferString(sample.inputJSON))
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

func TestUpdateQuestionById(t *testing.T) {
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
			inputJSON: 	`{"question" : "test question", "survey_id" : 1, "id" : 1}`,
			statusCode: http.StatusOK,
			message:	http.StatusText(http.StatusOK),
		},
		{
			inputJSON: 	`{}`,
			statusCode: http.StatusBadRequest,
			message:	"Required question",
		},
		{
			inputJSON: 	`{"question" : "test created"}`,
			statusCode: http.StatusBadRequest,
			message:	"Required survey id",
		},
	}

	router := api.Run()


	for _, sample := range samples {
		req, err := http.NewRequest("PUT", path.BaseUrl + path.SurveyQuestion, bytes.NewBufferString(sample.inputJSON))
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

func TestGetAllQuestionBySurveyId(t *testing.T) {
	err := refreshAllTable()

	if err != nil {
		log.Fatal(err)
	}

	surveys, err := SeedSurveys()
	_, err = SeedQuestions(surveys[0].ID)
	_, err = SeedQuestions(surveys[1].ID)

	router := api.Run()

	url := path.BaseUrl + path.SurveyQuestion + path.Survey + "/1"
	req, err := http.NewRequest("GET", url, bytes.NewBufferString("{}"))
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

	assert.Equal(t, rec.Code, http.StatusOK)
}

func TestDeleteQuestionById(t *testing.T) {
	err := refreshAllTable()

	if err != nil {
		log.Fatal(err)
	}

	surveys, err := SeedSurveys()
	_, err = SeedQuestions(surveys[0].ID)
	_, err = SeedQuestions(surveys[1].ID)

	if err != nil {
		log.Fatal("Error seeding")
	}

	samples := []struct{
		inputJSON	string
		id			int
		statusCode  int
		message		string
	} {
		{
			inputJSON: 	`{}`,
			id: 1,
			statusCode: http.StatusOK,
			message:	http.StatusText(http.StatusOK),
		},
		{
			inputJSON: 	`{}`,
			statusCode: http.StatusBadRequest,
			message:	http.StatusText(http.StatusBadRequest),
		},
		{
			inputJSON: 	`{}`,
			id: 34857,
			statusCode: http.StatusBadRequest,
			message:	"Not Found",
		},
	}

	router := api.Run()

	for _, sample := range samples {
		url := path.BaseUrl + path.SurveyQuestion + path.SurveyQuestion + "/" + strconv.Itoa(sample.id)
		req, err := http.NewRequest("DELETE", url, bytes.NewBufferString(sample.inputJSON))
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

func TestDeleteAllQuestionBySurveyId(t *testing.T) {
	err := refreshAllTable()

	if err != nil {
		log.Fatal(err)
	}

	surveys, err := SeedSurveys()
	_, err = SeedQuestions(surveys[0].ID)
	_, err = SeedQuestions(surveys[1].ID)

	if err != nil {
		log.Fatal("Error seeding")
	}

	samples := []struct{
		inputJSON	string
		id			int
		statusCode  int
		message		string
	} {
		{
			inputJSON: 	`{}`,
			id: 1,
			statusCode: http.StatusOK,
			message:	http.StatusText(http.StatusOK),
		},
		{
			inputJSON: 	`{}`,
			statusCode: http.StatusBadRequest,
			message:	http.StatusText(http.StatusBadRequest),
		},
	}

	router := api.Run()

	for _, sample := range samples {
		url := path.BaseUrl + path.SurveyQuestion + path.Survey + "/" + strconv.Itoa(sample.id)
		fmt.Println(url)
		req, err := http.NewRequest("DELETE", url, bytes.NewBufferString(sample.inputJSON))
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
