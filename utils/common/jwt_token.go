package common

import (
	"encoding/base64"
	"errors"
	"fmt"
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
	// decode private key
	decodedPrivateKey, err := base64.StdEncoding.DecodeString(j.cfg.JwtPrivateKey)
	if err != nil {
		return dto.AuthResponseDto{}, fmt.Errorf("could not decode key: %w", err)
	}

	// parse decoded private key
	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)
	if err != nil {
		return dto.AuthResponseDto{}, fmt.Errorf("failed to parse decoded private key: %w", err)
	}

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

	// sign token
	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return dto.AuthResponseDto{}, fmt.Errorf("failed to sign token: %w", err)
	}

	return dto.AuthResponseDto{Token: token}, nil
}

func (j *jwtToken) VerifyToken(tokenString string) (jwt.MapClaims, error) {
	// decode public key
	decodedPublicKey, err := base64.StdEncoding.DecodeString(j.cfg.JwtPublicKey)
	if err != nil {
		return nil, fmt.Errorf("could not decode public key %w", err)
	}

	// parse public key
	key, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
	if err != nil {
		return nil, fmt.Errorf("validate: parse key: %w", err)
	}

	// parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", token.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	// validate token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func (j *jwtToken) RefreshToken(oldTokenString string) (dto.AuthResponseDto, error) {
	// get private key
	decodedPrivateKey, err := base64.StdEncoding.DecodeString(j.cfg.JwtPrivateKey)
	if err != nil {
		return dto.AuthResponseDto{}, fmt.Errorf("could not decode key: %w", err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)
	if err != nil {
		return dto.AuthResponseDto{}, fmt.Errorf("failed to parse decoded private key: %w", err)
	}

	// get public key
	decodedPublicKey, err := base64.StdEncoding.DecodeString(j.cfg.JwtPublicKey)
	if err != nil {
		return dto.AuthResponseDto{}, fmt.Errorf("could not decode public key %w", err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
	if err != nil {
		return dto.AuthResponseDto{}, fmt.Errorf("validate: parse key: %w", err)
	}

	// parse token
	token, err := jwt.Parse(oldTokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", token.Header["alg"])
		}
		return publicKey, nil
	})

	if err != nil {
		return dto.AuthResponseDto{}, fmt.Errorf("error parsing token: %w", err)
	}

	// validate token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return dto.AuthResponseDto{}, errors.New("invalid token")
	}

	// set new token life time to 24 hours
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()

	newToken, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(privateKey)
	if err != nil {
		return dto.AuthResponseDto{}, fmt.Errorf("create: sign token: %w", err)
	}

	return dto.AuthResponseDto{Token: newToken}, nil
}

func NewJwtToken(cfg config.TokenConfig) JwtToken {
	return &jwtToken{cfg: cfg}
}
