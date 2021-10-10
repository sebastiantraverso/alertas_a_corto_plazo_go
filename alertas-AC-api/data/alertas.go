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
		return ResposneAlertasCortoPlazo{}, fmt.Errorf("%s - call getAllAlertsDataCommon - %s", funcName, err)
	}

	buildVersion, err := a.getBuildVersion()
	if err != nil {
		return ResposneAlertasCortoPlazo{}, fmt.Errorf("%s - call getBuildVersion - %s", funcName, err)
	}

	response := ResposneAlertasCortoPlazo{
		data:    responseObj,
		version: buildVersion,
	}

	// return responseObj, nil
	return response, nil
}

// func setUpResponseCortoPlazo(temp []data.AlertasData) data.ResposneAlertasCortoPlazo {
// 	response := ResposneAlertasCortoPlazo{
// 		data: ,
// 		version: "222"
// 	}
// }
