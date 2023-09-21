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

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Success  bool                    `json:"success"`
	Message  string                  `json:"message"`
	Response *database.LoginResponse `json:"response"`
}

func (b *BimsConfiguration) Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondJSON(w, 500, &LoginResponse{
			Success:  false,
			Message:  err.Error(),
			Response: nil,
		})
		return
	}

	// Restore request body after reading
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	req := &LoginRequest{}
	fmt.Println(req)

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, &LoginResponse{
			Success:  false,
			Message:  err.Error(),
			Response: nil,
		})
		return
	}

	LoginData, err := database.Login(b.BIMSdb, req.Username, req.Password)
	if err != nil {
		respondJSON(w, 200, &LoginResponse{
			Success:  false,
			Message:  err.Error(),
			Response: nil,
		})
		return
	}

	respondJSON(w, 200, &LoginResponse{
		Success:  true,
		Message:  "",
		Response: LoginData,
	})
}
