package usecase

import (
	"roomate/model/dto"
	"roomate/utils/common"
)

type AuthUseCase interface {
	Login(payload dto.AuthRequestDto) (dto.AuthResponseDto, error)
}

type authUseCase struct {
	userUc   UserUseCase
	jwtToken common.JwtToken
}

func (a *authUseCase) Login(payload dto.AuthRequestDto) (dto.AuthResponseDto, error) {
	user, err := a.userUc.GetByEmailPassword(payload.Email, payload.Password)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}

	token, err := a.jwtToken.GenerateToken(user)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}

	return token, nil
}

func NewAuthUseCase(userUc UserUseCase, jwtToken common.JwtToken) AuthUseCase {
	return &authUseCase{
		userUc:   userUc,
		jwtToken: jwtToken,
	}
}
