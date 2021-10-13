package data

import (
	"fmt"
)

type Alertas struct {
}

func NewAlertas() *Alertas {
	return &Alertas{}
}

func (a *Alertas) GetAllAlertsData() (ResposneAlertasCortoPlazo, error) {
	funcName := "GetAllAlertsData -"

	responseObj, err := a.getAllAlertsDataCommon()
	if err != nil {
		return ResposneAlertasCortoPlazo{}, fmt.Errorf("%s call getAllAlertsDataCommon - %s", funcName, err)
	}

	buildVersion, err := a.getBuildVersion()
	if err != nil {
		return ResposneAlertasCortoPlazo{}, fmt.Errorf("%s call getBuildVersion - %s", funcName, err)
	}

	var response ResposneAlertasCortoPlazo
	response.Data = responseObj
	response.Version = buildVersion

	return response, nil
}
