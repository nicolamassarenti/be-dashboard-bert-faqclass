package usecases

import (
	"fmt"

	"github.com/nicolamassarenti/be-dashboard-bert-faqclass/src/domain"
)

// Language defines the language
type Language struct {
	IsoName     string
	DisplayName string
}

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
	Lang   string
	Answer string
}

//Logger is the interface that manages the logs
type Logger interface {
	Info(message string)
	Debug(message string)
	Error(message string)
	Fatal(message string)
}

// LanguageRepository is the interface for the language repository
type LanguageRepository interface {
	GetAllLanguages() ([]Language, error)
}

//LanguageInteractor is the object that manages the interactions with the languages collection
type LanguageInteractor struct {
	LanguageRepository LanguageRepository
	Logger             Logger
}

//KnowledgeBaseInteractor is the object that manages the interactions with the KB collection
type KnowledgeBaseInteractor struct {
	FaqRepository domain.FaqRepository
	Logger        Logger
}

// AddFaq adds a faq
func (kbInteractor *KnowledgeBaseInteractor) AddFaq(faq Faq) error {

	faqDomain := faqToDomainLayer(faq)

	domainErr := kbInteractor.FaqRepository.AddFaq(faqDomain)
	if domainErr != nil {
		message := "Error adding a new faq. "
		err := fmt.Errorf(message, domainErr.Error())
		kbInteractor.Logger.Error(err.Error())
		return err
	}
	return nil
}

// ChangeTrainingStatus changes the training status of a FAQ
func (kbInteractor *KnowledgeBaseInteractor) ChangeTrainingStatus(ID string, newStatus bool) error {

	if ID == "" {
		message := "ID not valid - received %s"
		err := fmt.Errorf(message, ID)
		kbInteractor.Logger.Error(err.Error())
		return err
	}

	kbInteractor.Logger.Info(fmt.Sprintf("Chaging training status of Faq with ID: %s", ID))

	if domainErr := kbInteractor.FaqRepository.ChangeTrainingStatus(ID, newStatus); domainErr != nil {
		message := "Error changing the training status of Faq with ID: %s"
		err := fmt.Errorf(message, ID, domainErr.Error())
		kbInteractor.Logger.Error(err.Error())
		return err
	}
	return nil
}

// DeleteFaq deletes a faq
func (kbInteractor *KnowledgeBaseInteractor) DeleteFaq(ID string) error {
	message := "Deleting faq with id: %s"
	kbInteractor.Logger.Info(fmt.Sprintf(message, ID))

	domainErr := kbInteractor.FaqRepository.DeleteFaq(ID)
	if domainErr != nil {
		message = "Error deleting faq with id %s"
		err := fmt.Errorf(message, ID, domainErr.Error())
		kbInteractor.Logger.Error(err.Error())
		return err
	}
	return nil
}

// Faq returns a faq with a given ID
func (kbInteractor *KnowledgeBaseInteractor) Faq(ID string) (Faq, error) {
	var message string

	if ID == "" {
		message = "ID not valid - received %s"
		err := fmt.Errorf(message, ID)
		kbInteractor.Logger.Error(err.Error())
		return Faq{}, err
	}

	message = fmt.Sprintf("Retrieving Faq with ID: %s", ID)
	kbInteractor.Logger.Info(message)

	faq, domainErr := kbInteractor.FaqRepository.Faq(ID)
	if domainErr != nil {
		message = "Error retrieving the Faq with ID: %s - %s"
		err := fmt.Errorf(message, ID, domainErr.Error())
		kbInteractor.Logger.Error(err.Error())
		return Faq{}, err
	}

	return faqFromDomainLayer(faq), nil
}

// GetAllLanguages returns all the languages
func (langInteractor *LanguageInteractor) GetAllLanguages() (langs []Language, err error) {
	langInteractor.Logger.Info("Retrieving the languages")

	langs, err = langInteractor.LanguageRepository.GetAllLanguages()
	if err != nil {
		message := "Error retrieving the languages - %s"
		err = fmt.Errorf(message, err.Error())
		langInteractor.Logger.Error(err.Error())
	}
	return
}

// KnowledgeBase returns all the knowledge base, all the faqs
func (kbInteractor *KnowledgeBaseInteractor) KnowledgeBase() (faqs []Faq, err error) {
	var message string

	kbInteractor.Logger.Info("Retrieving the KB")
	kb, domainErr := kbInteractor.FaqRepository.KnowledgeBase()

	if domainErr != nil {
		message = "Error retrieving the KB - %s"
		err = fmt.Errorf(message, domainErr.Error())
		kbInteractor.Logger.Error(err.Error())
		return nil, domainErr
	}

	faqs = make([]Faq, len(kb))

	// Transforming the KB from domain struct to usecase struct
	for i, faq := range kb {
		faqs[i] = faqFromDomainLayer(faq)
	}

	return faqs, nil
}

// UpdateFaq updates an existing faq
func (kbInteractor *KnowledgeBaseInteractor) Update(ID string, faq Faq) error {
	message := "Updating faq with id: %s"
	kbInteractor.Logger.Info(fmt.Sprintf(message, ID))

	faqDomain := faqToDomainLayer(faq)

	domainErr := kbInteractor.FaqRepository.Update(ID, faqDomain)
	if domainErr != nil {
		message = "Error deleting faq with id %s"
		err := fmt.Errorf(message, ID, domainErr.Error())
		kbInteractor.Logger.Error(err.Error())
		return err
	}
	return nil
}
