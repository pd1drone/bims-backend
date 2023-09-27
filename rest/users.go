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

type DeleteRequest struct {
	ID         int64 `json:"ID"`
	ResidentID int64 `json:"ResidentID"`
}

type DeleteResponse struct {
	Success bool   `json:"Success"`
	Message string `json:"Message"`
}

type Users struct {
	ID           int64  `json:"ID"`
	FullName     string `json:"FullName"`
	FirstName    string `json:"FirstName"`
	MiddleName   string `json:"MiddleName"`
	LastName     string `json:"LastName"`
	PositionID   int64  `json:"PositionID"`
	PositionName string `json:"PositionName"`
	Email        string `json:"Email"`
	Username     string `json:"Username"`
	IsAdmin      bool   `json:"IsAdmin"`
	ProfileLink  string `json:"ProfileLink"`
}

type UpdateResponse struct {
	Success bool   `json:"Success"`
	Message string `json:"Message"`
}

type CreateResponse struct {
	Success bool   `json:"Success"`
	Message string `json:"Message"`
}

func (b *BimsConfiguration) ReadUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	UserData, err := database.ReadUsers(b.BIMSdb)
	if err != nil {
		log.Print(err)
		respondJSON(w, 400, nil)
		return
	}

	respondJSON(w, 200, UserData)
}

func (b *BimsConfiguration) DeleteUsers(w http.ResponseWriter, r *http.Request) {

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

	err = database.DeleteUser(b.BIMSdb, req.ID)
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

func (b *BimsConfiguration) UpdateUsers(w http.ResponseWriter, r *http.Request) {

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

	req := &Users{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, &UpdateResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	err = database.UpdateUser(b.BIMSdb, req.ID, req.FullName, req.FirstName, req.MiddleName, req.LastName,
		req.PositionID, req.Email, req.Username)
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

func (b *BimsConfiguration) CreateUsers(w http.ResponseWriter, r *http.Request) {

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

	req := &Users{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, &CreateResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	generatedPassword, err := GeneratePassword()
	if err != nil {
		respondJSON(w, 500, &CreateResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	md5HashPass := MD5HashPassword(generatedPassword)

	err = database.CreateUser(b.BIMSdb, req.FullName, req.FirstName, req.MiddleName, req.LastName, req.PositionID, req.Email,
		req.Username, md5HashPass, "")
	if err != nil {
		respondJSON(w, 200, &CreateResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	err = b.SendPasswordViaEmail(req.Email, req.FullName, req.Username, generatedPassword)
	if err != nil {
		respondJSON(w, 400, &CreateResponse{
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
