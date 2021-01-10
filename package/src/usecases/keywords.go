package usecases

import (
	"fmt"
	"github.com/nicolamassarenti/be-dashboard-bert-faqclass/src/domain"
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
