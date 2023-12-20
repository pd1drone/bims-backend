package rest

import (
	"bims/database"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type NewRequest struct {
	DocuTitle           string `json:"docuTitle"`
	LastName            string `json:"lastName"`
	FirstName           string `json:"firstName"`
	MiddleName          string `json:"middleName"`
	PhilHealthNumber    string `json:"philHealthNumber"`
	PhilHealthCategory  string `json:"philHealthCategory"`
	Address             string `json:"address"`
	HealthCardGGGNumber string `json:"healthCardGGGNumber"`
	TelNum              string `json:"telNum"`
	BirthPlace          string `json:"birthPlace"`
	Gender              string `json:"gender"`
	ParentName          string `json:"parentName"`
	CedulaNo            string `json:"cedulaNo"`
	PrecintNo           string `json:"precintNo"`
	CivilStatus         string `json:"civilStatus"`
	ParentContactNumber string `json:"parentContactNumber"`
	Purpose             string `json:"purpose"`
	ReasonForReferral   string `json:"reasonForReferral"`
	Remarks             string `json:"remarks"`
	Birthdate           string `json:"birthDate"`
	IssuingOfficer      string `json:"issuingOfficer"`
	ValidUntil          string `json:"validUntil"`
}

type NewResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (b *BimsConfiguration) New(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondJSON(w, 500, &NewResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	// Restore request body after reading
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	req := &NewRequest{}
	fmt.Println(req)

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, &NewResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	residentID, err := database.CreateResident(b.BIMSdb, req.LastName, req.FirstName, req.MiddleName, req.Address, req.Birthdate, req.BirthPlace,
		req.Gender, req.CivilStatus, req.TelNum, req.ParentName, req.ParentContactNumber, req.IssuingOfficer, req.DocuTitle)
	if err != nil {
		fmt.Println(err)
		respondJSON(w, 200, &NewResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	fmt.Print("RESIDENT ID: ")
	fmt.Println(residentID)
	currentTime := time.Now()
	newTime := currentTime.Add(720 * time.Hour)
	ValidUntil := newTime.Format("2006-01-02 03:04 PM")

	if req.DocuTitle == "Barangay Indigency" {

		documentIndigencyID, err := database.CreateIndigencies(b.BIMSdb, residentID, req.Purpose, ValidUntil, req.IssuingOfficer, req.Remarks)
		if err != nil {
			respondJSON(w, 200, &NewResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		fullName := req.FirstName + " " + req.MiddleName + " " + req.LastName
		err = CreateIndigencyPDF(residentID, documentIndigencyID, fullName, req.Address, req.Remarks)
		if err != nil {
			respondJSON(w, 200, &NewResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		err = database.UpdateDocumentID(b.BIMSdb, documentIndigencyID, residentID)
		if err != nil {
			respondJSON(w, 200, &NewResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		respondJSON(w, 200, &NewResponse{
			Success: true,
			Message: "",
		})
		return
	}

	if req.DocuTitle == "Barangay Clearance" {

		documentClearanceID, err := database.CreateClearance(b.BIMSdb, residentID, ValidUntil, req.IssuingOfficer, req.Remarks, req.LastName, req.FirstName,
			req.MiddleName, req.Purpose, req.CedulaNo, req.PrecintNo)
		if err != nil {
			respondJSON(w, 200, &NewResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		currentTimeClearance := time.Now()
		formattedTime := currentTimeClearance.Format("January 2, 2006")

		parsedDate, err := time.Parse("01/02/2006", req.Birthdate)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		Birthday := parsedDate.Format("January 2, 2006")
		fullName := req.FirstName + " " + req.MiddleName + " " + req.LastName

		err = CreateClearancePDF(residentID, documentClearanceID, formattedTime, Birthday, req.BirthPlace, fullName, req.Address, req.CivilStatus, req.Purpose, req.CedulaNo, req.PrecintNo, req.ValidUntil)
		if err != nil {
			respondJSON(w, 200, &NewResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		err = database.UpdateDocumentID(b.BIMSdb, documentClearanceID, residentID)
		if err != nil {
			respondJSON(w, 200, &NewResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		respondJSON(w, 200, &NewResponse{
			Success: true,
			Message: "",
		})
		return
	}

	if req.DocuTitle == "Referral Slip" {
		documentReferralsID, err := database.CreateReferrals(b.BIMSdb, residentID, req.HealthCardGGGNumber, req.PhilHealthNumber, req.PhilHealthCategory,
			req.ReasonForReferral, ValidUntil, req.IssuingOfficer, req.Remarks)
		if err != nil {
			respondJSON(w, 200, &NewResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}
		err = CreateReferralsPDF(residentID, documentReferralsID, req.LastName, req.MiddleName, req.FirstName, req.Address, req.TelNum,
			req.ParentName, req.ParentContactNumber, req.ReasonForReferral, req.HealthCardGGGNumber, req.PhilHealthNumber, req.PhilHealthCategory,
			req.Gender, req.Birthdate, req.CivilStatus, req.BirthPlace)
		if err != nil {
			respondJSON(w, 200, &NewResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		err = database.UpdateDocumentID(b.BIMSdb, documentReferralsID, residentID)
		if err != nil {
			respondJSON(w, 200, &NewResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		respondJSON(w, 200, &NewResponse{
			Success: true,
			Message: "",
		})
		return
	}

}
