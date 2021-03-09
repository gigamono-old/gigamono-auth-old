package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sageflow/sageflow/pkg/inits"
)

// HandleTokenRoutes handles "/token" routes.
func HandleTokenRoutes(group *gin.RouterGroup, app *inits.App) {
	token := group.Group("/token")
	token.POST("/", tokenHandler)
}

func tokenHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Handling the /token route")
}
