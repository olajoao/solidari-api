package schemas

import (
	"gorm.io/gorm"
)

type Donation struct {
	gorm.Model
	Title       string
	Description string
	Avatar      string
	Location    string
}
