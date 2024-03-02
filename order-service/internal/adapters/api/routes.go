package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func InitRouter() error {
	router := mux.NewRouter()
	http.Handle("/", router)
	setupRoutes(router)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		return err
	}
	return nil
}

func setupRoutes(r *mux.Router) {
	//r.HandleFunc()
}
