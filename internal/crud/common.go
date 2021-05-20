package crud

import (
	"github.com/gigamono/gigamono/pkg/auth"
	"github.com/gigamono/gigamono/pkg/errs"
	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gigamono/gigamono/pkg/messages"
	"github.com/gigamono/gigamono/pkg/services/session"
	"github.com/gin-gonic/gin"
)

// establishASession starts a new user session.
//
// It contains operations common to both sign-in and sign-up, which include
// verifying existing pre-session ids,
// creating new csrf id,
// creating access and refresh tokens, and
// sending them in response.
func establishASession(ctx *gin.Context, app *inits.App, sessionType string, userID string) *errs.ClientError {
	// Get public key.
	publicKey, err := app.Secrets.Get("AUTH_PUBLIC_KEY")
	if err != nil {
		panic(errs.NewSystemError(
			messages.Error[sessionType].(string),
			"trying to retrieve public key from secrets manager",
			err,
		))
	}

	// Get private key.
	privateKey, err := app.Secrets.Get("AUTH_PRIVATE_KEY")
	if err != nil {
		panic(errs.NewSystemError(
			messages.Error[sessionType].(string),
			"trying to retrieve private key from secrets manager",
			err,
		))
	}

	// Check for pre-session credentials and their validity.
	// Presession helps us prevent login CSRF. More info in ./presession.go.
	if err := session.VerifyPreSessionCredentials(ctx, []byte(publicKey)); err != nil {
		switch err.(type) {
		case *errs.ClientError:
			return err.(*errs.ClientError)
		default:
			panic(errs.NewSystemError(
				messages.Error[sessionType].(string),
				"verifying existing pre-session credentials",
				err,
			))
		}
	}

	// Generate new session random CSRF ID.
	plaintextCSRFID, err := auth.GenerateRandomBase64(32)
	if err != nil {
		panic(errs.NewSystemError(
			messages.Error[sessionType].(string),
			"generating random base64 string for session CSRF ID",
			err,
		))
	}

	// Sign/hash plaintext CSRF ID with publicKey key.
	signedCSRFID, err := auth.GenerateSignedCSRFID(plaintextCSRFID, []byte(publicKey))
	if err != nil {
		panic(errs.NewSystemError(
			messages.Error[sessionType].(string),
			"generating signed and hashed CSRF ID",
			err,
		))
	}

	// Generate access token.
	accessToken, err := auth.GenerateSignedJWT(
		auth.GenerateSessionClaims(userID, signedCSRFID, sessionType, 86400), // Expires in a day.
		[]byte(privateKey),
	)
	if err != nil {
		panic(errs.NewSystemError(
			messages.Error[sessionType].(string),
			"creating CSRF JWT access token",
			err,
		))
	}

	// Generate refresh token.
	// SEC: Refresh tokens have action "refresh" to distinguish them from access tokens and pre-session tokens.
	refreshToken, err := auth.GenerateSignedJWT(
		auth.GenerateSessionClaims(userID, signedCSRFID, "refresh", 86400), // Expires in a day.
		[]byte(privateKey),
	)
	if err != nil {
		panic(errs.NewSystemError(
			messages.Error["pre-session"].(string),
			"creating CSRF JWT refresh token",
			err,
		))
	}

	// Add tokens to response.
	if err = session.AttachSessionTokens(ctx, app, plaintextCSRFID, accessToken, refreshToken); err != nil {
		panic(errs.NewSystemError(
			messages.Error[sessionType].(string),
			"setting session cookies and csrf custom header",
			err,
		))
	}

	return nil
}
