package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	controllers "github.com/sebastiantraverso/alertas_a_corto_plazo_go/alertas-AC-api/config/controllers"
	config "github.com/sebastiantraverso/alertas_a_corto_plazo_go/alertas-AC-api/config/environment"
)

func main() {

	router := gin.Default()
	controllers.SetEndpoints(router)

	printBuilVersion()

	addSrv := getRunningPort()
	router.Run(addSrv)
}

func getRunningPort() string {
	addPort, err := config.GetPort()
	if err != nil {
		fmt.Printf("Variables del server (addPort): Error - %s", err)
		os.Exit(1)
	}
	return fmt.Sprintf(":%d", addPort)
}

func printBuilVersion() {
	buildVersion, err := config.GetBuildVersion()
	if err != nil {
		fmt.Printf("Variables del server (buildVersion): Error - %s", err)
	}
	fmt.Printf("Running version: %s \n", buildVersion)
}
