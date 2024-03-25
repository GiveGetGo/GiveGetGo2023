package controller

import (
	"match/utils"
	"net/http"

	"github.com/GiveGetGo/shared/types"
	"github.com/gin-gonic/gin"
)

// func AddMatchHandler - add a match
func AddMatchHandler(matchUtils *utils.MatchUtils) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.LoginRequest
		if err := c.BindJSON(&req); err != nil {
			types.ResponseError(c, http.StatusBadRequest, types.InvalidRequest())
			return
		}

		// TODO: Match
		// for now return success
		types.ResponseSuccess(c, http.StatusCreated, "register", 0, types.UserCreated())
	}
}
