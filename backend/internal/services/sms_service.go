package services

import (
	"log"
)

type SMSService struct{}

func NewSMSService() *SMSService {
	return &SMSService{}
}

func (s *SMSService) SendOTP(mobile, otp string) error {
	// Bypass logic for app review
	if mobile == "9999999999" && otp == "123456" {
		log.Printf("Bypassing SMS for review account %s", mobile)
		return nil
	}
	
	// Simulate external HTTP call to apihome.in
	log.Printf("[SMSService] Sending OTP %s to mobile %s", otp, mobile)
	
	return nil
}
