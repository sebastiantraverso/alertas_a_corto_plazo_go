package environment

import (
	"errors"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Error al leer el archivo de configuración %s", err)
	}
	log.Println("Archivo de configuraciòn leido correctamente")
}

func GetPort() (int, error) {
	value, ok := viper.Get("PORT").(int)
	if !ok {
		return 0, errors.New("GetPort: Error al recuperar la variable PORT")
	}
	return value, nil
}

func getServer() (string, error) {
	value, ok := viper.Get("SERVER").(string)
	if !ok {
		return "", errors.New("getServer: Error al recuperar la variable SERVER")
	}
	return value, nil
}

func GetServerAddress() (string, error) {
	port, err := GetPort()
	if err != nil {
		return "", fmt.Errorf("GetServerAddress: Error - %s", err)
	}

	server, err := getServer()
	if err != nil {
		return "", fmt.Errorf("GetServerAddress: Error - %s", err)
	}

	addr := fmt.Sprintf("%s:%d", server, port)
	return addr, nil
}
