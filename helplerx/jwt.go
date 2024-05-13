package helplerx

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claims struct {
	UserName string
	UserID   int64
	jwt.RegisteredClaims
}

type JwtPayload struct {
	UserName string
	UserID   int64
	Secret   []byte

	Issuer  string
	Subject string
	ID      string

	TokenExp time.Duration
}

func GenerateJwtToken(payload *JwtPayload) (string, error) {
	now := time.Now()
	expireTime := now.Add(payload.TokenExp)

	claims := Claims{
		UserName: payload.UserName,
		UserID:   payload.UserID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    payload.Issuer,
			Subject:   payload.Subject,
			ExpiresAt: jwt.NewNumericDate(expireTime),
			NotBefore: jwt.NewNumericDate(now.Add(-1 * time.Minute * time.Duration(1))),
			IssuedAt:  jwt.NewNumericDate(now),
			ID:        payload.ID,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(payload.Secret)

	return token, err
}

func ParseToken(token string, secret []byte) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(_ *jwt.Token) (any, error) {
		return secret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
