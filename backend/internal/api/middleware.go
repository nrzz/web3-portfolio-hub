package api

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/time/rate"
	"gorm.io/gorm"

	"web3-portfolio-dashboard/backend/internal/models"
	"web3-portfolio-dashboard/backend/internal/services"
)

// CORS middleware
func corsMiddleware() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})
}

// Logger middleware
func loggerMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		logger.WithFields(logrus.Fields{
			"client_ip":   param.ClientIP,
			"timestamp":   param.TimeStamp.Format(time.RFC3339),
			"method":      param.Method,
			"path":        param.Path,
			"protocol":    param.Request.Proto,
			"status_code": param.StatusCode,
			"latency":     param.Latency,
			"user_agent":  param.Request.UserAgent(),
			"error":       param.ErrorMessage,
		}).Info("HTTP Request")

		return ""
	})
}

// Rate limiting middleware
func rateLimitMiddleware() gin.HandlerFunc {
	// Create a rate limiter: 100 requests per minute per IP
	limiters := make(map[string]*rate.Limiter)

	return gin.HandlerFunc(func(c *gin.Context) {
		clientIP := c.ClientIP()

		limiter, exists := limiters[clientIP]
		if !exists {
			limiter = rate.NewLimiter(rate.Every(time.Minute/100), 100)
			limiters[clientIP] = limiter
		}

		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded. Please try again later.",
			})
			c.Abort()
			return
		}

		c.Next()
	})
}

// Authentication middleware
func authMiddleware(authService *services.AuthService) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header required",
			})
			c.Abort()
			return
		}

		// Check if the header starts with "Bearer "
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid authorization header format. Expected 'Bearer <token>'",
			})
			c.Abort()
			return
		}

		// Extract the token
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate the token
		userID, err := authService.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or expired token",
			})
			c.Abort()
			return
		}

		// Set user ID in context for handlers to use
		c.Set("user_id", userID)
		c.Next()
	})
}

// Request ID middleware
func requestIDMiddleware() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = generateRequestID()
		}

		c.Header("X-Request-ID", requestID)
		c.Set("request_id", requestID)
		c.Next()
	})
}

// Security headers middleware
func securityHeadersMiddleware() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		c.Header("Content-Security-Policy", "default-src 'self'")
		c.Next()
	})
}

// Error handling middleware
func errorHandlingMiddleware() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		c.Next()

		// Check if there are any errors
		if len(c.Errors) > 0 {
			err := c.Errors.Last()

			// Log the error
			logrus.WithFields(logrus.Fields{
				"request_id": c.GetString("request_id"),
				"path":       c.Request.URL.Path,
				"method":     c.Request.Method,
				"client_ip":  c.ClientIP(),
			}).Error("Request error: ", err.Error())

			// Return appropriate error response
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":      "Internal server error",
				"request_id": c.GetString("request_id"),
			})
		}
	})
}

// Recovery middleware
func recoveryMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return gin.RecoveryWithWriter(&recoveryWriter{logger: logger})
}

// Recovery writer for logging panics
type recoveryWriter struct {
	logger *logrus.Logger
}

func (w *recoveryWriter) Write(p []byte) (n int, err error) {
	w.logger.Error("Panic recovered: ", string(p))
	return len(p), nil
}

// Helper function to generate request ID
func generateRequestID() string {
	return time.Now().Format("20060102150405") + "-" + randomString(8)
}

// Helper function to generate random string
func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[time.Now().UnixNano()%int64(len(charset))]
	}
	return string(b)
}

// Plan-based access middleware
func planMiddleware(db *gorm.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			c.Abort()
			return
		}

		var user models.User
		if err := db.First(&user, "id = ?", userID).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		if user.SubscriptionTier == "basic" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Upgrade required for this feature. Subscribe to unlock analytics and alerts."})
			c.Abort()
			return
		}

		c.Next()
	})
}
