package server

import (
	"bid/controller"
	"bid/utils"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func NewRouter(DB *gorm.DB, redisClient *redis.Client) *gin.Engine {
	r := gin.Default()

	// Set up match utils
	bidUtils := utils.NewBidUtils(DB, redisClient)

	// Public routes - without auth middleware
	bidGroup := r.Group("/v1/bid")
	{
		bidGroup.POST("/bid", controller.AddBidHandler(bidUtils))
	}

	return r
}
