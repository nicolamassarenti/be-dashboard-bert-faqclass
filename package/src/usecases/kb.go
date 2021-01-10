package usecases

import (
	"fmt"
	"strconv"
)

// AddFaq adds a faq
func (interactor *KnowledgeBaseInteractor) AddFaq(faq Faq) error {

	faqDomain := faqToDomainLayer(faq)

	domainErr := interactor.Repository.AddFaq(faqDomain)
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

	if ID == "" {
		message := "ID not valid - received %s"
		err := fmt.Errorf(message, ID)
		interactor.Logger.Error(err.Error())
		return err
	}

	interactor.Logger.Info(fmt.Sprintf("Chaging training status of Faq with ID: %s", ID))

	if domainErr := interactor.Repository.ChangeTrainingStatus(ID, newStatus); domainErr != nil {
		message := "error changing the training status of Faq with ID: %s"
		err := fmt.Errorf(message, ID, domainErr.Error())
		interactor.Logger.Error(err.Error())
		return err
	}
	return nil
}

// DeleteFaq deletes a faq
func (interactor *KnowledgeBaseInteractor) DeleteFaq(ID string) error {
	message := "Deleting faq with id: %s"
	interactor.Logger.Info(fmt.Sprintf(message, ID))

	domainErr := interactor.Repository.DeleteFaq(ID)
	if domainErr != nil {
		message = "Error deleting faq with id %s"
		err := fmt.Errorf(message, ID, domainErr.Error())
		interactor.Logger.Error(err.Error())
		return err
	}
	return nil
}

// Faq returns a faq with a given ID
func (interactor *KnowledgeBaseInteractor) Faq(ID string) (Faq, error) {
	var message string

	message = fmt.Sprintf("Retrieving Faq with ID: %s", ID)
	interactor.Logger.Info(message)

	faq, domainErr := interactor.Repository.Faq(ID)
	if domainErr != nil {
		message = "Error retrieving the Faq with ID: %s - %s"
		err := fmt.Errorf(message, ID, domainErr.Error())
		interactor.Logger.Error(err.Error())
		return Faq{}, err
	}

	return faqFromDomainLayer(faq), nil
}

// KnowledgeBase returns all the knowledge base, all the faqs
func (interactor *KnowledgeBaseInteractor) KnowledgeBase() (faqs []Faq, err error) {
	var message string

	interactor.Logger.Info("Retrieving the KB")
	kb, domainErr := interactor.Repository.KnowledgeBase()

	if domainErr != nil {
		message = "Error retrieving the KB - %s"
		err = fmt.Errorf(message, domainErr.Error())
		interactor.Logger.Error(err.Error())
		return nil, domainErr
	}

	faqs = make([]Faq, len(kb))

	// Transforming the KB from domain struct to usecase struct
	for i, faq := range kb {
		faqs[i] = faqFromDomainLayer(faq)
	}

	interactor.Logger.Info("Retrieved " + strconv.Itoa(len(faqs)) + " faqs")
	return faqs, nil
}

// Update updates an existing faq
func (interactor *KnowledgeBaseInteractor) Update(ID string, faq Faq) error {
	message := "Updating faq with id: %s"
	interactor.Logger.Info(fmt.Sprintf(message, ID))

	faqDomain := faqToDomainLayer(faq)

	domainErr := interactor.Repository.Update(ID, faqDomain)
	if domainErr != nil {
		message = "Error deleting faq with id %s"
		err := fmt.Errorf(message, ID, domainErr.Error())
		interactor.Logger.Error(err.Error())
		return err
	}
	return nil
}

