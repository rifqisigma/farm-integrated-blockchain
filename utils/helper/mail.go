package helper

import (
	"fmt"
	"os"

	"gopkg.in/gomail.v2"
)

func SendEmailValidateEmail(toEmail, token string) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", os.Getenv("EMAIL_SENDER"))
	mailer.SetHeader("To", toEmail)
	mailer.SetHeader("Subject", "Verify Your Account")
	mailer.SetBody("text/html", fmt.Sprintf(`<a href="http://localhost:8080/gmail/verification?email=%s&token=%s">Klik di sini untuk verifikasi</a>`, toEmail, token))
	dialer := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("EMAIL_SENDER"), os.Getenv("APP_PASSWORD"))

	fmt.Println("Verification Link:")
	fmt.Printf("http://localhost:8080/gmail/verification?email=%s&token=%s\n", toEmail, token)
	if err := dialer.DialAndSend(mailer); err != nil {
		fmt.Println("Error sending email:", err)
	}
}

func SendEmailResetPassword(toEmail, token string) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", os.Getenv("EMAIL_SENDER"))
	mailer.SetHeader("To", toEmail)
	mailer.SetHeader("Subject", "Verify Your Account")
	// mailer.SetBody("text/html", fmt.Sprintf(`<a href="http://localhost:8080/gmail/reset-password?email=%s&token=%s">Klik di sini untuk verifikasi</a>`, toEmail, token))
	dialer := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("EMAIL_SENDER"), os.Getenv("APP_PASSWORD"))

	fmt.Println("Verification Link:")
	fmt.Printf("http://localhost:8080/gmail/reset-password?email=%s&token=%s\n", toEmail, token)
	if err := dialer.DialAndSend(mailer); err != nil {
		fmt.Println("Error sending email:", err)
	}
}
