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

type Referrals struct {
	ID                 int64  `json:"ID"`
	ResidentID         int64  `json:"ResidentID"`
	DateCreated        string `json:"DateCreated"`
	DateUpdated        string `json:"DateUpdated"`
	HCGGGNumber        string `json:"HCGGGNumber"`
	PhilHealthID       string `json:"PhilHealthID"`
	PhilHealthCategory string `json:"PhilHealthCategory"`
	ReasonForReferral  string `json:"ReasonForReferral"`
	ValidUntil         string `json:"ValidUntil"`
	IssuingOfficer     string `json:"IssuingOfficer"`
	Remarks            string `json:"Remarks"`
}

func (b *BimsConfiguration) ReadReferrals(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	ReferralsData, err := database.ReadReferrals(b.BIMSdb)
	if err != nil {
		log.Print(err)
		respondJSON(w, 400, nil)
		return
	}

	respondJSON(w, 200, ReferralsData)
}

func (b *BimsConfiguration) DeleteReferrals(w http.ResponseWriter, r *http.Request) {

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

	err = database.DeleteReferrals(b.BIMSdb, req.ID)
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

func (b *BimsConfiguration) UpdateReferrals(w http.ResponseWriter, r *http.Request) {

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

	req := &Referrals{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, &UpdateResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	err = database.UpdateReferrals(b.BIMSdb, req.ID, req.ResidentID, req.HCGGGNumber, req.PhilHealthID, req.PhilHealthCategory, req.ReasonForReferral,
		req.IssuingOfficer, req.Remarks)
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
