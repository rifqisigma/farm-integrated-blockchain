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
	mailer.SetBody("text/html", fmt.Sprintf(`<a href="http://localhost:8000/auth/gmail/verification?email=%s&token=%s">Klik di sini untuk verifikasi</a>`, toEmail, token))
	dialer := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("EMAIL_SENDER"), os.Getenv("APP_PASSWORD"))

	fmt.Println("Verification Link:")
	fmt.Printf("http://localhost:8000/auth/gmail/verification?email=%s&token=%s\n", toEmail, token)
	if err := dialer.DialAndSend(mailer); err != nil {
		fmt.Println("Error sending email:", err)
	}
}

func SendEmailResetPassword(toEmail, token string) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", os.Getenv("EMAIL_SENDER"))
	mailer.SetHeader("To", toEmail)
	mailer.SetHeader("Subject", "Reset password")
	mailer.SetBody("text/html", fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<body style="font-family: Arial, sans-serif; background-color: #f2f2f2; padding: 20px;">
		<div style="max-width: 480px; margin: auto; background: #ffffff; border-radius: 12px; padding: 24px; box-shadow: 0 4px 12px rgba(0,0,0,0.1);">
			<h2 style="color: #333;">Reset Password</h2>
			<p>Masukkan password baru kamu di halaman berikut:</p>
			<a href="http://localhost:8000/auth/gmail/reset-password?email=%s&token=%s"
			style="display:inline-block; background:#28a745; color:#fff; text-decoration:none; padding:10px 18px; border-radius:6px; margin-top:10px;">
			Buka Halaman Reset Password
			</a>
			<hr style="margin:20px 0;">
			<p style="font-size:12px; color:#888;">
			Jika tombol tidak berfungsi, salin dan buka link berikut di browser kamu:
			</p>
			<p style="font-size:12px; color:#555; word-break: break-all;">
			http://localhost:8000/auth/gmail/reset-password?email=%s&token=%s
			</p>
		</div>
		</body>
		</html>
`, toEmail, token, toEmail, token))

	dialer := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("EMAIL_SENDER"), os.Getenv("APP_PASSWORD"))

	fmt.Println("Verification Link:")
	fmt.Printf("http://localhost:8000/auth/gmail/reset-password?email=%s&token=%s\n", toEmail, token)
	if err := dialer.DialAndSend(mailer); err != nil {
		fmt.Println("Error sending email:", err)
	}
}
