package cmd

import (
	"E-Wallet-wallet/helpers"
	"E-Wallet-wallet/internal/api"
	"E-Wallet-wallet/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServerHTTP() {

	HealthCheckSvc := &services.HealthCheck{}
	healthcheckAPI := api.HealthCheckService{
		HealthCheckServices: HealthCheckSvc,
	}
	r := gin.Default()

	r.GET("/health", healthcheckAPI.HealthCheckHandlerHTTP)

	err := r.Run(":" + helpers.GetEnv("PORT", "8080"))
	if err != nil {
		log.Fatal(err)
	}
}
