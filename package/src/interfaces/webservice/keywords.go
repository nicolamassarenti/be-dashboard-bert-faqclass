package webservice

import (
	"encoding/json"
	"net/http"
)

// AddKeyword is the webservice handler that adds a keyword
func (handler WebserviceHandler) AddKeyword(res http.ResponseWriter, req *http.Request) {
	handler.Logger.Info("Received " + req.Method + " request at path: " + req.URL.Path)

	// Setting headers for CORS
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Authorization")
	if req.Method == http.MethodOptions {
		return
	}

	var err error
	var newKeyword Keyword

	// Parsing request body
	handler.Logger.Debug("Starting to parse request body")
	err = json.NewDecoder(req.Body).Decode(&newKeyword)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	handler.Logger.Debug("Request body parsed")

	// Data transformation for presentation
	handler.Logger.Debug("Starting to transform data for presentation")
	usecaseKeyword := webserviceKeywordToUsecaseKeyword(newKeyword)
	handler.Logger.Debug("Data transformation completed")

	// Adding keyword
	handler.Logger.Debug("Starting to add keyword")
	err = handler.KeywordsInteractor.Add(usecaseKeyword)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	handler.Logger.Debug("Keyword added")

	// Preparing response
	res.WriteHeader(200)
	handler.Logger.Info("Returning response")
	return
}

// DeleteKeyword is the webservice handler that deletes a keyword
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

	// Checking request, verifying if ID is in query params
	handler.Logger.Debug("Starting to check the ID")
	id = req.URL.Query().Get("id")
	ok := checkID(handler, res, req)
	if !ok {
		return
	}
	handler.Logger.Debug("Request correct, ID inserted as query params")

	handler.Logger.Info("ID: " + id)

	// Deleting keyword
	handler.Logger.Debug("Starting to delete keyword")
	err = handler.KeywordsInteractor.Delete(id)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	handler.Logger.Debug("Keyword deleted")

	// Preparing response
	res.WriteHeader(200)
	handler.Logger.Info("Returning response")
	return
}

// UpdateKeyword is the webservice handler that deletes a keyword
func (handler WebserviceHandler) UpdateKeyword(res http.ResponseWriter, req *http.Request) {
	handler.Logger.Info("Received " + req.Method + " request at path: " + req.URL.Path)

	// Setting headers for CORS
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Authorization")
	if req.Method == http.MethodOptions {
		return
	}

	var err error

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

	handler.Logger.Debug("Starting to retrieve value from queryparams")
	var value string
	values, ok := req.URL.Query()["value"]
	value = values[0]
	handler.Logger.Debug("Param value retrieved")

	handler.Logger.Info("ID: " + id + " value: " + value)

	var updatedKeyword = Keyword{ID: id, DisplayText: value}

	// Data transformation for presentation
	handler.Logger.Debug("Starting to transform data for presentation")
	usecasesKeyword := webserviceKeywordToUsecaseKeyword(updatedKeyword)
	handler.Logger.Debug("Data transformed")

	// Updating keyword
	handler.Logger.Debug("Starting to update keyword")
	err = handler.KeywordsInteractor.Update(id, usecasesKeyword)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	handler.Logger.Debug("Keyword updated")

	// Preparing response
	res.WriteHeader(200)
	handler.Logger.Info("Returning response")
	return
}

// GetKeywords is the webservice handler that deletes a keyword
func (handler WebserviceHandler) GetKeywords(res http.ResponseWriter, req *http.Request) {
	handler.Logger.Info("Received " + req.Method + " request at path: " + req.URL.Path)

	// Setting headers for CORS
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Authorization")
	if req.Method == http.MethodOptions {
		return
	}

	handler.Logger.Debug("Starting to retrieve keywords")
	keywordsUseCase, err := handler.KeywordsInteractor.Keywords()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	handler.Logger.Debug("Keywords retrieved")

	handler.Logger.Debug("Transforming data for presentation")
	var keywords []Keyword
	for _, keyword := range keywordsUseCase {
		keywords = append(keywords, Keyword{keyword.ID, keyword.DisplayText})
	}
	kb := Keywords{keywords}
	handler.Logger.Debug("Data transformed")

	handler.Logger.Debug("Starting writing body")
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
