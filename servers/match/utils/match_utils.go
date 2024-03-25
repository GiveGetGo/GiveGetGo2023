package utils

import (
	"match/db"
	"match/middleware"
	"match/schema"
)

type IMatchUtils interface {
	GetMatchByID(matchID uint) (schema.Match, error)
	AddMatch(match schema.Match) (schema.Match, error)
}

type MatchUtils struct {
	DB          db.Database
	RedisClient middleware.RedisClientInterface
}

// NewMatchUtils creates a new MatchUtils
func NewMatchUtils(DB db.Database, redisClient middleware.RedisClientInterface) *MatchUtils {
	return &MatchUtils{
		DB:          DB,
		RedisClient: redisClient,
	}
}

// func GetMatchByID retrieves a match by its ID
func (mu *MatchUtils) GetMatchByID(matchID uint) (schema.Match, error) {
	var match schema.Match
	err := mu.DB.First(&match, matchID).Error
	if err != nil {
		return schema.Match{}, err
	}

	return match, nil
}

// func AddMatch adds a match to the database
func (mu *MatchUtils) AddMatch(match schema.Match) (schema.Match, error) {
	err := mu.DB.Create(&match).Error
	if err != nil {
		return schema.Match{}, err
	}

	return match, nil
}
