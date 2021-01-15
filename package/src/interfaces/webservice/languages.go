package webservice

import (
	"encoding/json"
	"net/http"
)

// GetAllLanguages is the webservice handler returns the languages
func (handler WebserviceHandler) GetAllLanguages(res http.ResponseWriter, req *http.Request) {
	handler.Logger.Info("Received " + req.Method + " request at path: " + req.URL.Path)

	// Setting headers for CORS
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Authorization")
	if req.Method == http.MethodOptions {
		return
	}

	handler.Logger.Debug("Starting to retrieve languages")
	languages, err := handler.LanguagesInteractor.Languages()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	handler.Logger.Debug("Languages retrieved")

	handler.Logger.Debug("Starting to transform data for presentation")
	languagesMap := usecaseLanguageToMapStringString(languages)
	handler.Logger.Debug("Data transformed")

	handler.Logger.Debug("Starting to write body")
	var body []byte
	if body, err = json.Marshal(languagesMap); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	handler.Logger.Debug("Body written")

	res.Header().Add("Content-Type", "application/json")

	res.WriteHeader(200)
	res.Write(body)
	handler.Logger.Info("Returning response")

	return
}
