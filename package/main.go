package main

import (
	"github.com/nicolamassarenti/be-dashboard-bert-faqclass/src/infrastructure/db"
	"github.com/nicolamassarenti/be-dashboard-bert-faqclass/src/infrastructure/logging"
	"github.com/nicolamassarenti/be-dashboard-bert-faqclass/src/interfaces/repositories"
	"github.com/nicolamassarenti/be-dashboard-bert-faqclass/src/interfaces/webservice"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/nicolamassarenti/be-dashboard-bert-faqclass/src/usecases"
)

func main() {
	port := os.Getenv("PORT")
	projectID := "PROJECT_ID"

	// Handlers, interfaces and implementation
	dbHandler := db.Handler(projectID)
	kbInteractor := new(usecases.KnowledgeBaseInteractor)
	kbInteractor.Repository = repositories.NewFaqDBHandler(dbHandler, "KnowledgeBase")

	langInteractor := new(usecases.LanguageInteractor)
	langInteractor.Repository = repositories.NewLanguagesDBHandler(dbHandler, "Languages")

	keywordsInteractor := new(usecases.KeywordsInteractor)
	keywordsInteractor.Repository = repositories.NewKeywordsDBHandler(dbHandler, "keywords")

	logger := logging.NewLogger()
	kbInteractor.Logger = logger
	langInteractor.Logger = logger

	webserviceHandler := webservice.WebserviceHandler{}
	webserviceHandler.KnowledgeBaseInteractor = kbInteractor
	webserviceHandler.LanguagesInteractor = langInteractor
	webserviceHandler.Logger = logger

	logger.Info("Handlers created")

	// Routes
	rtr := mux.NewRouter()
	rtr.HandleFunc("/alive", webserviceHandler.Alive).
		Methods(http.MethodGet)

	// #################################################################################################################
	// # Languages #####################################################################################################
	// #################################################################################################################

	rtr.HandleFunc("/api/lang", webserviceHandler.GetAllLanguages).
		Methods(http.MethodGet, http.MethodOptions)

	// #################################################################################################################
	// # FAQs ##########################################################################################################
	// #################################################################################################################

	rtr.HandleFunc("/api/kb", webserviceHandler.KnowledgeBase).
		Methods(http.MethodGet, http.MethodOptions)

	rtr.HandleFunc("/api/faq", webserviceHandler.Faq).
		Methods(http.MethodGet, http.MethodOptions)

	rtr.HandleFunc("/api/faq", webserviceHandler.AddFaq).
		Methods(http.MethodPost, http.MethodOptions)

	rtr.HandleFunc("/api/faq", webserviceHandler.UpdateFaq).
		Methods(http.MethodPut, http.MethodOptions)

	rtr.HandleFunc("/api/faq", webserviceHandler.DeleteFaq).
		Methods(http.MethodDelete, http.MethodOptions)

	rtr.HandleFunc("/api/training/faq", webserviceHandler.ChangeTrainingStatus).
		Methods(http.MethodPut, http.MethodOptions)

	// #################################################################################################################
	// # Keywords ######################################################################################################
	// #################################################################################################################
	rtr.HandleFunc("/api/keyword", webserviceHandler.AddKeyword).
		Methods(http.MethodPost, http.MethodOptions)

	rtr.HandleFunc("/api/keyword", webserviceHandler.DeleteKeyword).
		Methods(http.MethodDelete, http.MethodOptions)

	http.Handle("/", rtr)

	logger.Info("Router and handler function created")

	// Server
	logger.Info("Server starting at port " + port)
	log.Fatal(http.ListenAndServe(":"+port, rtr))

}
