package handlers

import (
	"net/http"

	"uam-golang/internal/models"
	"uam-golang/internal/repository"

	"github.com/gin-gonic/gin"
)

// GetProfile godoc
// @Summary      Get user profile
// @Description  Get the profile of the currently authenticated user
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  models.User
// @Failure      401  {object}  models.ErrorResponse
// @Failure      404  {object}  models.ErrorResponse
// @Router       /users/me [get]
func GetProfile(c *gin.Context) {
	email, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var user models.User
	if result := repository.DB.Where("email = ?", email).First(&user); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"email":     user.Email,
		"role":      user.Role,
		"createdAt": user.CreatedAt,
		"updatedAt": user.UpdatedAt,
	})
}

// GetAllUsers godoc
// @Summary      Get all users
// @Description  Get a list of all registered users
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   models.User
// @Failure      500  {object}  models.ErrorResponse
// @Router       /users [get]
func GetAllUsers(c *gin.Context) {
	var users []models.User
	// Fetch all users but omit password in response via JSON tags
	if result := repository.DB.Find(&users); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	// Manual mapping or just relying on json:"-" for password
	c.JSON(http.StatusOK, users)
}
