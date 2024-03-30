package domain

import "context"

type AuthRepository interface {
	GenerateToken(userId string) (*string, error)
	DecodeToken(context context.Context, token string) (userId *string, err error)
}
