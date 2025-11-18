package middleware

import (
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// RequestLogger returns a middleware that logs HTTP requests
func RequestLogger() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339_nano} | ${status} | ${latency_human} | ${remote_ip} | ${method} ${uri} | ${error}\n",
		Output: log.Writer(),
	})
}

// CustomRequestLogger provides more detailed logging with structured format
func CustomRequestLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			
			// Process request
			err := next(c)
			
			// Log request details
			req := c.Request()
			res := c.Response()
			
			log.Printf(
				"[%s] %s %s - Status: %d - Duration: %v - IP: %s - User-Agent: %s",
				time.Now().Format("2006-01-02 15:04:05"),
				req.Method,
				req.RequestURI,
				res.Status,
				time.Since(start),
				c.RealIP(),
				req.UserAgent(),
			)
			
			return err
		}
	}
}
