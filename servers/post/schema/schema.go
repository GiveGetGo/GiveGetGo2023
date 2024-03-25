package schema

import "time"

type Post struct {
	PostID      uint `gorm:"primaryKey"`
	UserID      uint `gorm:"index"`
	Title       string
	Description string
	Status      string `gorm:"type:enum('Active', 'Closed', 'Expired');default:'Active'"`
	DatePosted  time.Time
	DateUpdated time.Time
}
