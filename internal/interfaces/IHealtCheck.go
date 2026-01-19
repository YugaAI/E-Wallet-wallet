package interfaces

import "github.com/gin-gonic/gin"

type IHealtCheckServices interface {
	HealtCheckServices() (string, error)
}
type IHealthCheckRepository interface {
}
type IHealthcheckAPI interface {
	HealthcheckHandlerHTTP(c *gin.Context)
}
