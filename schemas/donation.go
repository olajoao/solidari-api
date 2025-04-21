package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Donation struct {
	gorm.Model
	Title       string
	Description string
	Avatar      string
	Location    string
}

type DonationResponse struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at,omitempty"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Avatar      string    `json:"avatar"`
	Location    string    `json:"location"`
}
