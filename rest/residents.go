package rest

import (
	"bims/database"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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

type UpdateResidentsRequest struct {
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
	ResidentID            int64  `json:"ResidentID"`
	DocumentDateCreated   string `json:"DocumentDateCreated"`
	DocumentDateUpdated   string `json:"DocumentDateUpdated"`
	Purpose               string `json:"Purpose"`
	ValidUntil            string `json:"ValidUntil"`
	Remarks               string `json:"Remarks"`
	DocumentStatus        string `json:"DocumentStatus"`
	HCGGGNumber           string `json:"HCGGGNumber"`
	PhilHealthID          string `json:"PhilHealthID"`
	PhilHealthCategory    string `json:"PhilHealthCategory"`
	ReasonForReferral     string `json:"ReasonForReferral"`
	Reason                string `json:"Reason"`
	CedulaNo              string `json:"CedulaNo"`
	PrecintNo             string `json:"PrecintNo"`
	DocumentType          string `json:"DocumentType"`
	DocumentID            int64  `json:"DocumentID"`
	IssuingOfficer        string `json:"IssuingOfficer"`
}

type DeleteRequestResident struct {
	ID           int64  `json:"ID"`
	DocumentID   int64  `json:"DocumentID"`
	DocumentType string `json:"DocumentType"`
}

func (b *BimsConfiguration) ReadResidentsXL(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	ResidentsData, err := database.ReadResidentsXL(b.BIMSdb)
	if err != nil {
		log.Print(err)
		respondJSON(w, 400, nil)
		return
	}

	respondJSON(w, 200, ResidentsData)
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

	req := &DeleteRequestResident{}

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

	if req.DocumentType == "Barangay Indigency" {
		err = database.DeleteIndigencies(b.BIMSdb, req.DocumentID)
		if err != nil {
			respondJSON(w, 200, &DeleteResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}
	} else if req.DocumentType == "Barangay Clearance" {
		err = database.DeleteClearance(b.BIMSdb, req.DocumentID)
		if err != nil {
			respondJSON(w, 200, &DeleteResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}
	} else if req.DocumentType == "Referral Slip" {
		err = database.DeleteReferrals(b.BIMSdb, req.DocumentID)
		if err != nil {
			respondJSON(w, 200, &DeleteResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}
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

	req := &UpdateResidentsRequest{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, &UpdateResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	err = database.UpdateResidents(b.BIMSdb, req.ID, req.LastName, req.FirstName, req.MiddleName, req.Address, req.BirthDate, req.BirthPlace,
		req.Gender, req.CivilStatus, req.ContactNumber, req.GuardianName, req.GurdianContactNumbers)
	if err != nil {
		respondJSON(w, 200, &UpdateResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if req.DocumentType == "Barangay Indigency" {
		err = database.UpdateIndigencies(b.BIMSdb, req.DocumentID, req.Reason, req.Remarks, req.DocumentStatus)
		if err != nil {
			respondJSON(w, 200, &UpdateResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		err = CreateIndigencyPDF(req.ID, req.DocumentID, req.FirstName+" "+req.MiddleName+" "+req.LastName, req.Address, req.Reason)
		if err != nil {
			respondJSON(w, 200, &UpdateResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}

	} else if req.DocumentType == "Barangay Clearance" {
		err = database.UpdateClearance(b.BIMSdb, req.DocumentID, req.Remarks, req.LastName, req.FirstName, req.MiddleName, req.Purpose, req.CedulaNo, req.PrecintNo, req.DocumentStatus)
		if err != nil {
			respondJSON(w, 200, &UpdateResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		currentTimeClearance := time.Now()
		formattedTime := currentTimeClearance.Format("January 2, 2006")

		parsedDate, err := time.Parse("01/02/2006", req.BirthDate)
		if err != nil {
			respondJSON(w, 200, &UpdateResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		Birthday := parsedDate.Format("January 2, 2006")

		err = CreateClearancePDF(req.ResidentID, req.DocumentID, formattedTime, Birthday, req.BirthPlace, req.FirstName+" "+req.MiddleName+" "+req.LastName, req.Address, req.CivilStatus,
			req.Purpose, req.CedulaNo, req.PrecintNo, req.ValidUntil)
		if err != nil {
			respondJSON(w, 200, &UpdateResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}

	} else if req.DocumentType == "Referral Slip" {
		err = database.UpdateReferrals(b.BIMSdb, req.DocumentID, req.HCGGGNumber, req.PhilHealthID, req.PhilHealthCategory, req.ReasonForReferral, req.Remarks, req.DocumentStatus)
		if err != nil {
			respondJSON(w, 200, &UpdateResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		err = CreateReferralsPDF(req.ID, req.DocumentID, req.LastName, req.MiddleName, req.FirstName, req.Address, req.ContactNumber, req.GuardianName, req.GurdianContactNumbers, req.ReasonForReferral,
			req.HCGGGNumber, req.PhilHealthID, req.PhilHealthCategory, req.Gender, req.BirthDate, req.CivilStatus, req.BirthPlace)
		if err != nil {
			respondJSON(w, 200, &UpdateResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}
	}

	respondJSON(w, 200, &UpdateResponse{
		Success: true,
		Message: "",
	})
}
