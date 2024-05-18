package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hryhorskyi/gin-app/db"
	"github.com/hryhorskyi/gin-app/models"
)

// Subscribe godoc
// @Summary Subscribe to daily exchange rate updates
// @Description Subscribe an email to receive daily updates on the USD to UAH exchange rate
// @Tags subscription
// @Accept x-www-form-urlencoded
// @Produce json
// @Param email formData string true "Email Address"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/subscribe [post]
func Subscribe(c *gin.Context) {
	email := c.PostForm("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}

	var existingSubscription models.Subscription
	if db.DB.Where("email = ?", email).First(&existingSubscription).RecordNotFound() {
		subscription := models.Subscription{Email: email}
		if err := db.DB.Create(&subscription).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save subscription"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "E-mail added"})
	} else {
		c.JSON(http.StatusConflict, gin.H{"error": "E-mail already exists"})
	}
}

/*

// GetSubscriptions godoc
// @Summary Get all subscriptions
// @Description Get all email subscriptions from the database
// @Tags subscription
// @Produce json
// @Success 200 {array} models.Subscription
// @Failure 500 {object} map[string]string
// @Router /api/subscriptions [get]
func GetSubscriptions(c *gin.Context) {
	var subscriptions []models.Subscription
	if err := db.DB.Find(&subscriptions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch subscriptions"})
		return
	}
	c.JSON(http.StatusOK, subscriptions)
}

*/
