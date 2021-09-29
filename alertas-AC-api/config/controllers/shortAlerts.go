package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sebastiantraverso/alertas_a_corto_plazo_go/alertas-AC-api/handler"
)

func SetEndpoints(router *gin.Engine) {
	router.GET("/weather/sort-alert/all", handler.GetAllShortAlertsHandler)
}
