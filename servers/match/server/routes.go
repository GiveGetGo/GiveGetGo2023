package server

import (
	"match/controller"
	"match/utils"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func NewRouter(DB *gorm.DB, redisClient *redis.Client) *gin.Engine {
	r := gin.Default()

	// Set up match utils
	matchUtils := utils.NewMatchUtils(DB, redisClient)

	// Public routes - without auth middleware
	matchGroup := r.Group("/v1/match")
	{
		matchGroup.POST("/match", controller.AddMatchHandler(matchUtils))
	}

	return r
}
