package data

import (
	"encoding/json"
	"fmt"

	"io/ioutil"
	"net/http"

	config "github.com/sebastiantraverso/alertas_a_corto_plazo_go/alertas-AC-api/config/environment"
)

func (a *Alertas) getAllAlertsDataCommon() ([]AlertasData, error) {
	funcName := "getAllAlertsDataCommon -"

	response, err := http.Get("https://ws.smn.gob.ar/alerts/type/AC")
	if err != nil {
		return []AlertasData{}, fmt.Errorf("%s http.Get - %s", funcName, err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []AlertasData{}, fmt.Errorf("%s GetAllAlertsData - ioutil.ReadAll - %s", funcName, err)
	}

	var responseObj []AlertasData
	json.Unmarshal(responseData, &responseObj)
	if err != nil {
		return []AlertasData{}, fmt.Errorf("%s GetAllAlertsData - unmarshal - %s", funcName, err)
	}

	return responseObj, nil
}

func (a *Alertas) getBuildVersion() (string, error) {
	funcName := "getBuildVersion -"

	buildVersion, err := config.GetBuildVersion()
	if err != nil {
		return "", fmt.Errorf("%s GetBuildVersion - %s", funcName, err)
	}

	return buildVersion, nil
}
