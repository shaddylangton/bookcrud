package models
import(
	"time"

)

type issue struct{

	ID uint `json:"id" gorm:"primary_key"`
	ISBNNo string `json:"isbn_no"`
	MemberID string `json:"member_id"`
	IssueDate time.Time `json:"date"`
	IssueDays uint `json:"Days"`
	ReturnDate time.Time `json:"return_date"`
	Charge uint `json:"charges"`
	Token string `json:"token"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

}