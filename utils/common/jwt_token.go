package common

import (
	"errors"
	"roomate/config"
	"roomate/model/dto"
	"roomate/model/entity"
	modelutil "roomate/utils/model_util"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtToken interface {
	GenerateToken(payload entity.User) (dto.AuthResponseDto, error)
	VerifyToken(token string) (jwt.MapClaims, error)
	RefreshToken(oldTokenString string) (dto.AuthResponseDto, error)
}

type jwtToken struct {
	cfg config.TokenConfig
}

func (j *jwtToken) GenerateToken(payload entity.User) (dto.AuthResponseDto, error) {
	// claims is jwt payload
	claims := modelutil.JwtTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    payload.Name,
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(j.cfg.JwtLifeTime)),
		},
		UserId: payload.Id,
		Role:   payload.RoleName,
	}

	jwtNewClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtNewClaims.SignedString(j.cfg.JwtSignatureKey)
	if err != nil {
		return dto.AuthResponseDto{}, errors.New("failed to generate token")
	}

	return dto.AuthResponseDto{Token: token}, nil
}

func (j *jwtToken) VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return j.cfg.JwtSignatureKey, nil
	})

	if err != nil {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func (j *jwtToken) RefreshToken(oldTokenString string) (dto.AuthResponseDto, error) {
	token, err := jwt.Parse(oldTokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return j.cfg.JwtSignatureKey, nil
	})

	if err != nil {
		return dto.AuthResponseDto{}, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return dto.AuthResponseDto{}, errors.New("invalid token")
	}

	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	NewTokenString, err := newToken.SignedString(j.cfg.JwtSignatureKey)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}

	return dto.AuthResponseDto{Token: NewTokenString}, nil
}

func NewJwtToken(cfg config.TokenConfig) JwtToken {
	return &jwtToken{cfg: cfg}
}
