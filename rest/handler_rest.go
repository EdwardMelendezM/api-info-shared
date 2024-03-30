package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"

	errDomain "github.com/EdwardMelendezM/info-code-api-shared-v1/error-log"
)

func Json(c *gin.Context, httpStatus int, res any) {
	c.JSON(httpStatus, res)
	// REVIEW save log
	//ctx := c.Request.Context()
}

func ErrJson(c *gin.Context, err error) {
	customErr, ok := err.(errDomain.CustomError)
	if ok {
		c.JSON(customErr.HttpStatus, customErr)
		// REVIEW save log
		//ctx := c.Request.Context()
		return
	}
	customErr2, ok := err.(*errDomain.CustomError)
	if ok {
		c.JSON(customErr2.HttpStatus, customErr2)
	} else {
		customErr = errDomain.ErrUnknown
		customErr.Raw = err.Error()
		c.JSON(http.StatusInternalServerError, customErr)
	}
	// REVIEW save log
	//ctx := c.Request.Context()
}
