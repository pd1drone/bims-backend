package rest

import (
	"bims/database"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Incidents struct {
	ID                int64  `json:"ID"`
	DateCreated       string `json:"DateCreated"`
	DateUpdated       string `json:"DateUpdated"`
	CompliantFullName string `json:"CompliantFullName"`
	Respondent        string `json:"Respondent"`
	IncidentStatus    string `json:"IncidentStatus"`
	IncidentDateTime  string `json:"IncidentDateTime"`
	IncidentLocation  string `json:"IncidentLocation"`
	IncidentNarration string `json:"IncidentNarration"`
	IssuingOfficer    string `json:"IssuingOfficer"`
}

func (b *BimsConfiguration) ReadIncidents(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	incidentData, err := database.ReadIncidents(b.BIMSdb)
	if err != nil {
		log.Print(err)
		respondJSON(w, 400, nil)
		return
	}

	respondJSON(w, 200, incidentData)
}

func (b *BimsConfiguration) DeleteIncidents(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondJSON(w, 500, &DeleteResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	// Restore request body after reading
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	req := &DeleteRequest{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, &DeleteResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	err = database.DeleteIncidents(b.BIMSdb, req.ID)
	if err != nil {
		respondJSON(w, 200, &DeleteResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	respondJSON(w, 200, &DeleteResponse{
		Success: true,
		Message: "",
	})
}

func (b *BimsConfiguration) UpdateIncidents(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondJSON(w, 500, &UpdateResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	// Restore request body after reading
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	req := &Incidents{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, &UpdateResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	err = database.UpdateIncidents(b.BIMSdb, req.ID, req.CompliantFullName, req.Respondent, req.IncidentStatus, req.IncidentDateTime, req.IncidentLocation, req.IncidentNarration, req.IssuingOfficer)
	if err != nil {
		respondJSON(w, 200, &UpdateResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	err = CreateIncidentsPDF(req.ID, req.CompliantFullName, req.Respondent, req.IncidentDateTime, req.IncidentLocation, req.IncidentNarration, req.IssuingOfficer)
	if err != nil {
		respondJSON(w, 200, &CreateResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	respondJSON(w, 200, &UpdateResponse{
		Success: true,
		Message: "",
	})
}

func (b *BimsConfiguration) CreateIncidents(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondJSON(w, 500, nil)
		return
	}

	// Restore request body after reading
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	req := &Incidents{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, &CreateResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	id, err := database.CreateIncidents(b.BIMSdb, req.CompliantFullName, req.Respondent, req.IncidentStatus, req.IncidentDateTime, req.IncidentLocation, req.IncidentNarration, req.IssuingOfficer)
	if err != nil {
		respondJSON(w, 200, &CreateResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	err = CreateIncidentsPDF(id, req.CompliantFullName, req.Respondent, req.IncidentDateTime, req.IncidentLocation, req.IncidentNarration, req.IssuingOfficer)
	if err != nil {
		respondJSON(w, 200, &CreateResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	respondJSON(w, 200, &CreateResponse{
		Success: true,
		Message: "",
	})
}
