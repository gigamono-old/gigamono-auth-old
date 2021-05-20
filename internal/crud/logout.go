package crud

import (
	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gigamono/gigamono/pkg/messages"
	"github.com/gigamono/gigamono/pkg/services/rest/response"
	"github.com/gin-gonic/gin"
)

// LogUserOut logs a user out.
func LogUserOut(app *inits.App) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO: Sec: Invalidate session tokens immediately.

		response.Success(
			ctx,
			messages.Success["logged-out"].(string),
			nil,
		)
	}
}
