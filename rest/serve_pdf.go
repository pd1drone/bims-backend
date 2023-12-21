package rest

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
)

func (b *BimsConfiguration) ServeClearancePDF(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	residentID := chi.URLParam(r, "residentID")
	documentID := chi.URLParam(r, "documentID")
	filename := chi.URLParam(r, "filename")

	// Use the retrieved userID and filename as needed
	fmt.Println("residentID:", residentID)
	fmt.Println("documentID:", documentID)
	fmt.Println("filename:", filename)

	// Construct the file path
	filePath := fmt.Sprintf("/root/bims-backend/files/clearances/%s/%s/clearances_%s_%s.pdf", residentID, documentID, residentID, documentID)

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Set the appropriate Content-Type header
	contentType := mime.TypeByExtension(filepath.Ext(filePath))
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	w.Header().Set("Content-Type", contentType)

	// Copy the file contents to the response writer
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}
}

func (b *BimsConfiguration) ServeIndigenciesPDF(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	residentID := chi.URLParam(r, "residentID")
	documentID := chi.URLParam(r, "documentID")
	filename := chi.URLParam(r, "filename")

	// Use the retrieved userID and filename as needed
	fmt.Println("residentID:", residentID)
	fmt.Println("documentID:", documentID)
	fmt.Println("filename:", filename)

	// Construct the file path
	filePath := fmt.Sprintf("/root/bims-backend/files/indigencies/%s/%s/indigencies_%s_%s.pdf", residentID, documentID, residentID, documentID)

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Set the appropriate Content-Type header
	contentType := mime.TypeByExtension(filepath.Ext(filePath))
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	w.Header().Set("Content-Type", contentType)

	// Copy the file contents to the response writer
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}
}

func (b *BimsConfiguration) ServeReferralsPDF(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	residentID := chi.URLParam(r, "residentID")
	documentID := chi.URLParam(r, "documentID")
	filename := chi.URLParam(r, "filename")

	// Use the retrieved userID and filename as needed
	fmt.Println("residentID:", residentID)
	fmt.Println("documentID:", documentID)
	fmt.Println("filename:", filename)

	// Construct the file path
	filePath := fmt.Sprintf("/root/bims-backend/files/referrals/%s/%s/referrals_%s_%s.pdf", residentID, documentID, residentID, documentID)

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Set the appropriate Content-Type header
	contentType := mime.TypeByExtension(filepath.Ext(filePath))
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	w.Header().Set("Content-Type", contentType)

	// Copy the file contents to the response writer
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}
}

func (b *BimsConfiguration) ServeBDRRMCpdf(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	ID := chi.URLParam(r, "ID")
	filename := chi.URLParam(r, "filename")

	// Use the retrieved userID and filename as needed
	fmt.Println("ID:", ID)
	fmt.Println("filename:", filename)

	// Construct the file path
	filePath := fmt.Sprintf("/root/bims-backend/files/bdrrmc/%s/bdrrmc_%s.pdf", ID, ID)

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Set the appropriate Content-Type header
	contentType := mime.TypeByExtension(filepath.Ext(filePath))
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	w.Header().Set("Content-Type", contentType)

	// Copy the file contents to the response writer
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}
}

func (b *BimsConfiguration) ServeIncidentsPDF(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	ID := chi.URLParam(r, "ID")
	filename := chi.URLParam(r, "filename")

	// Use the retrieved userID and filename as needed
	fmt.Println("ID:", ID)
	fmt.Println("filename:", filename)

	// Construct the file path
	filePath := fmt.Sprintf("/root/bims-backend/files/incidents/%s/incidents_%s.pdf", ID, ID)

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Set the appropriate Content-Type header
	contentType := mime.TypeByExtension(filepath.Ext(filePath))
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	w.Header().Set("Content-Type", contentType)

	// Copy the file contents to the response writer
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}
}
