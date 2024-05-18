package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

// GetRate godoc
// @Summary Get current USD to UAH exchange rate
// @Description Get the current exchange rate of USD to UAH using a third-party API
// @Tags rate
// @Produce json
// @Success 200 {object} map[string]float64
// @Failure 400 {object} map[string]string
// @Router /api/rate [get]
func GetRate(c *gin.Context) {
	client := resty.New()
	resp, err := client.R().
		Get("https://api.exchangerate-api.com/v4/latest/USD")

	if err != nil || resp.StatusCode() != http.StatusOK {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to fetch exchange rate"})
		return
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse exchange rate"})
		return
	}

	rates, ok := result["rates"].(map[string]interface{})
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Rates not found in the response"})
		return
	}

	rate, ok := rates["UAH"].(float64)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Exchange rate for UAH not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"rate": rate})
}
