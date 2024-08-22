package token

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
)

type Payload struct {
	ID       uuid.UUID `json:"id"`
	UserID   uuid.UUID `json:"userId"`
	Username string    `json:"username"`
}

type jwtCustomClaims struct {
	Payload
	jwt.RegisteredClaims
}

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) *JWTMaker {
	return &JWTMaker{secretKey}
}

const issuer = "sleep-tracking-api"

func (maker *JWTMaker) CreateToken(userID uuid.UUID, duration time.Duration) (string, error) {

	claims := &jwtCustomClaims{
		Payload: Payload{
			ID:     uuid.New(),
			UserID: userID,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(maker.secretKey))
	return token, err
}

func (maker *JWTMaker) ValidateToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &jwtCustomClaims{}, keyFunc)
	if err != nil {
		return nil, ErrInvalidToken
	}

	customClaims, ok := jwtToken.Claims.(*jwtCustomClaims)
	if !ok {
		return nil, ErrInvalidToken
	}

	return &customClaims.Payload, nil
}
