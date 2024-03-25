package schema

import "time"

type User struct {
	UserID          uint `gorm:"primaryKey"`
	UserName        string
	Email           string
	HashedPassword  string
	ProfileInfo     string
	ReputationScore int
	EmailVerified   bool
	MFAVerified     bool
	MFASecret       string
	DateJoined      time.Time
	LastActiveDate  time.Time
}

type Post struct {
	PostID       uint `gorm:"primaryKey"`
	UserID       uint `gorm:"index"`
	Title        string
	Description  string
	Condition    string
	Availability string
	PostType     string `gorm:"type:enum('Offer', 'Need');default:'Offer'"`
	Status       string `gorm:"type:enum('Active', 'Closed', 'Expired');default:'Active'"`
	DatePosted   time.Time
	DateUpdated  time.Time
}

type PostSubscription struct {
	PostSubscriptionID uint `gorm:"primaryKey"`
	UserID             uint `gorm:"index"`
	PostID             uint `gorm:"index"`
	DateSubscribed     time.Time
}

type Category struct {
	CategoryID   uint `gorm:"primaryKey"`
	CategoryName string
}

type CategorySubscription struct {
	CategorySubscriptionID uint `gorm:"primaryKey"`
	UserID                 uint `gorm:"index"`
	CategoryID             uint `gorm:"index"`
	DateSubscribed         time.Time
}

type PostCategory struct {
	PostID     uint `gorm:"primaryKey;autoIncrement:false"`
	CategoryID uint `gorm:"primaryKey;autoIncrement:false"`
}

type Feedback struct {
	FeedbackID uint `gorm:"primaryKey"`
	FromUserID uint `gorm:"index"`
	ToUserID   uint `gorm:"index"`
	Rating     int
	Comment    string
	DateGiven  time.Time
}

type Exchange struct {
	ExchangeID       uint `gorm:"primaryKey"`
	OfferPostID      uint `gorm:"index"`
	BidID            uint `gorm:"index"`
	AgreementDetails string
	ExchangeDate     time.Time
	ExchangeStatus   string `gorm:"type:enum('Planned', 'Completed', 'Cancelled');default:'Planned'"`
}

type Notification struct {
	NotificationID   uint   `gorm:"primaryKey"`
	UserID           uint   `gorm:"index"`
	NotificationType string `gorm:"type:enum('Bid', 'Match', 'Feedback');default:'Bid'"`
	RelatedID        uint
	Message          string
	DateSent         time.Time
	IsRead           bool
}
