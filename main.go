package main

import (
	"fmt"
	"os"
	"strconv"

	gomail "gopkg.in/gomail.v2"
)

func envOrDefault(key, def string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return def
}

func main() {
	templateFile := "template/professional.html"
	subject := "Example Company Account Update"

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "otp-verify":
			templateFile = "template/otp_verify.html"
			subject = "Your Example Company verification code"
		case "login-success":
			templateFile = "template/login_success.html"
			subject = "Successful login to your Example Company account"
		case "suspicious-login":
			templateFile = "template/suspicious_login.html"
			subject = "Suspicious sign-in attempt detected"
		}
	}

	htmlBody, err := os.ReadFile(templateFile)
	if err != nil {
		panic(fmt.Errorf("failed to load email template: %w", err))
	}

	sender := envOrDefault("SMTP_FROM", "noreply@example.com")
	recipient := envOrDefault("SMTP_TO", "customer@example.com")
	smtpHost := envOrDefault("SMTP_HOST", "localhost")
	smtpPort, err := strconv.Atoi(envOrDefault("SMTP_PORT", "1025"))
	if err != nil {
		panic(fmt.Errorf("invalid SMTP_PORT: %w", err))
	}
	smtpUser := os.Getenv("SMTP_USERNAME") // optional, can be empty
	smtpPass := os.Getenv("SMTP_PASSWORD") // optional, can be empty

	m := gomail.NewMessage()
	m.SetHeader("From", sender)
	m.SetHeader("To", recipient)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", string(htmlBody))

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	fmt.Println("::Email sent successfully::")
}
