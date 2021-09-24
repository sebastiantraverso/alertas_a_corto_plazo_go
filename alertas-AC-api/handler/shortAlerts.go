package handler

import (
	"fmt"
	"log"
	"net/http"

	data "github.com/sebastiantraverso/alertas_a_corto_plazo_go/alertas-AC-api/data"
)

type ShortAlert struct {
	l *log.Logger
}

// NewShortAlert creates a short alerts handler whit the logger
func NewShortAlert(l *log.Logger) *ShortAlert {
	return &ShortAlert{l}
}

func (s *ShortAlert) GetAllShortAlerts(rw http.ResponseWriter, r *http.Request) {
	s.l.Println("Handler GET GetAllShortAlerts")

	alertasData := data.NewAlertas()
	// fetch alerts fron datastore
	response, err := alertasData.GetAllAlerts()
	if err != nil {
		// return data.Response{}, fmt.Errorf("GetAllShortAlerts - alertasData.GetAllAlerts - Error: %s", err)
		http.Error(rw, fmt.Sprintf("Error: GetAllShortAlerts - alertasData.GetAllAlerts - %s", err), http.StatusInternalServerError)
	}

	err = response.ToJSON(rw)
	if err != nil {
		http.Error(rw, fmt.Sprintf("Error: GetAllShortAlerts - response.ToJSON - %s", err), http.StatusInternalServerError)
	}
	// return response, nil
	// serialize the list to json
}
