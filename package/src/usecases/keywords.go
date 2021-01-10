package usecases

import (
	"fmt"
	"github.com/nicolamassarenti/be-dashboard-bert-faqclass/src/domain"
	"strconv"
)

// Add adds a keyword
func (interactor *KeywordsInteractor) Add(keyword Keyword) error {

	var keywordDomain domain.Keyword
	keywordDomain.Name = keyword.Name

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
	keywordDomain.Name = keyword.Name
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

// Faq returns a faq with a given ID
func (interactor *KeywordsInteractor) Keyword(ID string) (Keyword, error) {
	var message string

	message = fmt.Sprintf("Retrieving Faq with ID: %s", ID)
	interactor.Logger.Info(message)

	keyword, domainErr := interactor.Repository.Keyword(ID)
	if domainErr != nil {
		message = "Error retrieving the Faq with ID: %s - %s"
		err := fmt.Errorf(message, ID, domainErr.Error())
		interactor.Logger.Error(err.Error())
		return Keyword{}, err
	}

	return Keyword{Name: keyword.Name, ID: keyword.ID}, nil
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
		keywords[i] = Keyword{ID: item.ID, Name: item.Name}
	}

	interactor.Logger.Info("Retrieved " + strconv.Itoa(len(keywords)) + " faqs")
	return keywords, nil
}