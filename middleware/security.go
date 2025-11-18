package middleware

import (
	"financial-tracker-api/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// SecurityMiddleware adds security headers and CORS configuration
func SecurityMiddleware() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     utils.GetAllowedOrigins(),
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE, echo.OPTIONS},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, "X-Requested-With"},
		AllowCredentials: utils.GetEnvAsBool("CORS_ALLOW_CREDENTIALS", false),
		MaxAge:           utils.GetEnvAsInt("CORS_MAX_AGE", 86400), // 24 hours
	})
}

// SecurityHeaders adds security headers to all responses
func SecurityHeaders() echo.MiddlewareFunc {
	return middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "DENY",
		HSTSMaxAge:            31536000, // 1 year
		HSTSExcludeSubdomains: false,
		ContentSecurityPolicy: "default-src 'self'; script-src 'self'; style-src 'self' 'unsafe-inline'; img-src 'self' data: https:; font-src 'self'; connect-src 'self'; media-src 'self'; object-src 'none'; child-src 'none'; frame-ancestors 'none'; base-uri 'self'; form-action 'self';",
		ReferrerPolicy:        "strict-origin-when-cross-origin",
	})
}
