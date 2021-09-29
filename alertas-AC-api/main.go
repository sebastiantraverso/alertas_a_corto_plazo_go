package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	controllers "github.com/sebastiantraverso/alertas_a_corto_plazo_go/alertas-AC-api/config/controllers"
	config "github.com/sebastiantraverso/alertas_a_corto_plazo_go/alertas-AC-api/config/environment"
)

func main() {

	router := gin.Default()
	controllers.SetEndpoints(router)

	addSrv, err := config.GetServerAddress()
	if err != nil {
		fmt.Printf("Variables del server (addSrv): Error - %s", err)
	}

	router.Run(addSrv)
}
