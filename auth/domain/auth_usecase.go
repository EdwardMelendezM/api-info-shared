package domain

import "context"

type AuthUseCase interface {
	GenerateToken(userId string) (*string, error)
	DecodeToken(context context.Context, token string) (userId *string, err error)
}
