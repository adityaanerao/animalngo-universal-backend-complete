package jwtutil

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claims struct {
	UserID     string `json:"user_id"`
	PositionID string `json:"position_id"`
	jwt.RegisteredClaims
}

func GenerateToken(userID, positionID, secret string, hours int) (string, error) {
	now := time.Now()
	c := Claims{UserID: userID, PositionID: positionID, RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(hours) * time.Hour)), IssuedAt: jwt.NewNumericDate(now), Issuer: "animalngo-backend"}}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString([]byte(secret))
}
func ParseToken(tokenString, secret string) (*Claims, error) {
	t, e := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) { return []byte(secret), nil })
	if e != nil {
		return nil, e
	}
	c, ok := t.Claims.(*Claims)
	if !ok || !t.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}
	return c, nil
}
