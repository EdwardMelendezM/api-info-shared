package usecase

import "context"

func (u authUseCase) GenerateToken(userId string) (*string, error) {
	return u.authRepository.GenerateToken(userId)
}

func (u authUseCase) DecodeToken(context context.Context, token string) (userId *string, err error) {
	return u.authRepository.DecodeToken(context, token)
}
