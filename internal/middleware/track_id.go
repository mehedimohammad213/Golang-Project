package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/user/car-project/internal/utils"
)

// TrackIDMiddleware generates a unique track_id per request, sets it in context,
// and logs every request/response with that ID so you can grep logs by track_id.
// Use utils.GetTrackID(c) in handlers/response helpers to include it in JSON.
func TrackIDMiddleware() gin.HandlerFunc {
	logger := utils.GetLogger()
	return func(c *gin.Context) {
		trackID := uuid.New().String()
		c.Set(utils.TrackIDContextKey, trackID)
		c.Header("X-Track-ID", trackID)

		c.Next()

		status := c.Writer.Status()
		if status >= http.StatusBadRequest {
			logger.Printf("[track_id=%s] method=%s path=%s status=%d (error)", trackID, c.Request.Method, c.Request.URL.Path, status)
		} else {
			logger.Printf("[track_id=%s] method=%s path=%s status=%d", trackID, c.Request.Method, c.Request.URL.Path, status)
		}
	}
}
