package email

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/hryhorskyi/gin-app/db"
	"github.com/hryhorskyi/gin-app/models"
)

func SendEmails(rate float64) {
	var subscriptions []models.Subscription
	db.DB.Find(&subscriptions)

	for _, subscription := range subscriptions {
		sendEmail(subscription.Email, rate)
	}
}

func sendEmail(to string, rate float64) {
	from := os.Getenv("EMAIL_FROM")
	password := os.Getenv("EMAIL_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	msg := []byte("To: " + to + "\r\n" +
		"Subject: Daily USD to UAH Exchange Rate\r\n" +
		"\r\n" +
		"Current USD to UAH exchange rate is " + fmt.Sprintf("%.2f", rate) + ".\r\n")

	auth := smtp.PlainAuth("", from, password, smtpHost)
	if err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, msg); err != nil {
		log.Printf("Failed to send email to %s: %v", to, err)
	}
}
