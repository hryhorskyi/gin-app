package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"github.com/go-resty/resty/v2"
	"github.com/hryhorskyi/gin-app/db"
	"github.com/hryhorskyi/gin-app/email"
	"github.com/hryhorskyi/gin-app/handlers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/hryhorskyi/gin-app/docs"
)

// @title GSES BTC Application API
// @version 1.0
// @description This is a sample server for GSES BTC Application.
// @host localhost:8080
// @BasePath /

func main() {
	db.Init()

	r := gin.Default()

	r.GET("/api/rate", handlers.GetRate)
	r.POST("/api/subscribe", handlers.Subscribe)
	r.GET("/api/subscriptions", handlers.GetSubscriptions)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Day().At("08:00").Do(func() {
		rate := getExchangeRate()
		if rate != 0 {
			email.SendEmails(rate)
		}
	})
	s.StartAsync()

	r.Run(":8080")
}

func getExchangeRate() float64 {
	client := resty.New()
	resp, err := client.R().
		Get("https://api.exchangerate-api.com/v4/latest/USD")
	if err != nil || resp.StatusCode() != http.StatusOK {
		log.Println("Failed to fetch exchange rate")
		return 0
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		log.Println("Failed to parse exchange rate")
		return 0
	}

	rates, ok := result["rates"].(map[string]interface{})
	if !ok {
		log.Println("Failed to parse exchange rates")
		return 0
	}

	rate, ok := rates["UAH"].(float64)
	if !ok {
		log.Println("Exchange rate not found")
		return 0
	}

	return rate
}
