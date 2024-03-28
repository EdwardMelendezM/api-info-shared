package errorLog

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"sync"

	log "github.com/sirupsen/logrus"
)

// PanicRecovery Error recovery panic in goroutine
func PanicRecovery(ctx *context.Context, err *error) {
	if r := recover(); r != nil {
		stackTrace := getStackTrace()
		if err != nil {
			*err = fmt.Errorf("panic: %v", r)
		}
		log.WithFields(log.Fields{
			"error":      stackTrace,
			"project":    os.Getenv("PROJECT"),
			"request_id": getRequestIDFromContext(ctx),
		}).Panic("panic")
	}
}

// PanicThreadRecovery Error recovery in thread
func PanicThreadRecovery(ctx *context.Context, err *error, wg *sync.WaitGroup) {
	if r := recover(); r != nil {
		stackTrace := getStackTrace()
		if err != nil {
			*err = fmt.Errorf("panic: %v", r)
		}
		if wg != nil {
			wg.Done()
		}
		log.WithFields(log.Fields{
			"error":      stackTrace,
			"project":    os.Getenv("PROJECT"),
			"request_id": getRequestIDFromContext(ctx),
		}).Panic("panic in thread")
	}

}

func getStackTrace() string {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	return string(buf[:n])
}

func getRequestIDFromContext(ctx *context.Context) string {
	if ctx == nil {
		return ""
	}
	if requestID, ok := (*ctx).Value("request_id").(string); ok {
		return requestID
	}
	return ""
}
