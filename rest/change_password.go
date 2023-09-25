package rest

import (
	"bims/database"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type ChangePasswordRequest struct {
	Username    string `json:"Username"`
	Password    string `json:"Password"`
	NewPassword string `json:"NewPassword"`
	ID          int64  `json:"ID"`
}

type ChangePasswordResponse struct {
	Successful bool   `json:"Successful"`
	Message    string `json:"Message"`
}

func (b *BimsConfiguration) ChangePassword(w http.ResponseWriter, r *http.Request) {

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

	req := &ChangePasswordRequest{}
	response := &ChangePasswordResponse{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		fmt.Println(err)
		response.Successful = false
		response.Message = fmt.Sprintf(err.Error())
		respondJSON(w, 400, response)
		return
	}

	fmt.Println(req)

	err = database.ChangePassword(b.BIMSdb, req.ID, req.Username, req.Password, req.NewPassword)
	if err != nil {
		fmt.Println(err)
		response.Successful = false
		response.Message = fmt.Sprintf(err.Error())
		respondJSON(w, 400, response)
		return
	}

	response.Successful = true

	respondJSON(w, 200, response)
}
