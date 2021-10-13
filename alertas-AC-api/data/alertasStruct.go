package data

type ResposneAlertasCortoPlazo struct {
	Version string        `json:"version,omitempty"`
	Data    []AlertasData `json:"data,omitempty"`
}

type AlertasData struct {
	ID                 string            `json:"_id,omitempty"`
	IDAlert            int64             `json:"idAlert,omitempty"`
	NReport            interface{}       `json:"nReport,omitempty"`
	Type               string            `json:"type,omitempty"`
	Title              string            `json:"title,omitempty"`
	Status             interface{}       `json:"status,omitempty"`
	Date               string            `json:"date,omitempty"`
	Hour               string            `json:"hour,omitempty"`
	Description        interface{}       `json:"description,omitempty"`
	Zones              map[string]string `json:"zones,omitempty"`
	Severity           string            `json:"severity,omitempty"`
	Polygon            string            `json:"polygon,omitempty"`
	Urls               []URL             `json:"urls,omitempty"`
	RegionTopesNubosos string            `json:"region_topes_nubosos,omitempty"`
	Partial            string            `json:"partial,omitempty"`
	Update             interface{}       `json:"update,omitempty"`
}

type URL struct {
	Value       *string `json:"value"`
	Description string  `json:"description"`
}
