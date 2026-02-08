package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/user/car-project/internal/service"
)

func RequirePermission(permService service.PermissionService, requiredPerm string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get userID from context (set by AuthMiddleware)
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
			c.Abort()
			return
		}

		// Type assertion
		uid, ok := userID.(int64)
		if !ok {
			// This handles float64 which JWT parsing might return, depending on implementation
			// Safe conversion attempt if int64 fails directly
			if val, ok := userID.(float64); ok {
				uid = int64(val)
			} else {
				// Also handle int if platform dependent
				if val, ok := userID.(int); ok {
					uid = int64(val)
				} else {
					log.Printf("UserID type assertion failed: %T %v", userID, userID)
					c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
					c.Abort()
					return
				}
			}
		}

		// Fetch permissions
		perms, err := permService.GetUserPermissions(c.Request.Context(), uid)
		if err != nil {
			log.Printf("Failed to fetch permissions for user %d: %v", uid, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to check permissions"})
			c.Abort()
			return
		}

		// Check if user has the required permission
		hasPerm := false
		for _, p := range perms {
			if p == requiredPerm {
				hasPerm = true
				break
			}
		}

		if !hasPerm {
			c.JSON(http.StatusForbidden, gin.H{"error": "insufficient permissions"})
			c.Abort()
			return
		}

		c.Next()
	}
}
