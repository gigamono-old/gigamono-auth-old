package crud

import (
	"github.com/gigamono/gigamono/pkg/errs"
	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gigamono/gigamono/pkg/messages"
	"github.com/gigamono/gigamono/pkg/security"
	"github.com/gigamono/gigamono/pkg/services/session"
	"github.com/gin-gonic/gin"
)

// getSecurityKeys gets the service's private and public keys.
func getSecurityKeys(app *inits.App, sessionType string) ([]byte, []byte) {
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

	return []byte(privateKey), []byte(publicKey)
}

// verifyPreSession verifies existing pre-session ids, creating new csrf id,
func verifyPreSession(ctx *gin.Context, sessionType string, publicKey []byte) *errs.ClientError {
	// Check for pre-session credentials and their validity.
	// Presession helps us prevent login CSRF. More info in ./presession.go.
	if err := session.VerifyPreSessionCredentials(ctx, publicKey); err != nil {
		switch err.(type) {
		case errs.ClientError:
			clientErr := err.(errs.ClientError)
			return &clientErr
		default:
			panic(errs.NewSystemError(
				messages.Error[sessionType].(string),
				"verifying existing pre-session credentials",
				err,
			))
		}
	}

	return nil
}

// generateSessionTokens creates new CSRF IDs, access and refresh tokens, and sending them in response.
func generateSessionTokens(ctx *gin.Context, app *inits.App, sessionType string, userID string, email string, privateKey []byte, publicKey []byte) {
	// Generate new session random CSRF ID.
	plaintextCSRFID, err := security.GenerateRandomBase64(32)
	if err != nil {
		panic(errs.NewSystemError(
			messages.Error[sessionType].(string),
			"generating random base64 string for session CSRF ID",
			err,
		))
	}

	// Sign/hash plaintext CSRF ID with publicKey key.
	// TODO: Is there really a point to signing CSRF IDs inside JWT claims? JWT itself already requires signing.
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
		// TODO: Sec: Should expire in less time.
		security.GenerateSessionClaims(userID, email, signedCSRFID, security.SessionAccess, 604800), // Expires in a week.
		privateKey,
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
	refreshToken, err := security.GenerateSignedJWT(
		security.GenerateSessionClaims(userID, email, signedCSRFID, security.SessionRefresh, 604800), // Expires in a week.
		privateKey,
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
}
