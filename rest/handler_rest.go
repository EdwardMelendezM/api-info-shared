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
	smartErr, ok := err.(errDomain.CustomError)
	if ok {
		c.JSON(smartErr.HttpStatus, smartErr)
		// REVIEW save log
		//ctx := c.Request.Context()
		return
	}
	smartErr2, ok := err.(*errDomain.CustomError)
	if ok {
		c.JSON(smartErr2.HttpStatus, smartErr2)
	} else {
		smartErr = errDomain.ErrUnknown
		smartErr.Raw = err.Error()
		c.JSON(http.StatusInternalServerError, smartErr)
	}
	// REVIEW save log
	//ctx := c.Request.Context()
}
