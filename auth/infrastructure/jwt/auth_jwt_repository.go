package jwt

import "github.com/EdwardMelendezM/info-code-api-shared-v1/auth/domain"

type authJwtRepo struct{}

func NewAuthRepository() domain.AuthRepository {
	return &authJwtRepo{}
}
