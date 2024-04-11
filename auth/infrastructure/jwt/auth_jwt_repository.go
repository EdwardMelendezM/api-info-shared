package jwt

import "github.com/EdwardMelendezM/api-info-shared/auth/domain"

type authJwtRepo struct{}

func NewAuthRepository() domain.AuthRepository {
	return &authJwtRepo{}
}
