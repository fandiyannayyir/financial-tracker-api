package middleware

import (
	"financial-tracker-api/utils"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// BodyLimitMiddleware limits the size of request body to prevent memory exhaustion
func BodyLimitMiddleware() echo.MiddlewareFunc {
	return middleware.BodyLimit(utils.GetEnv("BODY_LIMIT", "10M"))
}

// TimeoutMiddleware adds request timeout to prevent hanging requests
func TimeoutMiddleware() echo.MiddlewareFunc {
	timeout := time.Duration(utils.GetEnvAsInt("REQUEST_TIMEOUT", 30)) * time.Second
	return middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: timeout,
	})
}

// RequestIDMiddleware adds unique request ID to each request for tracking
func RequestIDMiddleware() echo.MiddlewareFunc {
	return middleware.RequestID()
}

// CompressMiddleware adds gzip compression to responses
func CompressMiddleware() echo.MiddlewareFunc {
	return middleware.Gzip()
}
