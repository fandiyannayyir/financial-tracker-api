# Security Implementation

## Security Headers yang Diterapkan

### 1. Security Headers
- **X-XSS-Protection**: `1; mode=block` - Perlindungan dari XSS attacks
- **X-Content-Type-Options**: `nosniff` - Mencegah MIME type sniffing
- **X-Frame-Options**: `DENY` - Mencegah clickjacking attacks
- **Strict-Transport-Security**: `max-age=31536000` - Memaksa HTTPS selama 1 tahun
- **Content-Security-Policy**: Policy ketat untuk mencegah XSS dan injection attacks
- **Referrer-Policy**: `strict-origin-when-cross-origin` - Kontrol informasi referrer

### 2. CORS (Cross-Origin Resource Sharing)
- **Configurable Origins**: Dapat dikonfigurasi melalui environment variable `ALLOWED_ORIGINS`
- **Method Control**: Hanya methods yang diizinkan (GET, POST, PUT, PATCH, DELETE, OPTIONS)
- **Header Control**: Kontrol header yang diizinkan
- **Credentials Control**: Dapat dikonfigurasi melalui `CORS_ALLOW_CREDENTIALS`

### 3. Rate Limiting
- **Default Rate**: 20 requests per detik per IP
- **Configurable**: Dapat diubah melalui environment variable `RATE_LIMIT_PER_SECOND`
- **Memory Store**: Menggunakan in-memory storage untuk rate limiting
- **Error Response**: Response JSON yang informatif saat limit terlampaui

### 4. Request Protection
- **Timeout Protection**: Default 30 detik, mencegah hanging requests
- **Body Limit**: Default 10MB, mencegah memory exhaustion attacks
- **Panic Recovery**: Graceful handling saat terjadi runtime panic

## Request Logging

### 1. Detailed Logging
Setiap request dicatat dengan informasi:
- **Timestamp**: Waktu request masuk
- **Method**: HTTP method (GET, POST, dll)
- **URI**: Request URI yang diakses
- **Status Code**: Response status code
- **Duration**: Waktu pemrosesan request
- **Client IP**: IP address client
- **User-Agent**: Browser/client information

### 2. Request ID Tracking
- **Unique ID**: Setiap request mendapat ID unik
- **Header**: `X-Request-Id` ditambahkan ke response
- **Tracing**: Memudahkan debugging dan monitoring

## Environment Configuration

### Security Environment Variables
```bash
# CORS Configuration
ALLOWED_ORIGINS=http://localhost:3000,http://localhost:3001
CORS_ALLOW_CREDENTIALS=false
CORS_MAX_AGE=86400

# Rate Limiting
RATE_LIMIT_PER_SECOND=20

# Request Protection
REQUEST_TIMEOUT=30
BODY_LIMIT=10M
```

### Production Recommendations
1. **CORS Origins**: Selalu specify origins yang tepat, jangan gunakan "*" di production
2. **Rate Limiting**: Sesuaikan dengan kebutuhan traffic aplikasi
3. **Request Timeout**: Sesuaikan dengan kompleksitas operasi aplikasi
4. **Body Limit**: Sesuaikan dengan kebutuhan upload file

## Security Best Practices yang Diterapkan

### 1. Defense in Depth
- Multiple layers of security middleware
- Each middleware handles specific security concerns
- Fail-safe defaults untuk semua konfigurasi

### 2. Configurable Security
- Semua security settings dapat dikonfigurasi via environment variables
- Different settings untuk development dan production
- Easy to adjust based on security requirements

### 3. Monitoring and Logging
- Comprehensive request logging
- Unique request IDs untuk easy tracing
- Rate limiting logs untuk monitoring abuse attempts

### 4. Performance Considerations
- Gzip compression untuk optimize bandwidth
- Efficient rate limiting dengan memory store
- Minimal overhead dari security middleware

## Implementasi Middleware

Middleware diimplementasikan dengan urutan yang tepat:

1. **Recovery** - Catch panics (harus paling awal)
2. **Request ID** - Add tracking ID
3. **Timeout** - Request timeout protection
4. **Body Limit** - Prevent large payload attacks
5. **Compression** - Response compression
6. **Logging** - Request logging
7. **Security Headers** - Security headers
8. **CORS** - Cross-origin support
9. **Rate Limiting** - Abuse prevention (paling akhir)

## Testing Security

Untuk test security implementation:

```bash
# Test rate limiting
for i in {1..25}; do curl -w "%{http_code}\n" http://localhost:8080/api/users; done

# Test CORS
curl -H "Origin: http://localhost:3000" -H "Access-Control-Request-Method: POST" -H "Access-Control-Request-Headers: X-Requested-With" -X OPTIONS http://localhost:8080/api/users

# Test body limit
curl -X POST -H "Content-Type: application/json" -d '{"large_data":"...very large payload..."}' http://localhost:8080/api/users

# Test security headers
curl -I http://localhost:8080/api/users
```

## Next Steps

Security features yang bisa ditambahkan di masa depan:
1. **JWT Authentication**: User authentication dan authorization
2. **Input Validation**: Comprehensive input sanitization
3. **SQL Injection Protection**: Prepared statements dan ORM
4. **HTTPS Enforcement**: Force HTTPS in production
5. **API Versioning**: Version control untuk API endpoints
6. **Audit Logging**: Detailed audit trails untuk sensitive operations
