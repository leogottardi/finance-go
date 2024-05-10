package router

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Configure() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {
		// Get body
		body, err := io.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}

		bodyString := string(body)
		log.Println(bodyString)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(bodyString))
	}).Methods(http.MethodPost)
	return router
}
