package service

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/mariotj/survey-app/api/entity/model"
	"log"
	"os"
)

type DatabaseService struct {
	DB 				*gorm.DB
	IsInitialized 	bool
}

var databaseService DatabaseService

func Initialize() {
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

	databaseService = DatabaseService{DB:db, IsInitialized:true}

	Load(db)
}

func Load(db *gorm.DB)  {
	err := db.Debug().AutoMigrate(&model.Survey{}, &model.SurveyAnswer{}, &model.SurveyQuestion{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
}