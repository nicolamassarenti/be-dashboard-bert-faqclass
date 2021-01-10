package webservice

import (
	"encoding/json"
	"net/http"
)

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
