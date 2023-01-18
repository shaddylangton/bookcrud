package models
import(
	"time"

)
type  Book struct{

	ID uint  `json:"id" gorm:"primary_key"`
	Booktittle string  `json:"book_tittle"`
	Bookcategory string  `json:"book_category"`
	BookNo string `json:"book_no"`
	ISBNNo string `json:"isbn_no"`
	PublisherName string `json:"publisher_name"`
	AuthorName string `json:"author_name"`
	ShelfNo uint `json:"shelf_no"`
	RackNo uint `json:"rack_no"`
	Quantity uint `json:"quantity"`
	BookPrice uint `json:"book_price"`
	Description string `json:"description"`
	Hireprice  uint `json:"hire_price"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

}