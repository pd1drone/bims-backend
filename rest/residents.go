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

type Residents struct {
	ID                    int64  `json:"ID"`
	DateCreated           string `json:"DateCreated"`
	DateUpdated           string `json:"DateUpdated"`
	LastName              string `json:"LastName"`
	FirstName             string `json:"FirstName"`
	MiddleName            string `json:"MiddleName"`
	Address               string `json:"Address"`
	BirthDate             string `json:"BirthDate"`
	BirthPlace            string `json:"BirthPlace"`
	Gender                string `json:"Gender"`
	CivilStatus           string `json:"CivilStatus"`
	ContactNumber         string `json:"ContactNumber"`
	GuardianName          string `json:"GuardianName"`
	GurdianContactNumbers string `json:"GurdianContactNumbers"`
	Religion              string `json:"Religion"`
	Occupation            string `json:"Occupation"`
	IssuingOfficer        string `json:"IssuingOfficer"`
}

func (b *BimsConfiguration) ReadResidents(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	ResidentsData, err := database.ReadResidents(b.BIMSdb)
	if err != nil {
		log.Print(err)
		respondJSON(w, 400, nil)
		return
	}

	respondJSON(w, 200, ResidentsData)
}

func (b *BimsConfiguration) DeleteResidents(w http.ResponseWriter, r *http.Request) {

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

	err = database.DeleteResidents(b.BIMSdb, req.ID)
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

func (b *BimsConfiguration) UpdateResidents(w http.ResponseWriter, r *http.Request) {

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

	req := &Residents{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, &UpdateResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	err = database.UpdateResidents(b.BIMSdb, req.ID, req.LastName, req.FirstName, req.MiddleName, req.Address, req.BirthDate, req.BirthPlace,
		req.Gender, req.CivilStatus, req.ContactNumber, req.GuardianName, req.GurdianContactNumbers, req.Religion, req.Occupation)
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
