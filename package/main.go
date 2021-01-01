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
	projectID := "bert-faqclass"

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

	rtr.HandleFunc("/api/lang", webserviceHandler.GetAllLanguages).
		Methods(http.MethodGet, http.MethodOptions)

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

	http.Handle("/", rtr)

	logger.Info("Router and handler function created")

	rtr.Use(mux.CORSMethodMiddleware(rtr))
	// Server
	logger.Info("Server starting at port " + port)
	log.Fatal(http.ListenAndServe(":"+port, rtr))

}
