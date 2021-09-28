package data

import (
	"fmt"
)

type Alertas struct {
}

func NewAlertas() *Alertas {
	return &Alertas{}
}

func (a *Alertas) GetAllAlertsData() ([]AlertasData, error) {
	funcName := "GetAllAlertsData -"

	responseObj, err := a.getAllAlertsDataCommon()
	if err != nil {
		return []AlertasData{}, fmt.Errorf("%s - call getAllAlertsDataCommon - %s", funcName, err)
	}

	fmt.Println(responseObj)
	return a.validateNoAlerts(&responseObj)
}
