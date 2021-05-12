package crud

import (
	"github.com/gigamono/gigamono/pkg/auth"
	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gigamono/gigamono/pkg/messages"
	"github.com/gigamono/gigamono/pkg/services/rest/response"
	"github.com/gigamono/gigamono/pkg/services/session"
	"github.com/gin-gonic/gin"
)

// SignUserUp signs a new user up.
func SignUserUp(app *inits.App) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO: ???
	}
}

// PreSession creates a new presession csrf token.
func PreSession(app *inits.App) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Generate random CSRF token.
		csrfToken, err := auth.GenerateRandomBase64(32)
		if err != nil {
			// TODO!!
		}

		// Get private key.
		privateKey, err := app.Secrets.Get("AUTH_PRIVATE_KEY")
		if err != nil {
			// TODO!!
		}

		// Sign CSRF token with private key.
		signedCSRFToken, err := auth.GenerateSignedCSRFToken(csrfToken, []byte(privateKey))
		if err != nil {
			// TODO!!
		}

		// Add CSRF tokens to response.
		session.AttachPreSessionCSRFTokens(ctx, app, csrfToken, signedCSRFToken)

		response.Success(
			ctx,
			messages.Success["pre-session-created"].(string),
			nil,
		)
	}
}
