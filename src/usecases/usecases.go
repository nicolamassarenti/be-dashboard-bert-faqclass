package usecases

import (
	"fmt"

	"github.com/NicolaMassarenti/be-dashboard-bert-faqclass/src/domain"
)

// Faq contains the data that define a F.A.Q.
type Faq struct {
	ID               string
	MainExample      string
	Answers          []Answer
	IsTrained        bool
	TrainingExamples []TrainingExample
}

// TrainingExample contain the training examples of a specific language
type TrainingExample struct {
	Language string
	Examples []string
}

// Answer contains the answer in a language
type Answer struct {
	lang   string
	answer string
}

//Logger is the interface that manages the logs
type Logger interface {
	Log(message string)
	Error(message string)
}

//KnowledgeBaseInteractor is the object that manages the interactions
type KnowledgeBaseInteractor struct {
	FaqRepository domain.FaqRepository
	Logger        Logger
}

// KnowledgeBase returns all the knowledge base, all the faqs
func (kbInteractor *KnowledgeBaseInteractor) KnowledgeBase() (faqs []Faq, err error) {
	var message string

	kbInteractor.Logger.Log("Retrieving the KB")
	kb, domainErr := kbInteractor.FaqRepository.KnowledgeBase()

	if domainErr != nil {
		message = "Error retrieving the KB - %s"
		err = fmt.Errorf(message, domainErr.Error())
		kbInteractor.Logger.Log(err.Error())
		return nil, domainErr
	}

	faqs = make([]Faq, len(kb))

	// Transforming the KB from domain struct to usecase struct
	for i, faq := range kb {
		faqs[i] = faqFromDomainLayer(faq)
	}

	return faqs, nil
}

// Faq returns a faq with a given ID
func (kbInteractor *KnowledgeBaseInteractor) Faq(ID string) (Faq, error) {
	var message string

	message = fmt.Sprintf("Retrieving Faq with ID: %s", ID)
	kbInteractor.Logger.Log(message)

	faq, domainErr := kbInteractor.FaqRepository.Faq(ID)
	if domainErr != nil {
		message = "Error retrieving the Faq with ID: %s - %s"
		err := fmt.Errorf(message, ID, domainErr.Error())
		kbInteractor.Logger.Log(err.Error())
		return Faq{}, err
	}

	return faqFromDomainLayer(faq), nil
}

// ChangeTrainingStatus changes the training status of a FAQ
func (kbInteractor *KnowledgeBaseInteractor) ChangeTrainingStatus(ID string) error {
	kbInteractor.Logger.Log(fmt.Sprintf("Chaging training status of Faq with ID: %s", ID))

	if domainErr := kbInteractor.FaqRepository.ChangeTrainingStatus(ID); domainErr != nil {
		message := "Error changing the training status of Faq with ID: %s"
		err := fmt.Errorf(message, ID, domainErr.Error())
		kbInteractor.Logger.Log(err.Error())
		return err
	}
	return nil
}

// AddFaq adds a faq
func (kbInteractor *KnowledgeBaseInteractor) AddFaq(faq Faq) error {
	message := "Adding new faq with id %s"
	kbInteractor.Logger.Log(fmt.Sprintf(message, faq.ID))

	faqDomain := faqToDomainLayer(faq)

	domainErr := kbInteractor.FaqRepository.AddFaq(faqDomain)
	if domainErr != nil {
		message = "Error adding the new faq with id %s"
		err := fmt.Errorf(message, faq.ID, domainErr.Error())
		kbInteractor.Logger.Log(err.Error())
		return err
	}
	return nil
}

// DeleteFaq deletes a faq
func (kbInteractor *KnowledgeBaseInteractor) DeleteFaq(ID string) error {
	message := "Deleting faq with id: %s"
	kbInteractor.Logger.Log(fmt.Sprintf(message, ID))

	domainErr := kbInteractor.FaqRepository.DeleteFaq(ID)
	if domainErr != nil {
		message = "Error deleting faq with id %s"
		err := fmt.Errorf(message, ID, domainErr.Error())
		kbInteractor.Logger.Log(err.Error())
		return err
	}
	return nil
}
