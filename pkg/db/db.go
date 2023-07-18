package db

import (
	"log"

	"github.com/aulaleslie/asset-service/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	err = db.AutoMigrate(&models.AssetsGroup{})
	if err != nil {
		log.Fatalln(err)
	}

	err = db.AutoMigrate(&models.AssetSubGroup{})
	if err != nil {
		log.Fatalln(err)
	}

	err = db.AutoMigrate(&models.BusinessGroup{})
	if err != nil {
		log.Fatalln(err)
	}

	return Handler{db}
}
