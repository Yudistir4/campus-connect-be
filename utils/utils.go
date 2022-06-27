package utils

import (
	"log"
	"math/rand"

	"gopkg.in/gomail.v2"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "Campus Connect <nanasuharnaaa@gmail.com>"
const CONFIG_AUTH_EMAIL = "nanasuharnaaa@gmail.com"
const CONFIG_AUTH_PASSWORD = "cyqosyaezendfdkh"

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func Sendmail(email string, password string) {
	var message string = "<h1>Selamat datang di Campus Connect</h1>" + "<br/>" + "<p>Email: " + email + "</p><br/>" + "<p>Password: " + password + "</p>"

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", email)
	mailer.SetHeader("Subject", "Password Account Campus Connect")
	mailer.SetBody("text/html", message)

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")
}
func SendmailVerified(email string, namaUniversitas string) {
	var message string = "<h1>Selamat datang di Campus Connect</h1>" + "<br/><h4>Account " + namaUniversitas + " dengan email " + email + " sudah terverifikasi dan sudah bisa login</h4>"

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", email)
	mailer.SetHeader("Subject", "Universitas Anda sudah terverifikasi")
	mailer.SetBody("text/html", message)

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")
}
