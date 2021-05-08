package rest

import (
	"github.com/gigamono/gigamono-auth/internal/crud"
	"github.com/gigamono/gigamono/pkg/errs"
	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gigamono/gigamono/pkg/messages"
	"github.com/gigamono/gigamono/pkg/services/rest/response"
	"github.com/gin-gonic/gin"
)

// OauthTokenForm represents the x-www-form-urlencoded values.
type OauthTokenForm struct {
	GrantType string `form:"grant_type"`
	Email     string `form:"email" binding:"email"`
	Password  string `form:"password"`
	ClientID  string `form:"client_id"`
}

func oauthTokenHandler(app *inits.App) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO: Sec: Validate.

		// Get x-www-form-urlencoded values.
		var form OauthTokenForm
		if err := ctx.ShouldBind(&form); err != nil {
			response.BindErrors(ctx, err)
			return
		}

		// Switch over grant_type.
		switch form.GrantType {
		case "password":
			crud.SignUserUp(ctx)
		default:
			response.FormErrors(
				ctx,
				errs.UnsupportedGrantType,
				messages.Error["grant-type"].(string),
			)
		}
	}
}
