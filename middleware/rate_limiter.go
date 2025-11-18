package middleware

import (
	"financial-tracker-api/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

// RateLimiter adds rate limiting to prevent abuse
func RateLimiter() echo.MiddlewareFunc {
	rateLimit := rate.Limit(utils.GetEnvAsInt("RATE_LIMIT_PER_SECOND", 20))
	return middleware.RateLimiterWithConfig(middleware.RateLimiterConfig{
		Store: middleware.NewRateLimiterMemoryStore(rateLimit),
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusTooManyRequests, map[string]string{
				"error": "Too many requests, please try again later",
			})
		},
	})
}

// RecoveryMiddleware adds panic recovery
func RecoveryMiddleware() echo.MiddlewareFunc {
	return middleware.Recover()
}
