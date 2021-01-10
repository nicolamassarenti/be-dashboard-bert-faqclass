package usecases

import (
	"fmt"
	"github.com/nicolamassarenti/be-dashboard-bert-faqclass/src/domain"
	"strconv"
)

// Add adds a keyword
func (interactor *KeywordsInteractor) Add(keyword Keyword) error {

	var keywordDomain domain.Keyword
	keywordDomain.DisplayText = keyword.DisplayText

	domainErr := interactor.Repository.Add(keywordDomain)
	if domainErr != nil {
		message := "error adding a new faq"
		err := fmt.Errorf(message, domainErr.Error())
		interactor.Logger.Error(err.Error())
		return err
	}
	return nil
}

// Update updates an existing keyword
func (interactor *KeywordsInteractor) Update(ID string, keyword Keyword) error {
	message := "Updating faq with id: %s"
	interactor.Logger.Info(fmt.Sprintf(message, ID))

	var keywordDomain domain.Keyword
	keywordDomain.DisplayText = keyword.DisplayText
	keywordDomain.ID = keyword.ID

	domainErr := interactor.Repository.Update(ID, keywordDomain)
	if domainErr != nil {
		message = "Error deleting keyword with id %s"
		err := fmt.Errorf(message, ID, domainErr.Error())
		interactor.Logger.Error(err.Error())
		return err
	}
	return nil
}

// Delete deletes an existing keyword
func (interactor *KeywordsInteractor) Delete(ID string) error {
	message := "Deleting faq with id: %s"
	interactor.Logger.Info(fmt.Sprintf(message, ID))

	domainErr := interactor.Repository.Delete(ID)
	if domainErr != nil {
		message = "Error deleting faq with id %s"
		err := fmt.Errorf(message, ID, domainErr.Error())
		interactor.Logger.Error(err.Error())
		return err
	}
	return nil
}

// KnowledgeBase returns all the knowledge base, all the faqs
func (interactor *KeywordsInteractor) Keywords() (keywords []Keyword, err error) {
	var message string

	interactor.Logger.Info("Retrieving the KB")
	allKeywords, domainErr := interactor.Repository.Keywords()

	if domainErr != nil {
		message = "Error retrieving the KB - %s"
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