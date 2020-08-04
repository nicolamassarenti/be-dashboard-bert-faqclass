package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/NicolaMassarenti/be-dashboard-bert-faqclass/src/infrastructure"
	"github.com/NicolaMassarenti/be-dashboard-bert-faqclass/src/interfaces"
	"github.com/NicolaMassarenti/be-dashboard-bert-faqclass/src/usecases"
)

const (
	port     = "8080"
	authPath = "./auth/bert-faqclass-a96dec925432.json"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file")
		return
	}

	dbHandler := infrastructure.NewFirestoreHandler(authPath)
	kbInteractor := new(usecases.KnowledgeBaseInteractor)
	kbInteractor.FaqRepository = interfaces.NewFaqDBHandler(dbHandler, "Faq")

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
	http.ListenAndServe(port, nil)
}
