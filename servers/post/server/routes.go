package server

import (
	"post/controller"
	"post/utils"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func NewRouter(DB *gorm.DB, redisClient *redis.Client) *gin.Engine {
	r := gin.Default()

	// Set up verification utils
	postUtils := utils.NewPostUtils(DB, redisClient)

	// Public routes - without auth middleware

	// Public routes - with auth middleware
	postAuthGroup := r.Group("/v1/request")
	// TODO: Add auth middleware
	{
		postAuthGroup.POST("/request", controller.AddPostHandler(postUtils))
	}

	return r
}
