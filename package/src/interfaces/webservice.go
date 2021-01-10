package interfaces

import (
	"encoding/json"
	"github.com/nicolamassarenti/be-dashboard-bert-faqclass/src/usecases"
	"net/http"
	"strconv"
)

// KnowledgeBaseInteractor is the interactor that links the webservice to the usecases
type KnowledgeBaseInteractor interface {
	AddFaq(usecases.Faq) error
	ChangeTrainingStatus(ID string, newStatus bool) error
	DeleteFaq(ID string) error
	Faq(ID string) (usecases.Faq, error)
	KnowledgeBase() ([]usecases.Faq, error)
	Update(ID string, faq usecases.Faq) error
}

// LanguagesInteractor is the interactor that links the webservice to the usecases
type LanguagesInteractor interface {
	GetAllLanguages() ([]usecases.Language, error)
}

// KeywordsInteractor is the interactor that links the webservice to the usecases
type KeywordsInteractor interface {
	Add(usecases.Keyword) error
	Delete(ID string) error
	Update(ID string, keyword usecases.Keyword) error
}

// KB is the struct that contains the preview of all the KB
type KB struct {
	KB []FaqPreview `json:"kb,omitempty"`
}

// FaqPreview contains the preview of a faq
type FaqPreview struct {
	ID           string `json:"id,omitempty"`
	MainQuestion string `json:"mainQuestion,omitempty"`
	Trained      bool   `json:"trained"`
}

// Keyword is a keyword
type Keyword struct {
	ID string
	Name string
}

// Faq contains the data that define a F.A.Q, in the format required by the UI
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
	KeywordsInteractor 		KeywordsInteractor
	Logger                  usecases.Logger
}

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
		res.WriteHeader(http.StatusBadRequest)
	}
	return ok
}

// Alive returns 200 OK
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
	res.Write(body)

	handler.Logger.Debug("Response set-up, returning the request")
	handler.Logger.Info("Returning response")
	return
}

// AddFaq is the handler function that adds a new Faq
func (handler WebserviceHandler) AddFaq(res http.ResponseWriter, req *http.Request) {
	handler.Logger.Info("Received " + req.Method + " request at path: " + req.URL.Path)

	// Setting headers for CORS
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Authorization")
	if req.Method == http.MethodOptions {
		return
	}

	var err error
	var newFaq Faq

	// Parsing the request body
	err = json.NewDecoder(req.Body).Decode(&newFaq)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
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
	handler.Logger.Info("Returning response")
	return
}

// ChangeTrainingStatus is the handler function that returns a Faq
func (handler WebserviceHandler) ChangeTrainingStatus(res http.ResponseWriter, req *http.Request) {
	handler.Logger.Info("Received " + req.Method + " request at path: " + req.URL.Path)

	// Setting headers for CORS
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Authorization")
	if req.Method == http.MethodOptions {
		return
	}

	// Retrieving the ID from the url
	var id string
	var toTrain bool
	var err error

	// Retrieving the ID and toTrain
	ok := checkID(handler, res, req)
	if !ok {
		return
	}
	id = req.URL.Query().Get("id")

	toTrainString, ok := req.URL.Query()["toTrain"]
	if !ok || len(toTrainString) != 1 {
		if !ok {
			handler.Logger.Info("Error retrieving param toTrain. Returning BadRequest")
		} else if len(toTrainString) == 0 {
			handler.Logger.Info("No ID as query params. Returning BadRequest")
		} else {
			handler.Logger.Info("More than one ID in query params. Returning BadRequest")
		}
		res.WriteHeader(http.StatusBadRequest)
	}
	toTrain, err = strconv.ParseBool(req.URL.Query().Get("toTrain"))
	if err != nil {
		handler.Logger.Info("Error transforming toTrain from string to bool. Returning Internal Server Error")
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	handler.Logger.Info("ID: " + id + "\ttoTrain: " + strconv.FormatBool(toTrain))

	// Retrieving the Faq
	err = handler.KnowledgeBaseInteractor.ChangeTrainingStatus(id, toTrain)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Preparing the response
	res.WriteHeader(200)
	handler.Logger.Info("Returning response")
	return
}

// DeleteFaq is the handler function that adds a new Faq
func (handler WebserviceHandler) DeleteFaq(res http.ResponseWriter, req *http.Request) {
	handler.Logger.Info("Received " + req.Method + " request at path: " + req.URL.Path)

	// Setting headers for CORS
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Authorization")
	if req.Method == http.MethodOptions {
		return
	}

	// Retrieving the ID from the url
	var id string
	var err error

	// Retrieving the ID
	id = req.URL.Query().Get("id")
	ok := checkID(handler, res, req)
	if !ok {
		return
	}
	handler.Logger.Info("ID: " + id)

	// Deleting the new Faq
	err = handler.KnowledgeBaseInteractor.DeleteFaq(id)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Preparing the response
	res.WriteHeader(200)
	handler.Logger.Info("Returning response")
	return
}

// GetAllLanguages returns all the languages
func (handler WebserviceHandler) GetAllLanguages(res http.ResponseWriter, req *http.Request) {
	handler.Logger.Info("Received " + req.Method + " request at path: " + req.URL.Path)

	// Setting headers for CORS
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Authorization")
	if req.Method == http.MethodOptions {
		return
	}

	languages, err := handler.LanguagesInteractor.GetAllLanguages()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	languagesMap := languagesToFrontEnd(languages)

	var body []byte
	if body, err = json.Marshal(languagesMap); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Header().Add("Content-Type", "application/json")

	res.WriteHeader(200)
	res.Write(body)
	handler.Logger.Info("Returning response")

	return
}

// Faq is the handler function that returns a Faq
func (handler WebserviceHandler) Faq(res http.ResponseWriter, req *http.Request) {
	handler.Logger.Info("Received " + req.Method + " request at path: " + req.URL.Path)

	// Setting headers for CORS
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Authorization")
	if req.Method == http.MethodOptions {
		return
	}

	// Retrieving the ID from the url
	var id string
	ok := checkID(handler, res, req)
	if !ok {
		return
	}
	ids, ok := req.URL.Query()["id"]
	id = ids[0]

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
	handler.Logger.Info("Returning response")
	return
}

// KnowledgeBase is the handler function that returns the kb
func (handler WebserviceHandler) KnowledgeBase(res http.ResponseWriter, req *http.Request) {
	handler.Logger.Info("Received " + req.Method + " request at path: " + req.URL.Path)

	// Setting headers for CORS
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Authorization")
	if req.Method == http.MethodOptions {
		return
	}

	faqsUseCase, err := handler.KnowledgeBaseInteractor.KnowledgeBase()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	var faqs []FaqPreview
	for _, faq := range faqsUseCase {
		faqs = append(faqs, FaqPreview{faq.ID, faq.MainExample, faq.IsTrained})
	}
	kb := KB{faqs}

	var body []byte
	if body, err = json.Marshal(kb); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Header().Add("Content-Type", "application/json")

	res.WriteHeader(200)
	res.Write(body)
	handler.Logger.Info("Returning response")
	return
}

// UpdateFaq is the handler function that updates a faq
func (handler WebserviceHandler) UpdateFaq(res http.ResponseWriter, req *http.Request) {
	handler.Logger.Info("Received " + req.Method + " request at path: " + req.URL.Path)

	// Setting headers for CORS
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Authorization")
	if req.Method == http.MethodOptions {
		return
	}

	var err error
	var updatedFaq Faq

	// Parsing the request body
	err = json.NewDecoder(req.Body).Decode(&updatedFaq)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Retrieving the ID from the url
	var id string
	ok := checkID(handler, res, req)
	if !ok {
		return
	}
	ids, ok := req.URL.Query()["id"]
	id = ids[0]

	// Data transformation
	usecasesFaq := webserviceFaqToUsecaseFaq(updatedFaq)

	// Adding the new Faq
	err = handler.KnowledgeBaseInteractor.Update(id, usecasesFaq)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Preparing the response
	res.WriteHeader(200)
	handler.Logger.Info("Returning response")
	return
}

// AddKeyword is the handler function that adds a keyword
func (handler WebserviceHandler) AddKeyword(res http.ResponseWriter, req *http.Request){
	handler.Logger.Info("Received " + req.Method + " request at path: " + req.URL.Path)

	// Setting headers for CORS
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Authorization")
	if req.Method == http.MethodOptions {
		return
	}

	var err error
	var newKeyword Keyword

	// Parsing the request body
	err = json.NewDecoder(req.Body).Decode(&newKeyword)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Data transformation
	usecaseKeyword := webserviceKeywordToUsecaseKeyword(newKeyword)

	// Adding the new Faq
	err = handler.KeywordsInteractor.Add(usecaseKeyword)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Preparing the response
	res.WriteHeader(200)
	handler.Logger.Info("Returning response")
	return
}

// DeleteKeyword is the handler function that deletes a keyword
func (handler WebserviceHandler) DeleteKeyword(res http.ResponseWriter, req *http.Request) {
	handler.Logger.Info("Received " + req.Method + " request at path: " + req.URL.Path)

	// Setting headers for CORS
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Authorization")
	if req.Method == http.MethodOptions {
		return
	}

	// Retrieving the ID from the url
	var id string
	var err error

	// Retrieving the ID
	id = req.URL.Query().Get("id")
	ok := checkID(handler, res, req)
	if !ok {
		return
	}
	handler.Logger.Info("ID: " + id)

	// Deleting the new Faq
	err = handler.KeywordsInteractor.Delete(id)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Preparing the response
	res.WriteHeader(200)
	handler.Logger.Info("Returning response")
	return
}

// UpdateKeyword updates a keyword
func (handler WebserviceHandler) UpdateKeyword(res http.ResponseWriter, req *http.Request) {
	handler.Logger.Info("Received " + req.Method + " request at path: " + req.URL.Path)

	// Setting headers for CORS
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Authorization")
	if req.Method == http.MethodOptions {
		return
	}

	var err error
	var updatedKeyword Keyword

	// Parsing the request body
	err = json.NewDecoder(req.Body).Decode(&updatedKeyword)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Retrieving the ID from the url
	var id string
	ok := checkID(handler, res, req)
	if !ok {
		return
	}
	ids, ok := req.URL.Query()["id"]
	id = ids[0]

	// Data transformation
	usecasesKeyword := webserviceKeywordToUsecaseKeyword(updatedKeyword)

	// Adding the new Faq
	err = handler.KeywordsInteractor.Update(id, usecasesKeyword)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Preparing the response
	res.WriteHeader(200)
	handler.Logger.Info("Returning response")
	return
}
