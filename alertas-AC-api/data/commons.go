package data

import (
	"encoding/json"
	"fmt"

	"io/ioutil"
	"net/http"
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

	// var responseObj Response
	var responseObj []AlertasData
	json.Unmarshal(responseData, &responseObj)
	if err != nil {
		return []AlertasData{}, fmt.Errorf("%s GetAllAlertsData - unmarshal - %s", funcName, err)
	}

	return responseObj, nil
}
