package main

import (
	"github.com/gorilla/mux"

	"net/http"

	"github.com/NicolaMassarenti/be-dashboard-bert-faqclass/interfaces"
)

const (
	PORT = 8080
)

func main() {

	webserviceHandler := interfaces.WebserviceHandler{}

	rtr := mux.NewRouter()
	rtr.HandleFunc("/faq", webserviceHandler.KnowledgeBase).
		Methods(http.MethodGet)

	rtr.HandleFunc("/faq/{id}", webserviceHandler.Faq).
		Methods(http.MethodGet)

	rtr.HandleFunc("/faq/{id}", webserviceHandler.ChangeTrainingStatus).
		Methods(http.MethodPut).
		Queries("toTrain")

	rtr.HandleFunc("/faq/{id}", webserviceHandler.AddFaq).
		Methods(http.MethodPost)

	rtr.HandleFunc("/faq/{id}", webserviceHandler.DeleteFaq).
		Methods(http.MethodDelete)

	http.Handle("/", rtr)
	http.ListenAndServe(PORT, nil)
}
