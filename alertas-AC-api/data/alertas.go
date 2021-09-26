package data

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Alertas struct {
}

func NewAlertas() *Alertas {
	return &Alertas{}
}

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    comment, err := UnmarshalComment(bytes)
//    bytes, err = comment.Marshal()

func UnmarshalResponse(data []byte) (Response, error) {
	var r Response
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Response) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func (r *Response) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(r)
}

type Response struct {
	Ok   bool          `json:"ok"`
	Data []AlertasData `json:"data"`
}

type AlertasData struct {
	ID                 string            `json:"_id"`
	IDAlert            int64             `json:"idAlert"`
	NReport            interface{}       `json:"nReport"`
	Type               string            `json:"type"`
	Title              string            `json:"title"`
	Status             interface{}       `json:"status"`
	Date               string            `json:"date"`
	Hour               string            `json:"hour"`
	Description        interface{}       `json:"description"`
	Zones              map[string]string `json:"zones"`
	Severity           string            `json:"severity"`
	Polygon            string            `json:"polygon"`
	Urls               []URL             `json:"urls"`
	RegionTopesNubosos string            `json:"region_topes_nubosos"`
	Partial            string            `json:"partial"`
	Update             interface{}       `json:"update"`
}

type URL struct {
	Value       *string `json:"value"`
	Description string  `json:"description"`
}

func (a *Alertas) GetAllAlerts() (Response, error) {
	response, err := http.Get("https://ws.smn.gob.ar/alerts/type/AC")
	if err != nil {
		return Response{}, fmt.Errorf("Error: getAllAlerts - http.Get - %s", err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Response{}, fmt.Errorf("Error: getAllAlerts - ioutil.ReadAll - %s", err)
	}

	// responseObj, err := UnmarshalResponse(responseData)
	var responseObj Response
	json.Unmarshal(responseData, &responseObj)
	if err != nil {
		return Response{}, fmt.Errorf("Error: getAllAlerts - UnmarshalResponse - %s", err)
	}

	fmt.Print(responseObj)
	return responseObj, nil
}
