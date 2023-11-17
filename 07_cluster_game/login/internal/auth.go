package internal

import (
	"github.com/gin-gonic/gin"
	"learn-pitaya-with-demos/cluster_game/pkg/models"
	"net/http"
)

const (
	Authentication           = "Authentication"
	ContentType              = "Content-Type"
	ContentTypeJson          = "application/json"
	GinContextOfAccountToken = "token"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader(Authentication)
		at, err := ValidateToken(ctx, token)
		if err != nil {
			ctx.Writer.Header().Set(ContentType, ContentTypeJson)
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": "",
				"msg":  "Invalid Token",
			})
			return
		}

		// 将token信息写入context
		ctx.Set(GinContextOfAccountToken, at)
	}
}

func ValidateToken(ctx *gin.Context, token string) (*models.AccountToken, error) {

	return nil, nil
}
