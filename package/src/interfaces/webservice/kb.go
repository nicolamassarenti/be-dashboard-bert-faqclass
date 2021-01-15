package webservice

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// AddFaq is the webservice handler function that adds a new Faq
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

	// Transforming data from webservice definition to usecase definition
	handler.Logger.Debug("Starting to transform faq to usecase faq")
	usecasesFaq := webserviceFaqToUsecaseFaq(newFaq)
	handler.Logger.Debug("Transformation to usecase faq completed")

	// Adding new faq
	handler.Logger.Debug("Starting to add faq")
	err = handler.KnowledgeBaseInteractor.AddFaq(usecasesFaq)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	handler.Logger.Debug("Faq added")

	// Preparing response
	res.WriteHeader(200)
	handler.Logger.Info("Returning response")
	return
}

// ChangeTrainingStatus is the webservice handler that changes the training status of a faq
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

	// Checking request, verifying if ID is in query params
	handler.Logger.Debug("Starting to check the ID")
	ok := checkID(handler, res, req)
	if !ok {
		return
	}
	id = req.URL.Query().Get("id")
	handler.Logger.Debug("Request correct, ID inserted as query params")

	// Checking request, verifying if toTrain is in query params
	handler.Logger.Debug("Starting to check the toTrain")
	toTrainString, ok := req.URL.Query()["toTrain"]
	if !ok || len(toTrainString) != 1 {
		if !ok {
			handler.Logger.Info("Error retrieving param toTrain. Returning BadRequest")
		} else if len(toTrainString) == 0 {
			handler.Logger.Info("No toTrain param as query params. Returning BadRequest")
		} else {
			handler.Logger.Info("More than one toTrain param in query params. Returning BadRequest")
		}
		res.WriteHeader(http.StatusBadRequest)
	}
	toTrain, err = strconv.ParseBool(req.URL.Query().Get("toTrain"))
	if err != nil {
		handler.Logger.Info("Error transforming toTrain from string to bool. Returning Internal Server Error")
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	handler.Logger.Debug("Starting to check the toTrain")

	handler.Logger.Info("ID: " + id + "\ttoTrain: " + strconv.FormatBool(toTrain))

	// Executing change of training status
	handler.Logger.Debug("Starting to execute the change of training status")
	err = handler.KnowledgeBaseInteractor.ChangeTrainingStatus(id, toTrain)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	handler.Logger.Debug("Change of training status executed")

	// Preparing response
	res.WriteHeader(200)
	handler.Logger.Info("Returning response")
	return
}

// DeleteFaq is the webservice handler that deletes a faq
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

	// Checking request, verifying if ID is in query params
	handler.Logger.Debug("Starting to check the ID")
	id = req.URL.Query().Get("id")
	ok := checkID(handler, res, req)
	if !ok {
		return
	}
	handler.Logger.Debug("Request correct, ID inserted as query params")

	handler.Logger.Info("ID: " + id)

	// Deleting the new Data
	handler.Logger.Debug("Starting to delete faq")
	err = handler.KnowledgeBaseInteractor.DeleteFaq(id)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	handler.Logger.Debug("Faq deleted")

	// Preparing response
	res.WriteHeader(200)
	handler.Logger.Info("Returning response")
	return
}

// Faq is the webservice handler that returns a faq
func (handler WebserviceHandler) Faq(res http.ResponseWriter, req *http.Request) {
	handler.Logger.Info("Received " + req.Method + " request at path: " + req.URL.Path)

	// Setting headers for CORS
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Authorization")
	if req.Method == http.MethodOptions {
		return
	}

	// Checking request, verifying if ID is in query params
	handler.Logger.Debug("Starting to check the ID")
	var id string
	ok := checkID(handler, res, req)
	if !ok {
		return
	}
	ids, ok := req.URL.Query()["id"]
	id = ids[0]
	handler.Logger.Debug("Request correct, ID inserted as query params")

	handler.Logger.Info("ID: " + id)

	// Retrieving faq
	handler.Logger.Debug("Starting to retrieve faq")
	usecaseFaq, err := handler.KnowledgeBaseInteractor.Faq(id)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	handler.Logger.Debug("Faq retrieved")

	// Transforming data for presentation
	handler.Logger.Debug("Starting to transform data for presentation")
	faq := usecaseFaqToWebserviceFaq(usecaseFaq)
	handler.Logger.Debug("Data transformed")

	// Preparing response
	var body []byte
	if body, err = json.Marshal(faq); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.Header().Add("Content-Type", "application/json")

	res.WriteHeader(200)
	_, err = res.Write(body)
	if err != nil {
		return
	}
	handler.Logger.Info("Returning response")
	return
}

// KnowledgeBase is the webservice handler that returns the kb
func (handler WebserviceHandler) KnowledgeBase(res http.ResponseWriter, req *http.Request) {
	handler.Logger.Info("Received " + req.Method + " request at path: " + req.URL.Path)

	// Setting headers for CORS
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Authorization")
	if req.Method == http.MethodOptions {
		return
	}

	handler.Logger.Debug("Starting to retrieve knowledge base")
	faqsUseCase, err := handler.KnowledgeBaseInteractor.KnowledgeBase()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	handler.Logger.Debug("Knowledge base retrieved")

	handler.Logger.Debug("Starting to transform data for presentation")
	var faqs []FaqPresentation
	for _, faq := range faqsUseCase {
		faqs = append(faqs, FaqPresentation{faq.ID, faq.MainExample, faq.IsTrained})
	}
	kb := KB{faqs}
	handler.Logger.Debug("Data transformed")

	handler.Logger.Debug("Starting to write body")
	var body []byte
	if body, err = json.Marshal(kb); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	handler.Logger.Debug("Body written")

	res.Header().Add("Content-Type", "application/json")

	res.WriteHeader(200)
	_, err = res.Write(body)
	if err != nil {
		return
	}
	handler.Logger.Info("Returning response")
	return
}

// UpdateFaq is the webservice handler that updates a faq
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

	// Checking request, verifying if ID is in query params
	handler.Logger.Debug("Starting to check the ID")
	var id string
	ok := checkID(handler, res, req)
	if !ok {
		return
	}
	ids, ok := req.URL.Query()["id"]
	id = ids[0]
	handler.Logger.Debug("Request correct, ID inserted as query params")

	handler.Logger.Info("ID: " + id)

	// Transforming data for presentation
	handler.Logger.Debug("Starting to transform data for presentation")
	usecasesFaq := webserviceFaqToUsecaseFaq(updatedFaq)
	handler.Logger.Debug("Data transformed for presentation")

	// Updating faq
	handler.Logger.Debug("Starting to update faq")
	err = handler.KnowledgeBaseInteractor.Update(id, usecasesFaq)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	handler.Logger.Debug("Faq updated")

	// Preparing response
	res.WriteHeader(200)
	handler.Logger.Info("Returning response")
	return
}
