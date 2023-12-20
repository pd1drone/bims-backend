package rest

import (
	"bims/database"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Clearance struct {
	ID                 int64     `json:"ID"`
	ResidentID         int64     `json:"ResidentID"`
	DateCreated        string    `json:"DateCreated"`
	DateUpdated        string    `json:"DateUpdated"`
	ValidUntil         string    `json:"ValidUntil"`
	IssuingOfficer     string    `json:"IssuingOfficer"`
	Remarks            string    `json:"Remarks"`
	ResidentLastName   string    `json:"ResidentLastName"`
	ResidentFirstName  string    `json:"ResidentFirstName"`
	ResidentMiddleName string    `json:"ResidentMiddleName"`
	Purpose            string    `json:"Purpose"`
	CedulaNo           string    `json:"CedulaNo"`
	PrecintNo          string    `json:"PrecintNo"`
	DocumentStatus     string    `json:"DocumentStatus"`
	ResidentData       Residents `json:"ResidentData"`
}

func (b *BimsConfiguration) ReadClearanceXL(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	ClearanceData, err := database.ReadClearanceXL(b.BIMSdb)
	if err != nil {
		log.Print(err)
		respondJSON(w, 400, nil)
		return
	}

	respondJSON(w, 200, ClearanceData)
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

	fmt.Println(req)

	err = database.UpdateClearance(b.BIMSdb, req.ID, req.Remarks, req.ResidentLastName,
		req.ResidentFirstName, req.ResidentMiddleName, req.Purpose, req.CedulaNo, req.PrecintNo, req.DocumentStatus)
	if err != nil {
		respondJSON(w, 200, &UpdateResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

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
	// residentData, err := database.ReadResidentData(b.BIMSdb, req.ResidentID)
	// if err != nil {
	// 	respondJSON(w, 200, &UpdateResponse{
	// 		Success: false,
	// 		Message: err.Error(),
	// 	})
	// 	return
	// }

	currentTimeClearance := time.Now()
	formattedTime := currentTimeClearance.Format("January 2, 2006")

	parsedDate, err := time.Parse("01/02/2006", req.ResidentData.BirthDate)
	if err != nil {
		respondJSON(w, 200, &UpdateResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	Birthday := parsedDate.Format("January 2, 2006")
	fullName := req.ResidentFirstName + " " + req.ResidentMiddleName + " " + req.ResidentLastName

	err = CreateClearancePDF(req.ResidentID, req.ID, formattedTime, Birthday, req.ResidentData.BirthPlace, fullName, req.ResidentData.Address, req.ResidentData.CivilStatus, req.Purpose, req.CedulaNo, req.PrecintNo, req.ValidUntil)
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
