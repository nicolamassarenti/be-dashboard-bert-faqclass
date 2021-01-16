package usecases

import "fmt"

// Languages returns all the languages
func (langInteractor *LanguageInteractor) Languages() (langs []Language, err error) {
	langInteractor.Logger.Info("Retrieving the languages")

	langs, err = langInteractor.Repository.Languages()
	if err != nil {
		message := "error retrieving the languages - %s"
		err = fmt.Errorf(message, err.Error())
		langInteractor.Logger.Error(err.Error())
	}
	return
}
