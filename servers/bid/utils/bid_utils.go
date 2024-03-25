package utils

import (
	"bid/db"
	"bid/middleware"
	"bid/schema"
)

type IBidUtils interface {
	GetbidByID(bidID uint) (schema.Bid, error)
	Addbid(bid schema.Bid) (schema.Bid, error)
}

type BidUtils struct {
	DB          db.Database
	RedisClient middleware.RedisClientInterface
}

// NewbidUtils creates a new bidUtils
func NewBidUtils(DB db.Database, redisClient middleware.RedisClientInterface) *BidUtils {
	return &BidUtils{
		DB:          DB,
		RedisClient: redisClient,
	}
}

// func GetBidByID retrieves a bid by its ID
func (bu *BidUtils) GetBidByID(bidID uint) (schema.Bid, error) {
	var bid schema.Bid
	err := bu.DB.First(&bid, bidID).Error
	if err != nil {
		return schema.Bid{}, err
	}

	return bid, nil
}

// func Addbid adds a bid to the database
func (bu *BidUtils) AddBid(bid schema.Bid) (schema.Bid, error) {
	err := bu.DB.Create(&bid).Error
	if err != nil {
		return schema.Bid{}, err
	}

	return bid, nil
}
