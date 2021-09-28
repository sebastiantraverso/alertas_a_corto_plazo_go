package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	data "github.com/sebastiantraverso/alertas_a_corto_plazo_go/alertas-AC-api/data"
)

func GetAllShortAlertsHandler(c *gin.Context) {
	alertasData := data.NewAlertas()
	// fetch alerts fron datastore
	response, err := alertasData.GetAllAlertsData()
	if err != nil {
		http.Error(c.Writer, fmt.Sprintf("Error: GetAllShortAlertsHandler - alertasData.GetAllAlertsData - %s", err), http.StatusInternalServerError)
	}

	c.IndentedJSON(http.StatusOK, response)
}
