package endpoints

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sebastiantraverso/alertas_a_corto_plazo_go/alertas-AC-api/handler"
)

func SetShortAlertsEndpoints(h *handler.ShortAlert) *mux.Router {
	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/weather/sort-alert/all", h.GetAllShortAlerts)

	return sm
}
