package middleware

import "github.com/labstack/echo/v4"

// SetupMiddleware configures all middleware for the Echo instance
func SetupMiddleware(e *echo.Echo) {
	// Recovery middleware should be first to catch panics
	e.Use(RecoveryMiddleware())
	
	// Request ID for tracing
	e.Use(RequestIDMiddleware())
	
	// Request timeout
	e.Use(TimeoutMiddleware())
	
	// Body limit to prevent memory exhaustion
	e.Use(BodyLimitMiddleware())
	
	// Compression
	e.Use(CompressMiddleware())
	
	// Request logging
	e.Use(CustomRequestLogger())
	
	// Security headers
	e.Use(SecurityHeaders())
	
	// CORS
	e.Use(SecurityMiddleware())
	
	// Rate limiting
	e.Use(RateLimiter())
}
