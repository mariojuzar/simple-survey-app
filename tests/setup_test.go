package tests

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/mariotj/survey-app/api/entity/model"
	"github.com/mariotj/survey-app/api/service"
	"log"
	"os"
	"testing"
)

var database service.DatabaseService

func TestMain(m *testing.M)  {
	var err error
	err = godotenv.Load(os.ExpandEnv(".env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}

	Databases()

	os.Exit(m.Run())
}

func Databases()  {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	DbUser, DbPassword, DbPort, DbHost, DbName := os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME")
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
	db, err := gorm.Open("mysql", dbURL)

	if err != nil {
		log.Fatal("Failed to init db: ", err)
	}
	db.LogMode(true)

	database = service.DatabaseService{DB:db, IsInitialized:true}
}

func refreshAllTable() error {
	err := database.DB.DropTableIfExists(&model.Survey{}, &model.SurveyQuestion{}, &model.SurveyAnswer{}).Error
	if err != nil {
		return err
	}
	err = database.DB.AutoMigrate(&model.Survey{}, &model.SurveyQuestion{}, &model.SurveyAnswer{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed table")
	return nil
}

func SeedSurveys() ([]model.Survey, error)  {
	surveys := []model.Survey{
		model.Survey{
			Name: "test",
		},
		model.Survey{
			Name: "test 2",
		},
	}

	for i := range surveys {
		err := database.DB.Create(&surveys[i]).Error
		if err != nil {
			return []model.Survey{}, err
		}
	}
	return surveys, nil
}

func SeedQuestions(surveyId uint32) ([]model.SurveyQuestion, error)  {
	questions := []model.SurveyQuestion{
		model.SurveyQuestion{
			Question: "test seed",
			SurveyId: surveyId,
		},
		model.SurveyQuestion{
			Question: "test 2",
			SurveyId: surveyId,
		},
	}

	for i := range questions {
		err := database.DB.Create(&questions[i]).Error
		if err != nil {
			return []model.SurveyQuestion{}, err
		}
	}
	return questions, nil
}

func SeedAnswers(questionId uint32) ([]model.SurveyAnswer, error) {
	answer := []model.SurveyAnswer{
		model.SurveyAnswer{
			Answer: "test answer",
			SurveyQuestionId:questionId,
			Username:"test",
		},
		model.SurveyAnswer{
			Answer: "test answer",
			SurveyQuestionId:questionId,
			Username:"test",
		},
	}

	for i := range answer {
		err := database.DB.Create(&answer[i]).Error
		if err != nil {
			return []model.SurveyAnswer{}, nil
		}
	}

	return answer, nil
}