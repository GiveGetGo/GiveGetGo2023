package schema

import "time"

type Bid struct {
	BidID           uint `gorm:"primaryKey"`
	PostID          uint `gorm:"index"`
	UserID          uint `gorm:"index"`
	BidDescription  string
	TermsConditions string
	DateSubmitted   time.Time
	Status          string `gorm:"type:enum('Submitted', 'Accepted', 'Rejected');default:'Submitted'"`
}
