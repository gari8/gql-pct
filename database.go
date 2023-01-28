package main

import (
	"github.com/gari8/gqlgen-pct/repository/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func NewDatabase(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(
		&model.Program{},
		&model.Place{},
	); err != nil {
		return nil, err
	}
	if err := seed(db); err != nil {
		log.Println("already make seed!")
	}
	db = db.Debug()
	return db, nil
}

func seed(db *gorm.DB) error {
	db.Create([]*model.Program{
		{
			ID:          "1",
			Name:        "Program 1",
			Description: "Description 1",
			Image:       "Image 1",
			ProgramType: 0,
			PlaceID:     "1",
		},
		{
			ID:          "2",
			Name:        "Program 2",
			Description: "Description 2",
			Image:       "Image 2",
			ProgramType: 1,
			PlaceID:     "1",
		},
		{
			ID:          "3",
			Name:        "Program 3",
			Description: "Description 3",
			Image:       "Image 3",
			ProgramType: 2,
			PlaceID:     "2",
		},
	})
	db.Create([]*model.Place{
		{
			ID:        "1",
			Name:      "Place 1",
			Latitude:  1.0,
			Longitude: 1.0,
			PlaceType: 0,
		},
		{
			ID:        "2",
			Name:      "Place 2",
			Latitude:  2.0,
			Longitude: 2.0,
			PlaceType: 1,
		},
	})
	return nil
}
