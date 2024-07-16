package models

import( 
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `gorm:"unique"`
	Email     string    `gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"Update_at"`
}
