package helper

import (
	"fmt"
	"log"
	"math/rand"
	"net/smtp"
	"time"

	"github.com/bibin-zoz/social-match-room-svc/pkg/config"
)

var otpMap = make(map[string]string)

func GenerateOTP() string {
	source := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(source)
	return fmt.Sprintf("%06d", randGen.Intn(1000000))
}

func SendOTP(email string, cfg config.Config) (string, error) {
	otp := GenerateOTP()
	from := cfg.Email
	password := cfg.Password
	to := email
	log.Println("email", email, otp)
	smtpServer := "smtp.gmail.com"
	smtpPort := "587"
	otpMap[email] = otp

	auth := smtp.PlainAuth("", from, password, smtpServer)
	fmt.Println("auth", auth)

	message := fmt.Sprintf("Subject: Your OTP\n\nYour OTP is: %s", otp)

	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, from, []string{to}, []byte(message))
	if err != nil {
		fmt.Println("error", err)
		return "", err
	}

	return otp, nil
}

func VerifyOTP(otp string, email string) bool {
	userEmail := email
	enteredOTP := otp

	storedOTP, ok := otpMap[userEmail]
	if !ok {
		fmt.Println("otp not found")
		return false
	}
	fmt.Println("storedOTP", storedOTP, otp)
	if enteredOTP == storedOTP {
		// Clear the OTP from the map after successful verification
		delete(otpMap, userEmail)
		// Render HTML page with a success message
		// c.HTML(http.StatusOK, "verify.html", gin.H{"message": "OTP verified successfully"})
		// // Send JSON response with the same success message
		// c.JSON(http.StatusOK, gin.H{"message": "OTP verified successfully"})
		return true
	} else {
		return false
	}
}
