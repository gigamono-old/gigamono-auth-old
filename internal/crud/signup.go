package crud

import (
	controllers "github.com/gigamono/gigamono/pkg/database/controllers/auth"
	"github.com/gigamono/gigamono/pkg/errs"
	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gigamono/gigamono/pkg/messages"
	"github.com/gigamono/gigamono/pkg/security"
	"github.com/gigamono/gigamono/pkg/services/rest/response"
	"github.com/gigamono/gigamono/pkg/services/session"
	"github.com/gin-gonic/gin"
)

// SignUpResponse is the handler's response.
type SignUpResponse struct {
	ID string `json:"id"`
}

// SignUserUp adds a new user to the system provided an email, a password and presession tokens.
func SignUserUp(app *inits.App) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO: Sec: Validation
		sessionType := "signup"

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

		// Get security keys.
		privateKey, publicKey := getSecurityKeys(app, sessionType)

		// Verify pre-session.
		if clientErr := verifyPreSession(ctx, sessionType, publicKey); clientErr != nil {
			response.BadRequestErrors(ctx, clientErr)
			return
		}

		// Hash password using argon2id with 10 iterations.
		passwordHash, err := security.GeneratePasswordHash(plaintextPassword, 10)
		if err != nil {
			panic(errs.NewSystemError(
				messages.Error[sessionType].(string),
				"generating password key/hash",
				err,
			))
		}

		// TODO: Duplicate email check. Get from pg error?
		// Create new user account access in db.
		userID, err := controllers.CreateUserAccountCreds(&app.DB, email, passwordHash)
		if err != nil {
			panic(errs.NewSystemError(
				messages.Error[sessionType].(string),
				"resgistering user account credentials in the database",
				err,
			))
		}

		// Generate session tokens.
		generateSessionTokens(ctx, app, sessionType, userID.String(), privateKey, publicKey)

		response.Success(
			ctx,
			messages.Success["user-created"].(string),
			SignUpResponse{ID: userID.String()},
		)
	}
}
