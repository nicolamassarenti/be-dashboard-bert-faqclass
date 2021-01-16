package usecases

import (
	"github.com/nicolamassarenti/be-dashboard-bert-faqclass/src/domain"
)

// Language defines a language
type Language struct {
	IsoName     string
	DisplayName string
}

// Faq defines a faq
type Faq struct {
	ID               string
	MainExample      string
	Answers          []Answer
	IsTrained        bool
	TrainingExamples []TrainingExamples
}

// Keyword defines a keyword
type Keyword struct {
	ID          string
	DisplayText string
}

// TrainingExamples defines a set of training example
type TrainingExamples struct {
	Language string
	Examples []string
}

// Answer contains the answer in a language
type Answer struct {
	Language   string
	Answers []string
}

// Logger defines the interface for the logger
type Logger interface {
	Info(message string)
	Debug(message string)
	Error(message string)
	Fatal(message string)
}

// LanguageRepository defines interface for the language repository
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
