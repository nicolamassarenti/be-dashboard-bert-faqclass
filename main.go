package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
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
	kbInteractor.FaqRepository = interfaces.NewFaqDBHandler(dbHandler, "KnowledgeBase")

	langInteractor := new(usecases.LanguageInteractor)
	langInteractor.LanguageRepository = interfaces.NewLanguagesDBHandler(dbHandler, "Languages")

	logger := infrastructure.NewLogger()
	kbInteractor.Logger = logger
	langInteractor.Logger = logger

	webserviceHandler := interfaces.WebserviceHandler{}
	webserviceHandler.KnowledgeBaseInteractor = kbInteractor
	webserviceHandler.LanguagesInteractor = langInteractor
	webserviceHandler.Logger = logger

	logger.Info("Handlers created")

	// Routes
	rtr := mux.NewRouter()
	rtr.HandleFunc("/alive", webserviceHandler.Alive)

	rtr.HandleFunc("/api/lang", webserviceHandler.GetAllLanguages).
		Methods(http.MethodGet)

	rtr.HandleFunc("/api/kb", webserviceHandler.KnowledgeBase).
		Methods(http.MethodGet)

	rtr.HandleFunc("/api/faq", webserviceHandler.Faq).
		Methods(http.MethodGet)

	rtr.HandleFunc("/api/faq/{id}", webserviceHandler.AddFaq).
		Methods(http.MethodPost)

	rtr.HandleFunc("/api/faq/{id}", webserviceHandler.DeleteFaq).
		Methods(http.MethodDelete)

	rtr.HandleFunc("/api/training/faq", webserviceHandler.ChangeTrainingStatus).
		Methods(http.MethodPut).
		Queries("toTrain")

	http.Handle("/", rtr)

	logger.Info("Router and handler function created")

	// Server
	logger.Info("Server starting at port " + port)
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS()(rtr)))

}
