package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/nicolamassarenti/be-dashboard-bert-faqclass/src/infrastructure"
	"github.com/nicolamassarenti/be-dashboard-bert-faqclass/src/interfaces"
	"github.com/nicolamassarenti/be-dashboard-bert-faqclass/src/usecases"
)

func main() {
	port := os.Getenv("PORT")
	projectID := "PROJECT_ID"

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

	http.Handle("/", rtr)

	logger.Info("Router and handler function created")

	// Server
	logger.Info("Server starting at port " + port)
	log.Fatal(http.ListenAndServe(":"+port, rtr))

}
