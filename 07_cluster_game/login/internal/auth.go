package internal

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"learn-pitaya-with-demos/cluster_game/pkg/models"
	"net/http"
	"time"
)

const (
	Authentication           = "Authentication"
	ContentType              = "Content-Type"
	ContentTypeJson          = "application/json"
	GinContextOfAccountToken = "token"
)

type tokenClaims struct {
	jwt.RegisteredClaims
	Data []byte `json:"data,omitempty"`
}

func Auth(secure string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader(Authentication)
		at, err := ValidateToken(token, secure)
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

func ValidateToken(tokenString, tokenSecure string) (*models.AccountToken, error) {
	if data, err := parseToken(tokenString, []byte(tokenSecure)); err != nil {
		return nil, err
	} else {
		var ut models.AccountToken
		if err = json.Unmarshal(data, &ut); err != nil {
			return nil, err
		}

		return &ut, nil
	}
}

func signToken(data []byte, expInSeconds int64, key []byte) (string, error) {
	claims := tokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(expInSeconds))),
		},
		Data: data,
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(key)
}

func parseToken(tokenString string, key []byte) ([]byte, error) {
	token, err := jwt.ParseWithClaims(tokenString, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}

	if token != nil {
		if claims, ok := token.Claims.(*tokenClaims); ok && token.Valid {
			return claims.Data, nil
		}
	}

	return nil, fmt.Errorf("token is invalid")
}
