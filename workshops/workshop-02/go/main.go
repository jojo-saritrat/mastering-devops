package main

import (
	"log"
	"net/http"
	"os"
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello this is code from Go!"))
}

func healthCheckReadinessHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if _, err := os.Stat("/tmp/ready"); err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("200 OK, it's ready!"))
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("503 Service unavailable, it's not ready!"))
	}
}

func healthCheckLivenessHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if _, err := os.Stat("/tmp/ready"); err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("200 OK, it lives!"))
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("503 Service unavailable, it doesn't live!"))
	}
}

func main() {
	http.HandleFunc("/", greetHandler)
	http.HandleFunc("/health/readiness", healthCheckReadinessHandler)
	http.HandleFunc("/health/liveness", healthCheckLivenessHandler)

	log.Println("Server is running on port 80")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
