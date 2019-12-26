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

func TestGetReportBySurveyId(t *testing.T) {
	err := refreshAllTable()

	if err != nil {
		log.Fatal(err)
	}

	router := api.Run()

	url := path.BaseUrl + path.Report
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
