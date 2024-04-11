package usecase

import (
	"time"

	"github.com/EdwardMelendezM/api-info-shared/auth/domain"
)

type authUseCase struct {
	authRepository domain.AuthRepository
}

func NewAuthUseCase(
	authRepository domain.AuthRepository,
	timeout time.Duration,
) domain.AuthUseCase {
	return &authUseCase{
		authRepository: authRepository,
	}
}
