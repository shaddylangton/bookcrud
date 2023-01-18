package models

import (
	 "time"
)

type category struct{
	ID uint `json:"id" gorm:"primary_key"`
	CategoryTitle string `json:"category_title"`
	Description string `json:"category_description"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}