package interfaces

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/NicolaMassarenti/be-dashboard-bert-faqclass/src/usecases"
	"github.com/gorilla/mux"
)

// KnowledgeBaseInteractor is the interactor that links the webservice to the usecases
type KnowledgeBaseInteractor interface {
	KnowledgeBase() ([]usecases.Faq, error)
	Faq(ID string) (usecases.Faq, error)
	ChangeTrainingStatus(ID string, newStatus bool) error
	AddFaq(usecases.Faq) error
	DeleteFaq(ID string) error
}

// Faq contains the data that define a F.A.Q, in the format required by the UI
type Faq struct {
	ID           string              `json:"id,omitempty"`
	MainQuestion string              `json:"mainQuestion,omitempty"`
	Answers      []Answer            `json:"answers,omitempty"`
	Trained      bool                `json:"trained,omitempty"`
	Examples     map[string][]string `json:"examples,omitempty"`
}

// Answer contains the answer in a language
type Answer struct {
	lang   string
	answer string
}

// faqOverview represents the overview of a Faq
type faqOverview struct {
	ID           string `json:"id,omitempty"`
	MainQuestion string `json:"mainQuestion,omitempty"`
	Trained      bool   `trained:"id,omitempty"`
}

// WebserviceHandler it's the handler for REST api
type WebserviceHandler struct {
	KnowledgeBaseInteractor KnowledgeBaseInteractor
}

// KnowledgeBase is the handler function that returns the kb
func (handler WebserviceHandler) KnowledgeBase(res http.ResponseWriter, req *http.Request) {
	kb, err := handler.KnowledgeBaseInteractor.KnowledgeBase()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	var body []byte
	if body, err = json.Marshal(kb); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.Header().Add("Content-Type", "application/json")

	res.WriteHeader(200)
	res.Write(body)
	return
}

// Faq is the handler function that returns a Faq
func (handler WebserviceHandler) Faq(res http.ResponseWriter, req *http.Request) {
	// Retrieving the ID from the url
	var id string
	id = mux.Vars(req)["id"]

	// Retrieving the Faq
	usecaseFaq, err := handler.KnowledgeBaseInteractor.Faq(id)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Transforming the data for the UI
	faq := usecaseFaqToWebserviceFaq(usecaseFaq)

	// Preparing the response
	var body []byte
	if body, err = json.Marshal(faq); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.Header().Add("Content-Type", "application/json")

	res.WriteHeader(200)
	res.Write(body)
	return
}

// ChangeTrainingStatus is the handler function that returns a Faq
func (handler WebserviceHandler) ChangeTrainingStatus(res http.ResponseWriter, req *http.Request) {
	// Retrieving the ID from the url
	var id string
	var toTrain bool
	var err error

	// Retrieving the ID
	id = mux.Vars(req)["id"]

	// Retrieving the "toTrain" from the query string parameters
	toTrain, err = strconv.ParseBool(mux.Vars(req)["toTrain"])
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Retrieving the Faq
	err = handler.KnowledgeBaseInteractor.ChangeTrainingStatus(id, toTrain)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Preparing the response
	res.WriteHeader(200)
	return
}

// AddFaq is the handler function that adds a new Faq
func (handler WebserviceHandler) AddFaq(res http.ResponseWriter, req *http.Request) {

	// Retrieving the ID from the url
	var id string
	var err error
	var newFaq Faq

	// Retrieving the ID
	id = mux.Vars(req)["id"]

	// Parsing the request body
	err = json.NewDecoder(req.Body).Decode(&newFaq)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Checking if the ID of the query string is the same of the body
	if id != newFaq.ID {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	// Data transformation
	usecasesFaq := webserviceFaqToUsecaseFaq(newFaq)

	// Adding the new Faq
	err = handler.KnowledgeBaseInteractor.AddFaq(usecasesFaq)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Preparing the response
	res.WriteHeader(200)
	return
}

// DeleteFaq is the handler function that adds a new Faq
func (handler WebserviceHandler) DeleteFaq(res http.ResponseWriter, req *http.Request) {

	// Retrieving the ID from the url
	var id string
	var err error

	// Retrieving the ID
	id = mux.Vars(req)["id"]

	// Deleting the new Faq
	err = handler.KnowledgeBaseInteractor.DeleteFaq(id)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Preparing the response
	res.WriteHeader(200)
	return
}