package models

import(
	"time"

)

type member struct{
	ID uint `json:"id" gorm:"primary_key"`
	MemberID string `json:"member_id"`
    Name string `json:"member_name"`
    Type string `json:"member_type"`
	Mail string `json:"member_mail"`
	Mobile string `json:"member_phone"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

}