package data

import (
	"encoding/json"
	"fmt"

	"io/ioutil"
	"net/http"
)

func (a *Alertas) getAllAlertsDataCommon() (Response, error) {
	funcName := "getAllAlertsDataCommon -"

	response, err := http.Get("https://ws.smn.gob.ar/alerts/type/AC")
	//response, err := http.NewRequest("GET", "https://ws.smn.gob.ar/alerts/type/AC", nil)
	if err != nil {
		return Response{}, fmt.Errorf("%s http.Get - %s", funcName, err)
	}
	fmt.Println(response.Body)

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Response{}, fmt.Errorf("%s GetAllAlertsData - ioutil.ReadAll - %s", funcName, err)
	}

	var responseObj Response
	json.Unmarshal(responseData, &responseObj)
	if err != nil {
		return Response{}, fmt.Errorf("%s GetAllAlertsData - unmarshal - %s", funcName, err)
	}

	return responseObj, nil
}

func (a *Alertas) validateNoAlerts(resp *Response) ([]AlertasData, error) {
	funcName := "validateNoAlerts -"

	if !resp.Ok {
		s := []byte(`[ { "novalue": "No hay alertas" } ]`)
		var newAlerData []AlertasData
		err := json.Unmarshal(s, &newAlerData)
		if err != nil {
			return []AlertasData{}, fmt.Errorf("%s Unmarshal - %s", funcName, err)
		}
		return newAlerData, nil
	}
	return resp.Data, nil
}
