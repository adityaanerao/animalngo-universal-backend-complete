package services

import (
	"animalngo-universal-backend/config"
	"animalngo-universal-backend/internal/models"
	"animalngo-universal-backend/internal/repositories"
	jwtutil "animalngo-universal-backend/pkg/jwt"
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthService struct {
	users    *repositories.UserRepository
	db       *gorm.DB
	cfg      *config.Config
	otpStore *OTPStore
	smsSvc   *SMSService
}

func NewAuthService(db *gorm.DB, cfg *config.Config, otpStore *OTPStore, smsSvc *SMSService) *AuthService {
	return &AuthService{repositories.NewUserRepository(db), db, cfg, otpStore, smsSvc}
}

func generateOTP() string {
	b := make([]byte, 3)
	rand.Read(b)
	return fmt.Sprintf("%06d", int(b[0])<<16|int(b[1])<<8|int(b[2]))[:6]
}

func (s *AuthService) SendSignupOTP(mobile, name string) error {
	if _, e := s.users.ByMobileNumber(mobile); e == nil {
		return errors.New("mobile already registered")
	}
	otp := generateOTP()
	if mobile == "9999999999" {
		otp = "123456"
	}
	s.otpStore.Set(mobile, otp, name)
	return s.smsSvc.SendOTP(mobile, otp)
}

func (s *AuthService) SendLoginOTP(mobile string) error {
	if _, e := s.users.ByMobileNumber(mobile); e != nil {
		return errors.New("user not found")
	}
	otp := generateOTP()
	if mobile == "9999999999" {
		otp = "123456"
	}
	s.otpStore.Set(mobile, otp, "")
	return s.smsSvc.SendOTP(mobile, otp)
}

func (s *AuthService) VerifyOTP(mobile, otp string) (*models.User, string, error) {
	name, ok := s.otpStore.Verify(mobile, otp)
	if !ok {
		// allow review bypass
		if mobile == "9999999999" && otp == "123456" {
			// proceed
		} else {
			return nil, "", errors.New("invalid or expired OTP")
		}
	}

	u, e := s.users.ByMobileNumber(mobile)
	if e != nil {
		// User does not exist, this is a signup
		u = &models.User{
			FullName:     name,
			MobileNumber: mobile,
			PositionID:   nil,
			IsActive:     true,
			SessionToken: uuid.New().String(),
		}
		if e = s.users.Create(u); e != nil {
			return nil, "", e
		}
	} else {
		// Login
		if !u.IsActive {
			return nil, "", errors.New("account is disabled")
		}
		u.SessionToken = uuid.New().String()
		s.users.Update(u)
	}

	pidStr := ""
	if u.PositionID != nil {
		pidStr = u.PositionID.String()
	}
	t, e := jwtutil.GenerateToken(u.ID.String(), pidStr, s.cfg.JWTSecret, s.cfg.JWTExpiryHours)
	return u, t, e
}

func (s *AuthService) ByID(id uuid.UUID) (*models.User, error) { return s.users.ByID(id) }
