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
	ID             int64     `json:"ID"`
	ResidentID     int64     `json:"ResidentID"`
	DateCreated    string    `json:"DateCreated"`
	DateUpdated    string    `json:"DateUpdated"`
	Purpose        string    `json:"Purpose"`
	ValidUntil     string    `json:"ValidUntil"`
	IssuingOfficer string    `json:"IssuingOfficer"`
	Remarks        string    `json:"Remarks"`
	DocumentStatus string    `json:"DocumentStatus"`
	ResidentData   Residents `json:"ResidentData"`
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

func (b *BimsConfiguration) ReadIndigenciesXL(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	IndigenciesData, err := database.ReadIndigenciesXL(b.BIMSdb)
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

	err = database.DeleteResidents(b.BIMSdb, req.ResidentID)
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

	err = database.UpdateIndigencies(b.BIMSdb, req.ID, req.Purpose, req.Remarks, req.DocumentStatus)
	if err != nil {
		respondJSON(w, 200, &UpdateResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	// residentData, err := database.ReadResidentData(b.BIMSdb, req.ResidentID)
	// if err != nil {
	// 	respondJSON(w, 200, &UpdateResponse{
	// 		Success: false,
	// 		Message: err.Error(),
	// 	})
	// 	return
	// }
	err = database.UpdateResidents(b.BIMSdb, req.ResidentData.ID, req.ResidentData.LastName, req.ResidentData.FirstName,
		req.ResidentData.MiddleName, req.ResidentData.Address, req.ResidentData.BirthDate, req.ResidentData.BirthPlace, req.ResidentData.Gender,
		req.ResidentData.CivilStatus, req.ResidentData.ContactNumber, req.ResidentData.GuardianName, req.ResidentData.GurdianContactNumbers)
	if err != nil {
		respondJSON(w, 200, &UpdateResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	fullName := req.ResidentData.FirstName + " " + req.ResidentData.MiddleName + " " + req.ResidentData.LastName
	err = CreateIndigencyPDF(req.ResidentID, req.ID, fullName, req.ResidentData.Address, req.Purpose)
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
