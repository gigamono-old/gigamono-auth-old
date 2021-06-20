package crud

import (
	"github.com/gigamono/gigamono/pkg/errs"
	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gigamono/gigamono/pkg/messages"
	"github.com/gigamono/gigamono/pkg/security"
	"github.com/gigamono/gigamono/pkg/services/rest/response"
	"github.com/gigamono/gigamono/pkg/services/session"
	"github.com/gin-gonic/gin"
)

// PreSession creates a pre-session CSRF ID and access token for additional security.
//
// Pre-session credentials prevent login CSRF where attacker logs into attacker's session on victim's browser.
//
// Login or signup cannot be established without existing pre-session credentials expected to be found in
// a HttpOnly cookie and a custom header which an attacker cannot replicate in a cross-origin request from victim's browser.
//
// Without the HttpOnly cookie and custom header challenge, an attacker may log into their session easily.
//
// A crypto-secure pseudo-random CSRF ID is generated on behalf of the user and stored in the custom header.
// A signed form of the plaintext CSRF ID is stored in a JWT access token and set as a HttpOnly domain-only cookie.
func PreSession(app *inits.App) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get security keys.
		sessionType := "pre-session"

		privateKey, publicKey := getSecurityKeys(app, sessionType)

		// Generate random CSRF ID.
		plaintextCSRFID, err := security.GenerateRandomBase64(32)
		if err != nil {
			panic(errs.NewSystemError(
				messages.Error[sessionType].(string),
				"generating random base64 string for pre-session CSRF ID",
				err,
			))
		}

		// Sign/hash plaintext CSRF ID with private key.
		signedCSRFID, err := security.GenerateSignedCSRFID(plaintextCSRFID, publicKey)
		if err != nil {
			panic(errs.NewSystemError(
				messages.Error[sessionType].(string),
				"generating signed and hashed CSRF ID",
				err,
			))
		}

		// Generate access token.
		accessToken, err := security.GenerateSignedJWT(
			security.GeneratePreSessionClaims(signedCSRFID, 86400), // Expires in a day.
			privateKey,
		)
		if err != nil {
			panic(errs.NewSystemError(
				messages.Error[sessionType].(string),
				"creating CSRF JWT access token",
				err,
			))
		}

		// Add tokens to response.
		if err = session.AttachPreSessionTokens(ctx, app, plaintextCSRFID, accessToken); err != nil {
			panic(errs.NewSystemError(
				messages.Error[sessionType].(string),
				"setting csrf token cookie and custom header",
				err,
			))
		}

		response.Success(
			ctx,
			messages.Success["pre-session-created"].(string),
			nil,
		)
	}
}
