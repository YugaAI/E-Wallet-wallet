package cmd

import (
	"E-Wallet-wallet/helpers"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (d *Dependency) MiddlewareValidateToken(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		helpers.SendResponse(c, http.StatusUnauthorized, "unauthorized", nil)
		c.Abort()
		return
	}

	const prefix = "Bearer "
	token := authHeader
	if strings.HasPrefix(authHeader, prefix) {
		token = strings.TrimSpace(authHeader[len(prefix):])
	}
	log.Printf("RAW Authorization header: [%q]", authHeader)
	log.Printf("TOKEN AFTER STRIP: [%q]", token)
	log.Printf("TOKEN DOT COUNT: %d", strings.Count(token, "."))

	if token == "" {
		helpers.SendResponse(c, http.StatusUnauthorized, "unauthorized", nil)
		c.Abort()
		return
	}

	tokenData, err := d.External.ValidateToken(c.Request.Context(), token)
	if err != nil {
		log.Printf("validate token failed:", err)
		helpers.SendResponse(c, http.StatusUnauthorized, "unauthorized", nil)
		c.Abort()
		return
	}

	c.Set("token", tokenData)
	c.Next()
}
