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

type Indigencies struct {
	ID             int64  `json:"ID"`
	ResidentID     int64  `json:"ResidentID"`
	DateCreated    string `json:"DateCreated"`
	DateUpdated    string `json:"DateUpdated"`
	Reason         string `json:"Reason"`
	ValidUntil     string `json:"ValidUntil"`
	IssuingOfficer string `json:"IssuingOfficer"`
	Remarks        string `json:"Remarks"`
}

func (b *BimsConfiguration) ReadIndigencies(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	IndigenciesData, err := database.ReadIndigencies(b.BIMSdb)
	if err != nil {
		log.Print(err)
		respondJSON(w, 400, nil)
		return
	}

	respondJSON(w, 200, IndigenciesData)
}

func (b *BimsConfiguration) DeleteIndigencies(w http.ResponseWriter, r *http.Request) {

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

	err = database.DeleteIndigencies(b.BIMSdb, req.ID)
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

func (b *BimsConfiguration) UpdateIndigencies(w http.ResponseWriter, r *http.Request) {

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

	req := &Indigencies{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, &UpdateResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	err = database.UpdateIndigencies(b.BIMSdb, req.ID, req.Reason, req.Remarks)
	if err != nil {
		respondJSON(w, 200, &UpdateResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	residentData, err := database.ReadResidentData(b.BIMSdb, req.ResidentID)
	if err != nil {
		respondJSON(w, 200, &UpdateResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	fullName := residentData.FirstName + " " + residentData.MiddleName + " " + residentData.LastName
	err = CreateIndigencyPDF(req.ResidentID, req.ID, fullName, residentData.Address, req.Reason)
	if err != nil {
		respondJSON(w, 200, &UpdateResponse{
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
