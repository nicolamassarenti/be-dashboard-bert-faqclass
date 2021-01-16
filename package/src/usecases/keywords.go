package usecases

import (
	"fmt"
	"github.com/nicolamassarenti/be-dashboard-bert-faqclass/src/domain"
	"strconv"
)

// Add adds a keyword
func (interactor *KeywordsInteractor) Add(keyword Keyword) error {
	interactor.Logger.Debug("Starting to add keyword")
	var keywordDomain domain.Keyword
	keywordDomain.DisplayText = keyword.DisplayText

	domainErr := interactor.Repository.Add(keywordDomain)
	if domainErr != nil {
		message := "error adding a new faq"
		err := fmt.Errorf(message, domainErr.Error())
		interactor.Logger.Error(err.Error())
		return err
	}
	interactor.Logger.Debug("Keyword added")
	return nil
}

// Update updates a keyword
func (interactor *KeywordsInteractor) Update(ID string, keyword Keyword) error {
	interactor.Logger.Debug("Starting to update keyword")

	var keywordDomain domain.Keyword
	keywordDomain.DisplayText = keyword.DisplayText
	keywordDomain.ID = keyword.ID

	domainErr := interactor.Repository.Update(ID, keywordDomain)
	if domainErr != nil {
		message := "Error deleting keyword with id %s"
		err := fmt.Errorf(message, ID, domainErr.Error())
		interactor.Logger.Error(err.Error())
		return err
	}
	interactor.Logger.Debug("Keyword updated")
	return nil
}

// Delete deletes a keyword
func (interactor *KeywordsInteractor) Delete(ID string) error {
	interactor.Logger.Debug("Starting to delete keyword")

	domainErr := interactor.Repository.Delete(ID)
	if domainErr != nil {
		message := "Error deleting faq with id %s"
		err := fmt.Errorf(message, ID, domainErr.Error())
		interactor.Logger.Error(err.Error())
		return err
	}
	interactor.Logger.Debug("Keyword deleted")
	return nil
}

// Keywords returns all the keywords
func (interactor *KeywordsInteractor) Keywords() (keywords []Keyword, err error) {
	interactor.Logger.Debug("Starting to retrieve KB")
	allKeywords, domainErr := interactor.Repository.Keywords()

	if domainErr != nil {
		message := "Error retrieving the KB - %s"
		err = fmt.Errorf(message, domainErr.Error())
		interactor.Logger.Error(err.Error())
		return nil, domainErr
	}

	keywords = make([]Keyword, len(allKeywords))

	// Transforming the KB from domain struct to usecase struct
	for i, item := range allKeywords {
		keywords[i] = Keyword{ID: item.ID, DisplayText: item.DisplayText}
	}

	interactor.Logger.Info("Retrieved " + strconv.Itoa(len(keywords)) + " keywords")
	return keywords, nil
}