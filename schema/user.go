package schema

import "time"

// GORM provides a defined gorm.Model struct which includes sensible defaults,
// have a look in the future: https://gorm.io/docs/models.html
type User struct {
	ID				 uint
	Name			 string
	Password	 string
	CreatedAt	 time.Time
	UpdatedAt	 time.Time
}

