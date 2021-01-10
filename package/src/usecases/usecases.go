package usecases

import (
	"fmt"
	"strconv"

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

// Keyword is a keyword
type Keyword struct {
	ID string
	Name string
}

// TrainingExample contain the training examples of a specific language
type TrainingExample struct {
	Language string
	Examples []string
}

// Answer contains the answer in a language
type Answer struct {
	Language   string
	Answers []string
}

// Logger is the interface that manages the logs
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

// LanguageInteractor is the object that manages the interactions with the languages collection
type LanguageInteractor struct {
	Repository LanguageRepository
	Logger     Logger
}

// KnowledgeBaseInteractor is the object that manages the interactions with the KB collection
type KnowledgeBaseInteractor struct {
	Repository domain.FaqRepository
	Logger     Logger
}

// KeywordsInteractor is the object that manages the interactions with the KB collection
type KeywordsInteractor struct {
	Repository domain.KeywordRepository
	Logger     Logger
}

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