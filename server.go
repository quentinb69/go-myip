package main

import (
	"log"
	"net/http"
)

// set handler for and start listening
func LoadServer() error {

	r := http.NewServeMux()
	r.HandleFunc("/", ShowHomeHandler)
	r.HandleFunc("/health", HealthHandler)

	log.Print("Starting server on port 8000")
	return http.ListenAndServe(":8000", r)
}

// health handler
func HealthHandler(w http.ResponseWriter, _ *http.Request) {
	log.Print("Print Health")
	w.Write([]byte("OK"))
}

// default handler
func ShowHomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Print IP")
	ip := GetIp(r)
	w.Write([]byte(ip))
}
