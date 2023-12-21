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

type BDRRMC struct {
	ID              int64  `json:"ID"`
	TypeOfRecord    string `json:"TypeOfRecord"`
	PartiesInvolved string `json:"PartiesInvolved"`
	DateTime        string `json:"DateTime"`
	Location        string `json:"Location"`
	RecordDetails   string `json:"RecordDetails"`
	IssuingOfficer  string `json:"IssuingOfficer"`
}

func (b *BimsConfiguration) ReadBDRRMC(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	bdrrmcData, err := database.ReadBDRRMC(b.BIMSdb)
	if err != nil {
		log.Print(err)
		respondJSON(w, 400, nil)
		return
	}

	respondJSON(w, 200, bdrrmcData)
}

func (b *BimsConfiguration) DeleteBDRRMC(w http.ResponseWriter, r *http.Request) {

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

	err = database.DeleteBDRRMC(b.BIMSdb, req.ID)
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

func (b *BimsConfiguration) UpdateBDRRMC(w http.ResponseWriter, r *http.Request) {

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

	req := &BDRRMC{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, &UpdateResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	err = database.UpdateBDRRMC(b.BIMSdb, req.ID, req.TypeOfRecord, req.PartiesInvolved, req.DateTime, req.Location, req.RecordDetails, req.IssuingOfficer)
	if err != nil {
		respondJSON(w, 200, &UpdateResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	err = CreateBDRRMCpdf(req.ID, req.TypeOfRecord, req.PartiesInvolved, req.DateTime, req.Location, req.RecordDetails, req.IssuingOfficer)
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

func (b *BimsConfiguration) CreateBDRRMC(w http.ResponseWriter, r *http.Request) {

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

	req := &BDRRMC{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, &CreateResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	id, err := database.CreateBDRRMC(b.BIMSdb, req.TypeOfRecord, req.PartiesInvolved, req.DateTime, req.Location, req.RecordDetails, req.IssuingOfficer)
	if err != nil {
		respondJSON(w, 200, &CreateResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	err = CreateBDRRMCpdf(id, req.TypeOfRecord, req.PartiesInvolved, req.DateTime, req.Location, req.RecordDetails, req.IssuingOfficer)
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
