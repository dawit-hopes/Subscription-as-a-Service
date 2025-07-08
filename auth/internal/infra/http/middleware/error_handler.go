// Package middleware provides HTTP middleware for error handling and other concerns.
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	appErr "github.com/dawit_hopes/saas/auth/internal/common/errors"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ErrorHandler middleware catches panics and errors and sends JSON response
func ErrorHandler(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			last := c.Errors.Last()

			// If it's an AppError, extract info
			if err, ok := last.Err.(*appErr.AppError); ok {
				logger.Error("Handled application error",
					zap.String("method", c.Request.Method),
					zap.String("path", c.Request.URL.Path),
					zap.String("message", err.Message),
					zap.String("internal", err.Internal),
					zap.Int("status", err.Code),
				)

				c.JSON(err.Code, ErrorResponse{
					Code:    err.Code,
					Message: err.Message,
				})
				return
			}

			// Otherwise fallback to 500
			logger.Error("Unhandled error",
				zap.String("method", c.Request.Method),
				zap.String("path", c.Request.URL.Path),
				zap.Error(last.Err),
			)

			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "Internal Server Error",
			})
		}
	}
}
