package main

import (
	"fmt"
	"net/http"

	"crudmongomux2/api/upload_api"

	"github.com/gorilla/handlers"

	"crudmongomux2/api/productapi"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/cooking/findall", productapi.FindAll).Methods("GET")
	r.HandleFunc("/api/cooking/find/{id}", productapi.Find).Methods("GET")
	r.HandleFunc("/api/cooking/create", productapi.Create).Methods("POST")
	r.HandleFunc("/api/cooking/update", productapi.Update).Methods("PUT")
	r.HandleFunc("/api/cooking/{id}", productapi.Delete).Methods("DELETE")

	r.HandleFunc("/api/upload", upload_api.UploadFile).Methods("POST")

	//testing rest api
	r.HandleFunc("/api/status/find/{id}", productapi.Getdata).Methods("GET")

	err := http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r))

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("listening 3000.....")
	}
}
