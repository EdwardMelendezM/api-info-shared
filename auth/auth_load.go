package auth

import (
	"time"

	authRepository "github.com/EdwardMelendezM/api-info-shared/auth/infrastructure/jwt"
	auth "github.com/EdwardMelendezM/api-info-shared/auth/interface/rest"
	authRest "github.com/EdwardMelendezM/api-info-shared/auth/interface/rest"
	authUseCase "github.com/EdwardMelendezM/api-info-shared/auth/usecase"
)

func LoadAuthMiddleware() auth.AuthMiddleware {
	timeoutContext := time.Duration(60) * time.Second
	authJWTRepository := authRepository.NewAuthRepository()
	authUCase := authUseCase.NewAuthUseCase(authJWTRepository, timeoutContext)
	authMiddleware := authRest.NewAuthMiddleware(authUCase)
	return authMiddleware
}
