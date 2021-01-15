package webservice

import (
	"encoding/json"
	"github.com/nicolamassarenti/be-dashboard-bert-faqclass/src/usecases"
	"net/http"
)

// KnowledgeBaseInteractor defines the interface for the knowledge base
type KnowledgeBaseInteractor interface {
	AddFaq(usecases.Faq) error
	ChangeTrainingStatus(ID string, newStatus bool) error
	DeleteFaq(ID string) error
	Faq(ID string) (usecases.Faq, error)
	KnowledgeBase() ([]usecases.Faq, error)
	Update(ID string, faq usecases.Faq) error
}

// LanguagesInteractor defines the interface for the languages
type LanguagesInteractor interface {
	Languages() ([]usecases.Language, error)
}

// KeywordsInteractor defines the interface for the keywords
type KeywordsInteractor interface {
	Add(usecases.Keyword) error
	Delete(ID string) error
	Update(ID string, keyword usecases.Keyword) error
	Keywords() ([]usecases.Keyword, error)
}

// KB defines the object that represents the knowledge base
type KB struct {
	KB []FaqPresentation `json:"kb,omitempty"`
}

// FaqPresentation defines the faq for presentation
type FaqPresentation struct {
	ID           string `json:"id,omitempty"`
	MainQuestion string `json:"mainQuestion,omitempty"`
	Trained      bool   `json:"trained"`
}

// Keyword defines a keyword
type Keyword struct {
	ID          string `json:"id,omitempty"`
	DisplayText string `json:"keyword,omitempty"`
}

// Keywords defines the keywords
type Keywords struct {
	Keywords []Keyword `json:"keywords,omitempty"`
}

// Faq defines the faq detail for presentation
type Faq struct {
	MainQuestion string              			`json:"mainQuestion"`
	Answers      map[string][]string            `json:"answers"`
	Trained      bool                			`json:"trained"`
	Examples     map[string][]string 			`json:"examples"`
}

// WebserviceHandler it's the handler for REST api
type WebserviceHandler struct {
	KnowledgeBaseInteractor KnowledgeBaseInteractor
	LanguagesInteractor     LanguagesInteractor
	KeywordsInteractor      KeywordsInteractor
	Logger                  usecases.Logger
}

// checkID is the function that checks if an id has been passed as query param
func checkID(handler WebserviceHandler, res http.ResponseWriter, req *http.Request) bool{
	ids, ok := req.URL.Query()["id"]
	if !ok || len(ids) != 1 {
		if !ok {
			handler.Logger.Info("Error retrieving the ID. Returning BadRequest")
		} else if len(ids) == 0 {
			handler.Logger.Info("No ID as query params. Returning BadRequest")
		} else {
			handler.Logger.Info("More than one ID in query params. Returning BadRequest")
		}
		handler.Logger.Info("Returning response")
		res.WriteHeader(http.StatusBadRequest)
	}
	return ok
}

// Alive returns 200 OK is the webservice is alive
func (handler WebserviceHandler) Alive(res http.ResponseWriter, req *http.Request) {
	handler.Logger.Info("Received " + req.Method + " request at path: " + req.URL.Path)

	// Setting headers for CORS
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Authorization")
	if req.Method == http.MethodOptions {
		return
	}


	body, err := json.Marshal("I am alive")
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.Header().Add("Content-Type", "application/json")

	res.WriteHeader(200)
	_, err = res.Write(body)
	if err != nil {
		return
	}

	handler.Logger.Debug("Response set-up, returning the request")
	handler.Logger.Info("Returning response")
	return
}

