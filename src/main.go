package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/NicolaMassarenti/be-dashboard-bert-faqclass/src/infrastructure"
	"github.com/NicolaMassarenti/be-dashboard-bert-faqclass/src/interfaces"
	"github.com/NicolaMassarenti/be-dashboard-bert-faqclass/src/usecases"
)

func main() {
	// Loading ENV variables
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file")
		return
	}

	port := os.Getenv("PORT")
	projectID := os.Getenv("PROJECT_ID")

	// Handlers, interfaces and implementation
	dbHandler := infrastructure.NewFirestoreHandler(projectID)
	kbInteractor := new(usecases.KnowledgeBaseInteractor)
	kbInteractor.FaqRepository = interfaces.NewFaqDBHandler(dbHandler, "Faq")

	logger := infrastructure.NewLogger()
	kbInteractor.Logger = logger

	webserviceHandler := interfaces.WebserviceHandler{}
	webserviceHandler.KnowledgeBaseInteractor = kbInteractor
	webserviceHandler.Logger = logger

	logger.Info("Handlers created")

	// Routes
	rtr := mux.NewRouter()
	rtr.HandleFunc("/alive", webserviceHandler.Alive)

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

	logger.Info("Router and handler function created")

	// Server
	logger.Info("Server starting at port " + port)
	log.Fatal(http.ListenAndServe(port, nil))

}
