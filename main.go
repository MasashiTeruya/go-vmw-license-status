package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// LicenseStatus contains if the license is enabled or not.
type LicenseStatus struct {
	Enabled bool `json:"enabled"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		status := &LicenseStatus{Enabled: true}
		json.NewEncoder(w).Encode(status)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
