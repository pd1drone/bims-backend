package rest

import (
	"bims/database"
	"fmt"
	"io"
	"log"
	"mime"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/go-chi/chi"
)

func (b *BimsConfiguration) UploadUserProfile(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	// Parse our multipart form, 32 << 20 specifies a maximum
	// upload of 32 MB files.
	r.ParseMultipartForm(32 << 20)

	file, handler, err := r.FormFile("uploadUserProfile")
	if err != nil {
		w.Write([]byte("Error in uploading file!"))
		respondJSON(w, 400, nil)
		return
	}
	defer file.Close()

	UserID := r.FormValue("UserID")
	fmt.Println(UserID)

	err = createDirectoryIfNotExist("/root/bims-backend/files/" + UserID)
	if err != nil {
		w.Write([]byte("Error in uploading file!"))
		respondJSON(w, 500, nil)
		return
	}

	filePath := fmt.Sprintf("/root/bims-backend/files/%s/%s", UserID, handler.Filename)

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		w.Write([]byte("Error in uploading file!"))
		respondJSON(w, 500, nil)
		return
	}
	// Save the file to disk with the provided file path
	err = os.WriteFile(filePath, fileBytes, 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	intUserID, err := strconv.Atoi(UserID)
	if err != nil {
		w.Write([]byte("Error in uploading file!"))
		respondJSON(w, 500, nil)
		return
	}

	ipaddr := GetOutboundIP()

	link := fmt.Sprintf("http://" + ipaddr.String() + ":8085/files/" + UserID + "/" + handler.Filename)

	err = database.UploadUserPicture(b.BIMSdb, int64(intUserID), link)
	if err != nil {
		w.Write([]byte("Error in uploading file!"))
		respondJSON(w, 500, nil)
		return
	}

	// Return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")

	// Set the response headers
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Write the JavaScript code that performs the redirection
	redirectScript := `
			<script>
				window.location.href = 'http://` + b.FrontEndIP + `:` + b.PortNumber + `/dashboard';
			</script>
		`

	// Write the response containing the redirection script
	w.Write([]byte(redirectScript))

}

func (b *BimsConfiguration) ServeFile(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	userID := chi.URLParam(r, "userID")
	filename := chi.URLParam(r, "filename")

	// Use the retrieved userID and filename as needed
	fmt.Println("userID:", userID)
	fmt.Println("filename:", filename)

	// Construct the file path
	filePath := fmt.Sprintf("/root/bims-backend/files/%s/%s", userID, filename)

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

func createDirectoryIfNotExist(directoryPath string) error {
	// Check if the directory already exists
	if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
		// Directory does not exist, create it
		err := os.MkdirAll(directoryPath, 0755)
		if err != nil {
			return fmt.Errorf("error creating directory: %w", err)
		}
		fmt.Println("Directory created:", directoryPath)
	} else if err != nil {
		return fmt.Errorf("error checking directory: %w", err)
	} else {
		fmt.Println("Directory already exists:", directoryPath)
	}

	return nil
}

// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
