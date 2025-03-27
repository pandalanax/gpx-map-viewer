package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const gpxDir = "./gpx_tracks"

func main() {
	// Serve index.html at root
	http.HandleFunc("/", serveIndex)

	// Serve GPX files from the gpx_tracks directory
	http.HandleFunc("/gpx/", serveGPX)

	// List available GPX files
	http.HandleFunc("/gpx-files", listGPXFiles)

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Serve index.html
func serveIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

// Serve GPX file from directory
func serveGPX(w http.ResponseWriter, r *http.Request) {
	fileName := strings.TrimPrefix(r.URL.Path, "/gpx/")
	filePath := filepath.Join(gpxDir, fileName)

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.NotFound(w, r)
		return
	}

	http.ServeFile(w, r, filePath)
}

// List all GPX files in the directory
func listGPXFiles(w http.ResponseWriter, r *http.Request) {
	files, err := filepath.Glob(filepath.Join(gpxDir, "*.gpx"))
	if err != nil {
		http.Error(w, "Error reading GPX files", http.StatusInternalServerError)
		return
	}

	// Extract file names and send them as a JSON response
	var fileNames []string
	for _, file := range files {
		_, fileName := filepath.Split(file)
		fileNames = append(fileNames, fileName)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("[\"" + strings.Join(fileNames, "\",\"") + "\"]"))
}
