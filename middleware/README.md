# Middleware Documentation

This document describes the security and logging middleware implemented in the Financial Tracker API.

## Security Middleware

### 1. Recovery Middleware
- **Purpose**: Catches panics and prevents server crashes
- **Position**: First in the middleware chain
- **Configuration**: Default Echo recovery middleware

### 2. Request ID Middleware
- **Purpose**: Adds unique request ID to each request for tracing
- **Headers Added**: `X-Request-Id`
- **Configuration**: Default Echo request ID middleware

### 3. Timeout Middleware
- **Purpose**: Prevents hanging requests
- **Default Timeout**: 30 seconds
- **Environment Variable**: `REQUEST_TIMEOUT` (in seconds)

### 4. Body Limit Middleware
- **Purpose**: Prevents memory exhaustion from large request bodies
- **Default Limit**: 10MB
- **Environment Variable**: `BODY_LIMIT` (e.g., "10M", "5MB")

### 5. Compression Middleware
- **Purpose**: Compresses response bodies using gzip
- **Configuration**: Default Echo gzip middleware

### 6. Security Headers Middleware
- **Purpose**: Adds security headers to all responses
- **Headers Added**:
  - `X-XSS-Protection: 1; mode=block`
  - `X-Content-Type-Options: nosniff`
  - `X-Frame-Options: DENY`
  - `Strict-Transport-Security: max-age=31536000`
  - `Content-Security-Policy: default-src 'self'; ...`
  - `Referrer-Policy: strict-origin-when-cross-origin`

### 7. CORS Middleware
- **Purpose**: Handles Cross-Origin Resource Sharing
- **Environment Variables**:
  - `ALLOWED_ORIGINS`: Comma-separated list of allowed origins (default: "*")
  - `CORS_ALLOW_CREDENTIALS`: Allow credentials (default: false)
  - `CORS_MAX_AGE`: Preflight cache time in seconds (default: 86400)

### 8. Rate Limiting Middleware
- **Purpose**: Prevents API abuse
- **Default Rate**: 20 requests per second
- **Environment Variable**: `RATE_LIMIT_PER_SECOND`
- **Storage**: In-memory store
- **Response**: 429 Too Many Requests when limit exceeded

## Logging Middleware

### Request Logger
- **Purpose**: Logs all HTTP requests with detailed information
- **Log Format**: 
  ```
  [YYYY-MM-DD HH:MM:SS] METHOD URI - Status: XXX - Duration: XXXms - IP: XXX.XXX.XXX.XXX - User-Agent: ...
  ```
- **Information Logged**:
  - Timestamp
  - HTTP method
  - Request URI
  - Response status code
  - Request duration
  - Client IP address
  - User-Agent header

## Environment Configuration

Create a `.env` file based on `.env.example` to configure the middleware:

```bash
# Security Configuration
ALLOWED_ORIGINS=http://localhost:3000,http://localhost:3001
CORS_ALLOW_CREDENTIALS=false
CORS_MAX_AGE=86400

# Rate Limiting
RATE_LIMIT_PER_SECOND=20

# Request Timeout (in seconds)
REQUEST_TIMEOUT=30

# Body Limit
BODY_LIMIT=10M
```

## Security Best Practices

1. **CORS Configuration**: In production, always specify exact allowed origins instead of using "*"
2. **Rate Limiting**: Adjust rate limits based on your API usage patterns
3. **Request Timeout**: Set appropriate timeouts to prevent resource exhaustion
4. **Body Limits**: Set reasonable body size limits to prevent memory attacks
5. **Security Headers**: The implemented headers provide protection against common web vulnerabilities

## Monitoring and Debugging

- All requests are logged with unique request IDs for easy tracing
- Rate limiting errors return structured JSON responses
- Request duration is logged for performance monitoring
- Client IP addresses are logged for security analysis

## Middleware Order

The middleware are applied in the following order (important for proper functionality):

1. Recovery (catch panics)
2. Request ID (add tracing)
3. Timeout (request timeout)
4. Body Limit (prevent large payloads)
5. Compression (response compression)
6. Logging (request logging)
7. Security Headers (security headers)
8. CORS (cross-origin support)
9. Rate Limiting (abuse prevention)
