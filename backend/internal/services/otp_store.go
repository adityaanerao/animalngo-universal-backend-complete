package services

import (
	"sync"
	"time"
)

type otpEntry struct {
	Code      string
	Name      string
	ExpiresAt time.Time
}

type OTPStore struct {
	data map[string]otpEntry
	mu   sync.Mutex
}

func NewOTPStore() *OTPStore {
	return &OTPStore{
		data: make(map[string]otpEntry),
	}
}

func (s *OTPStore) Set(mobile, code, name string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[mobile] = otpEntry{
		Code:      code,
		Name:      name,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}
}

// Verify returns the name if successful, otherwise an error boolean
func (s *OTPStore) Verify(mobile, code string) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	entry, exists := s.data[mobile]
	if !exists {
		return "", false
	}

	if time.Now().After(entry.ExpiresAt) || entry.Code != code {
		return "", false
	}

	delete(s.data, mobile)
	return entry.Name, true
}
