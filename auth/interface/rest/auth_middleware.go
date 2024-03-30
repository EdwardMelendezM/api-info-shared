package rest

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"

	"github.com/EdwardMelendezM/info-code-api-shared-v1/auth/domain"
)

type AuthMiddleware interface {
	Auth(c *gin.Context)
	Cors(c *gin.Context)
}

type authMiddleware struct {
	AuthUseCase domain.AuthUseCase
}

func NewAuthMiddleware(authUseCase domain.AuthUseCase) AuthMiddleware {
	authTmp := &authMiddleware{
		AuthUseCase: authUseCase,
	}
	return authTmp
}

func (h authMiddleware) Cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
	c.Header("Access-Control-Allow-Headers", "*")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
}

func (h authMiddleware) Auth(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
	c.Header("Access-Control-Allow-Headers", "*")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	var err error
	var bodyBytes []byte
	var userId *string
	var token string
	requestID := uuid.New().String()
	requestStartAt := time.Now()

	ctx := context.WithValue(c.Request.Context(), "request_id", requestID)
	ctx = context.WithValue(ctx, "request_start_at", requestStartAt)

	authHeader := c.GetHeader("Authorization")
	if authHeader != "" {
		splitToken := strings.Split(authHeader, "Bearer ")
		if len(splitToken) == 2 {
			token = splitToken[1]
		}
	} else {
		c.JSON(http.StatusUnauthorized, errors.New("401"))
		c.Abort()
		return
	}
	userId, err = h.AuthUseCase.DecodeToken(c, token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, errors.New("401"))
		c.Abort()
		return
	}
	if userId != nil {
		c.Set("userId", *userId)
	}
	body := ""
	if c.Request.Body != nil {
		bodyBytes, err = c.GetRawData()
		if err != nil {
			c.JSON(
				http.StatusUnauthorized, gin.H{
					"error":   "ErrSCP2001",
					"message": err.Error(),
				},
			)
			c.Abort()
			return
		}
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		body = string(bodyBytes)
		body = strings.ReplaceAll(body, `"`, `'`)
		body = strings.ReplaceAll(body, "\n", "")
		body = strings.ReplaceAll(body, "\t", "")
		body = strings.ReplaceAll(body, " ", "")
	}

	fields := log.Fields{
		"project":          os.Getenv("PROJECT"),
		"request_id":       requestID,
		"request_start_at": requestStartAt,
		"method":           c.Request.Method,
		"url":              c.Request.URL.String(),
		"request_body":     body,
	}
	if len(body) > 1000 {
		fields["request_body"] = body[0:1000]
	}
	log.WithFields(fields).Info("Request received")

	c.Request = c.Request.WithContext(ctx)
	c.Next()
}
