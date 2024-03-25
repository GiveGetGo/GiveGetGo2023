package schema

import "time"

type Match struct {
	MatchID            uint   `gorm:"primaryKey"`
	NeedPostID         uint   `gorm:"index"`
	HelperUserID       uint   `gorm:"index"`
	MatchStatus        string `gorm:"type:enum('Matching', 'Matched', 'Unfulfilled');default:'Matching'"`
	DateMatched        time.Time
	FulfillmentDetails string
}
