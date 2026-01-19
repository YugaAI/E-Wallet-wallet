package cmd

import (
	"E-Wallet-wallet/external"
	"E-Wallet-wallet/helpers"
	"E-Wallet-wallet/internal/api"
	"E-Wallet-wallet/internal/interfaces"
	"E-Wallet-wallet/internal/repository"
	"E-Wallet-wallet/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	d := dependencyInject()

	r := gin.Default()

	r.GET("/health", d.HealthcheckAPI.HealthcheckHandlerHTTP)

	walletV1 := r.Group("/wallet/v1")
	walletV1.POST("/", d.WalletAPI.Create)
	walletV1.PUT("/balance/credit", d.MiddlewareValidateToken, d.WalletAPI.CreditBalance)
	walletV1.PUT("/balance/debit", d.MiddlewareValidateToken, d.WalletAPI.DebitBalance)
	walletV1.GET("/balance", d.MiddlewareValidateToken, d.WalletAPI.GetBalance)
	walletV1.GET("/history", d.MiddlewareValidateToken, d.WalletAPI.GetWalletHistory)

	err := r.Run(":" + helpers.GetEnv("PORT", ""))
	if err != nil {
		log.Fatal(err)
	}
}

type Dependency struct {
	HealthcheckAPI interfaces.IHealthcheckAPI
	WalletAPI      interfaces.IWalletAPI
	External       interfaces.IExternal
}

func dependencyInject() Dependency {
	healthcheckSvc := &services.HealthCheck{}
	healthcheckAPI := &api.HealthCheckService{
		HealthCheckServices: healthcheckSvc,
	}

	walletRepo := &repository.WalletRepo{
		DB: helpers.DB,
	}

	walletSvc := &services.WalletService{
		WalletRepo: walletRepo,
	}
	walletAPI := &api.WalletAPI{
		WalletService: walletSvc,
	}

	external := &external.External{}

	return Dependency{
		HealthcheckAPI: healthcheckAPI,
		WalletAPI:      walletAPI,
		External:       external,
	}
}
