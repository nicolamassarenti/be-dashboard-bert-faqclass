package webservice

import (
	"encoding/json"
	"net/http"
)

// AddKeyword is the handler function that adds a keyword
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

// GetKeywords gets all keywords
func (handler WebserviceHandler) GetKeywords(res http.ResponseWriter, req *http.Request) {
	handler.Logger.Info("Received " + req.Method + " request at path: " + req.URL.Path)

	// Setting headers for CORS
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Authorization")
	if req.Method == http.MethodOptions {
		return
	}

	keywordsUseCase, err := handler.KeywordsInteractor.Keywords()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	var keywords []Keyword
	for _, keyword := range keywordsUseCase {
		keywords = append(keywords, Keyword{keyword.ID, keyword.Name})
	}
	kb := Keywords{keywords}

	var body []byte
	if body, err = json.Marshal(kb); err != nil {
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

// GetKeyword gets a specific keyword
func (handler WebserviceHandler) GetKeyword(res http.ResponseWriter, req *http.Request) {
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
	keyword, err := handler.KeywordsInteractor.Keyword(id)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Preparing the response
	var body []byte
	if body, err = json.Marshal(keyword); err != nil {
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
