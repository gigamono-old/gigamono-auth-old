package crud

import (
	"github.com/gigamono/gigamono/pkg/database/models/auth"
	"github.com/gigamono/gigamono/pkg/errs"
	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gigamono/gigamono/pkg/messages"
	"github.com/gigamono/gigamono/pkg/security"
	"github.com/gigamono/gigamono/pkg/services/rest/response"
	"github.com/gigamono/gigamono/pkg/services/session"
	"github.com/gin-gonic/gin"
)

// SignInResponse is the handler's response.
type SignInResponse struct {
	ID string `json:"id"`
}

// SignUserIn authenticates a user using provided an email, a password and presession tokens.
func SignUserIn(app *inits.App) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO: Sec: Validation
		sessionType := "signin"

		// Using Basic Authentication Scheme, get email and password.
		email, plaintextPassword, err := session.GetBasicAuthCreds(ctx)
		if err != nil {
			switch err.(type) {
			case errs.ClientError:
				response.BasicAuthErrors(ctx, err.Error())
				return
			default:
				panic(errs.NewSystemError(
					messages.Error[sessionType].(string),
					"getting basic auth credentials",
					err,
				))
			}
		}

		// Create new user account access in db.
		accountCreds := auth.UserAccountCreds{Email: email}
		if err = accountCreds.GetByEmail(&app.DB); err != nil {
			panic(errs.NewSystemError(
				messages.Error[sessionType].(string),
				"getting user account credentials in the database",
				err,
			))
		}

		// Compare passwords.
		if err = security.VerifyPasswordHash(plaintextPassword, accountCreds.PasswordHash); err != nil {
			panic(errs.NewSystemError(
				messages.Error[sessionType].(string),
				"verifying user's password",
				err,
			))
		}

		// Get security keys.
		privateKey, publicKey := getSecurityKeys(app, sessionType)

		// Verify pre-session.
		if clientErr := verifyPreSession(ctx, sessionType, publicKey); clientErr != nil {
			response.BadRequestErrors(ctx, clientErr)
			return
		}

		// Generate session tokens.
		generateSessionTokens(ctx, app, sessionType, accountCreds.ID.String(), accountCreds.Email, privateKey, publicKey)

		response.Success(
			ctx,
			messages.Success["user-signed-in"].(string),
			SignInResponse{ID: accountCreds.ID.String()},
		)
	}
}
