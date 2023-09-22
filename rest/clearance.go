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

type Clearance struct {
	ID                 int64  `json:"ID"`
	ResidentID         int64  `json:"ResidentID"`
	DateCreated        string `json:"DateCreated"`
	DateUpdated        string `json:"DateUpdated"`
	ValidUntil         string `json:"ValidUntil"`
	IssuingOfficer     string `json:"IssuingOfficer"`
	Remarks            string `json:"Remarks"`
	ResidentLastName   string `json:"ResidentLastName"`
	ResidentFirstName  string `json:"ResidentFirstName"`
	ResidentMiddleName string `json:"ResidentMiddleName"`
	Purpose            string `json:"Purpose"`
}

func (b *BimsConfiguration) ReadClearance(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	ClearanceData, err := database.ReadClearance(b.BIMSdb)
	if err != nil {
		log.Print(err)
		respondJSON(w, 400, nil)
		return
	}

	respondJSON(w, 200, ClearanceData)
}

func (b *BimsConfiguration) DeleteClearance(w http.ResponseWriter, r *http.Request) {

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

	err = database.DeleteClearance(b.BIMSdb, req.ID)
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

func (b *BimsConfiguration) UpdateClearance(w http.ResponseWriter, r *http.Request) {

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

	req := &Clearance{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, &UpdateResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	err = database.UpdateClearance(b.BIMSdb, req.ID, req.ResidentID, req.IssuingOfficer, req.Remarks, req.ResidentLastName,
		req.ResidentFirstName, req.ResidentMiddleName, req.Purpose)
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
