package usecases

import "fmt"

// GetAllLanguages returns all the languages
func (langInteractor *LanguageInteractor) GetAllLanguages() (langs []Language, err error) {
	langInteractor.Logger.Info("Retrieving the languages")

	langs, err = langInteractor.Repository.GetAllLanguages()
	if err != nil {
		message := "error retrieving the languages - %s"
		err = fmt.Errorf(message, err.Error())
		langInteractor.Logger.Error(err.Error())
	}
	return
}
