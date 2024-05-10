package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID uint `json:"userID"`
}

func New(key []byte) HS256Signer {
	return HS256Signer{
		key: key,
	}
}

type HS256Signer struct {
	key []byte
}

func (s HS256Signer) GenerateToken(userID uint, valid time.Duration) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(valid)),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
		UserID: userID,
	}).SignedString(s.key)
}

func (s HS256Signer) ParseToken(token string) (*Claims, error) {
	var t *jwt.Token
	var claims Claims
	var err error
	t, err = jwt.ParseWithClaims(
		token, &claims, func(t *jwt.Token) (interface{}, error) {
			return s.key, nil
		},
		jwt.WithLeeway(time.Second*3),
	)
	if err != nil {
		return nil, err
	} else if !t.Valid {
		return nil, errors.New("token invalid")
	}
	return &claims, nil
}
