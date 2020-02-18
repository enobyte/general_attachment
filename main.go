package main

import (
	"fmt"
	"general_attachment/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/upload_file", controllers.UploadFile).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "7878" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, RequestLogger(router)) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}

func RequestLogger(targetMux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		targetMux.ServeHTTP(w, r)

		// log request by who(IP address)
		requesterIP := r.RemoteAddr

		log.Printf(
			"%s\t\t%s\t\t%s\t\t%v",
			r.Method,
			r.RequestURI,
			requesterIP,
			time.Since(start),
		)
	})
}
