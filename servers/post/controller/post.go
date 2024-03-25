package controller

import (
	"net/http"
	"post/utils"

	"github.com/GiveGetGo/shared/types"
	"github.com/gin-gonic/gin"
)

// AddPostHandler is a function that handles the request to add a post
func AddPostHandler(postUtils utils.IPostUtils) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.LoginRequest
		if err := c.BindJSON(&req); err != nil {
			types.ResponseError(c, http.StatusBadRequest, types.InvalidRequest())
			return
		}

		// TODO: Post validation
		// for now return success
		types.ResponseSuccess(c, http.StatusCreated, "register", 0, types.UserCreated())
	}
}
