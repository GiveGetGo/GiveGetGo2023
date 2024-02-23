package user

import (
	"errors"
	"net/http"
	"regexp"
	"user_server/schema"
	"user_server/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// userRegisterRequest is the request body for the user registration endpoint
type UserRegisterRequest struct {
	UserName string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UserRegisterResponse is the response body for the user registration endpoint
type UserRegisterResponse struct {
	UserID  uint   `json:"userID"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// UserRegisterHandler is the handler for the user registration endpoint
func UserRegisterHandler(userUtils utils.IUserUtils) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the request body
		var req UserRegisterRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, GeneralUserResponse{
				Code:    "40001",
				Message: "Invalid request body",
			})
			return
		}

		// Check if the email is in the correct format
		matched, _ := regexp.MatchString(`^[a-zA-Z0-9]+@purdue\.edu$`, req.Email)
		if !matched {
			c.JSON(http.StatusBadRequest, GeneralUserResponse{
				Code:    "40002",
				Message: "Email must be a @purdue.edu address",
			})
			return
		}

		// Check if the email is already registered
		user, err := userUtils.GetUserByEmail(req.Email)

		// If no error, or the error is not a record not found error, return an internal server error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusInternalServerError, GeneralUserResponse{
				Code:    "50001",
				Message: "Internal server error",
			})
			return
		}

		// If the user exists, return an error
		if (schema.User{}) != user {
			c.JSON(http.StatusBadRequest, GeneralUserResponse{
				Code:    "40003",
				Message: "Email already registered",
			})
			return
		}

		// Check if the password is valid
		err = userUtils.ValidatePassword(req.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, GeneralUserResponse{
				Code:    "40004",
				Message: "Invalid password",
			})
			return
		}

		// hash the password
		hashedPassword, err := userUtils.HashPassword(req.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, GeneralUserResponse{
				Code:    "50002",
				Message: "Failed to hash password",
			})
			return
		}

		// Create the user
		user, err = userUtils.CreateUser(req.UserName, req.Email, hashedPassword)
		if err != nil {
			c.JSON(http.StatusInternalServerError, GeneralUserResponse{
				Code:    "50003",
				Message: "Failed to create user",
			})
			return
		}

		// call /verification/request-email endpoint in the verification server to send the verification email
		// if the verification server is down, return an internal server error
		err = userUtils.RequestRegisterVerificationEmail(user.UserID, req.UserName, req.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, GeneralUserResponse{
				Code:    "50004",
				Message: "Failed to request verification email",
			})
			return
		}

		// Return the user
		c.JSON(http.StatusCreated, UserRegisterResponse{
			UserID:  user.UserID,
			Code:    "20101",
			Message: "Wait for verification email",
		})
	}
}