package v1

import (
	"net/http"

	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gin-gonic/gin"
)

// HandleTokenRoutes handles "/token" routes.
func HandleTokenRoutes(group *gin.RouterGroup, app *inits.App) {
	token := group.Group("/token")
	token.POST("/", tokenHandler)
}

func tokenHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Handling the /token route")
}
