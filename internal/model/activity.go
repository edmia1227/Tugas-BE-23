package model

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	ID           int
	name         string
	Description  string
	ActivityDate time.Time
}

type ActivityModel struct {
	db *gorm.DB
}

func NewActivityModel(connection *gorm.DB) *ActivityModel {
	return &ActivityModel{
		db: connection,
	}
}

func (am *ActivityModel) AddActivity(newData Activity) (Activity, error) {
	err := am.db.Create(&newData).Error
	if err != nil {
		return Activity{}, err
	}
	return newData, nil
}
