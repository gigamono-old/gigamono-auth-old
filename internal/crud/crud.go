package crud

import (
	"github.com/gigamono/gigamono/pkg/messages"
	"github.com/gigamono/gigamono/pkg/services/rest/response"
	"github.com/gin-gonic/gin"
)

// SignUserUp signs a new user up.
func SignUserUp(ctx *gin.Context) {
	response.SendSuccess(
		ctx,
		messages.Success["user-created"].(string),
		nil,
	)
}
