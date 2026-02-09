package middleware

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body        *bytes.Buffer
	wroteHeader bool
	status      int
}

func (w *responseBodyWriter) WriteHeader(status int) {
	w.status = status
	// We don't call w.ResponseWriter.WriteHeader(status) yet
}

func (w *responseBodyWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = http.StatusOK
	}
	return w.body.Write(b)
}

func (w *responseBodyWriter) WriteString(s string) (int, error) {
	if w.status == 0 {
		w.status = http.StatusOK
	}
	return w.body.WriteString(s)
}

func (w *responseBodyWriter) Status() int {
	if w.status == 0 {
		return w.ResponseWriter.Status()
	}
	return w.status
}

func ETagMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Only handle GET and HEAD requests
		if c.Request.Method != http.MethodGet && c.Request.Method != http.MethodHead {
			c.Next()
			return
		}

		// Save the original writer
		originalWriter := c.Writer
		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: originalWriter}
		c.Writer = w

		c.Next()

		// If the response was not 200 OK, don't generate ETag
		if w.Status() != http.StatusOK {
			// Write whatever we captured to the original writer
			if w.status != 0 {
				originalWriter.WriteHeader(w.status)
			}
			originalWriter.Write(w.body.Bytes())
			return
		}

		data := w.body.Bytes()
		if len(data) == 0 {
			if w.status != 0 {
				originalWriter.WriteHeader(w.status)
			}
			return
		}

		// Calculate SHA-1 hash
		hash := sha1.Sum(data)
		etag := "\"" + hex.EncodeToString(hash[:]) + "\""

		// Set ETag header
		c.Header("ETag", etag)

		// Check If-None-Match header
		if c.Request.Header.Get("If-None-Match") == etag {
			originalWriter.WriteHeader(http.StatusNotModified)
			return
		}

		// If not modified, we already returned. Otherwise, write the full response.
		// Important: Set the status before writing the body
		if w.status != 0 {
			originalWriter.WriteHeader(w.status)
		}
		originalWriter.Write(data)
	}
}
