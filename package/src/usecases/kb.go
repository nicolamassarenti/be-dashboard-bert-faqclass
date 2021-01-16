package usecases

import (
	"fmt"
	"strconv"
)

// AddFaq adds a faq
func (interactor *KnowledgeBaseInteractor) AddFaq(faq Faq) error {
	interactor.Logger.Debug("Starting to transform data")
	faqDomain := usecaseFaqToDomainFaq(faq)
	interactor.Logger.Debug("Data transformed")

	interactor.Logger.Debug("Starting to add faq")
	domainErr := interactor.Repository.AddFaq(faqDomain)
	interactor.Logger.Debug("Faq added")
	if domainErr != nil {
		message := "error adding a new faq"
		err := fmt.Errorf(message, domainErr.Error())
		interactor.Logger.Error(err.Error())
		return err
	}
	return nil
}

// ChangeTrainingStatus changes the training status of a FAQ
func (interactor *KnowledgeBaseInteractor) ChangeTrainingStatus(ID string, newStatus bool) error {
	interactor.Logger.Debug("Starting to change training status")
	if domainErr := interactor.Repository.ChangeTrainingStatus(ID, newStatus); domainErr != nil {
		message := "error changing the training status of Data with ID: %s"
		err := fmt.Errorf(message, ID, domainErr.Error())
		interactor.Logger.Error(err.Error())
		return err
	}
	interactor.Logger.Debug("Training status changed")
	return nil
}

// DeleteFaq deletes a faq
func (interactor *KnowledgeBaseInteractor) DeleteFaq(ID string) error {
	var message string
	interactor.Logger.Debug("Starting to delete faq")
	domainErr := interactor.Repository.DeleteFaq(ID)
	if domainErr != nil {
		message = "Error deleting faq with id %s"
		err := fmt.Errorf(message, ID, domainErr.Error())
		interactor.Logger.Error(err.Error())
		return err
	}
	interactor.Logger.Debug("Faq deleted")
	return nil
}

// Faq returns a faq
func (interactor *KnowledgeBaseInteractor) Faq(ID string) (Faq, error) {
	var message string
	interactor.Logger.Debug("Starting to retrieve faq")
	faq, domainErr := interactor.Repository.Faq(ID)
	if domainErr != nil {
		message = "Error retrieving the Data with ID: %s - %s"
		err := fmt.Errorf(message, ID, domainErr.Error())
		interactor.Logger.Error(err.Error())
		return Faq{}, err
	}
	interactor.Logger.Debug("Faq retrieved")
	return domainFaqToUsecaseFaq(faq), nil
}

// KnowledgeBase returns all the knowledge base, all the faqs
func (interactor *KnowledgeBaseInteractor) KnowledgeBase() (faqs []Faq, err error) {
	var message string

	interactor.Logger.Debug("Starting to retrieve KB")
	kb, domainErr := interactor.Repository.KnowledgeBase()

	if domainErr != nil {
		message = "Error retrieving KB - %s"
		err = fmt.Errorf(message, domainErr.Error())
		interactor.Logger.Error(err.Error())
		return nil, domainErr
	}

	faqs = make([]Faq, len(kb))

	// Transforming the KB from domain struct to usecase struct
	for i, faq := range kb {
		faqs[i] = domainFaqToUsecaseFaq(faq)
	}
	interactor.Logger.Debug("Knowledge base retrieved")

	interactor.Logger.Info("Retrieved " + strconv.Itoa(len(faqs)) + " faqs")
	return faqs, nil
}

// Update updates an existing faq
func (interactor *KnowledgeBaseInteractor) Update(ID string, faq Faq) error {
	interactor.Logger.Debug("Starting to transform data")
	faqDomain := usecaseFaqToDomainFaq(faq)
	interactor.Logger.Debug("Data transformed")

	interactor.Logger.Debug("Starting to update faq")
	domainErr := interactor.Repository.Update(ID, faqDomain)
	if domainErr != nil {
		var message string
		message = "Error deleting faq with id %s"
		err := fmt.Errorf(message, ID, domainErr.Error())
		interactor.Logger.Error(err.Error())
		return err
	}
	interactor.Logger.Debug("Faq updated")
	return nil
}

