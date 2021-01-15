package usecases

import (
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
	ID          string
	DisplayText string
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
	Languages() ([]Language, error)
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
