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

func TestCreateSurvey(t *testing.T) {
	err := refreshAllTable()

	if err != nil {
		log.Fatal(err)
	}

	samples := []struct{
		inputJSON	string
		statusCode  int
		message		string
	} {
		{
			inputJSON: 	`{"name" : "test"}`,
			statusCode: http.StatusCreated,
			message:	http.StatusText(http.StatusCreated),
		},
		{
			inputJSON: 	`{}`,
			statusCode: http.StatusBadRequest,
			message:	http.StatusText(http.StatusBadRequest),
		},
	}

	router := api.Run()


	for _, sample := range samples {
		req, err := http.NewRequest("POST", path.BaseUrl + path.Survey, bytes.NewBufferString(sample.inputJSON))
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

func TestUpdateSurvey(t *testing.T)  {
	err := refreshAllTable()

	if err != nil {
		log.Fatal(err)
	}

	_, err = SeedSurveys()

	if err != nil {
		log.Fatal("Error seeding")
	}

	samples := []struct{
		inputJSON	string
		statusCode  int
		message		string
	} {
		{
			inputJSON: 	`{"name" : "test update", "id" : 1}`,
			statusCode: http.StatusOK,
			message:	http.StatusText(http.StatusOK),
		},
		{
			inputJSON: 	`{}`,
			statusCode: http.StatusBadRequest,
			message:	"Required name",
		},
		{
			inputJSON: 	`{"name" : "test update"}`,
			statusCode: http.StatusBadRequest,
			message:	"Required id",
		},
	}

	router := api.Run()

	for _, sample := range samples {
		req, err := http.NewRequest("PUT", path.BaseUrl + path.Survey, bytes.NewBufferString(sample.inputJSON))
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

func TestDeleteSurvey(t *testing.T) {
	err := refreshAllTable()

	if err != nil {
		log.Fatal(err)
	}

	_, err = SeedSurveys()

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
		url := path.BaseUrl + path.Survey + "/" + strconv.Itoa(sample.id)
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

func TestGetSurveyById(t *testing.T) {
	err := refreshAllTable()

	if err != nil {
		log.Fatal(err)
	}

	_, err = SeedSurveys()

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
		url := path.BaseUrl + path.Survey + "/" + strconv.Itoa(sample.id)
		req, err := http.NewRequest("GET", url, bytes.NewBufferString(sample.inputJSON))
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

func TestGetAllSurvey(t *testing.T) {
	err := refreshAllTable()

	if err != nil {
		log.Fatal(err)
	}

	router := api.Run()

	url := path.BaseUrl + path.Survey
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