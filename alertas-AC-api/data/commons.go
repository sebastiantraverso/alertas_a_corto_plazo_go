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
	fmt.Println(response.Body) // TODO: borrar

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []AlertasData{}, fmt.Errorf("%s GetAllAlertsData - ioutil.ReadAll - %s", funcName, err)
	}
	fmt.Println(string(responseData)) // TODO: borrar

	// var responseObj Response
	var responseObj []AlertasData
	json.Unmarshal(responseData, &responseObj)
	if err != nil {
		return []AlertasData{}, fmt.Errorf("%s GetAllAlertsData - unmarshal - %s", funcName, err)
	}

	fmt.Println(responseObj)
	return responseObj, nil
}
