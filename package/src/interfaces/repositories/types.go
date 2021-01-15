package repositories

import (
	"github.com/nicolamassarenti/be-dashboard-bert-faqclass/src/domain"
)

// DBHandler is the interface for the handler of the DB for the FAQ
type DBHandler interface {
	Add(collection string, faq *map[string]interface{}) error
	ChangeBool(collection string, ID, path string, value bool) error
	Delete(collection string, ID string) error
	Get(collection string, ID string) (map[string]interface{}, error)
	GetAll(collection string) ([]map[string]interface{}, error)
	Update(collection string, ID string, data map[string]interface{}) error
}

// repositoryFaq is the Data retrieved by the repository
type repositoryFaq struct {
	ID   string     `json:"ID,omitempty"`
	Data domain.Faq `json:"data,omitempty"`
}

// repositoryKeywordWithID is the Data retrieved by the repository
type repositoryKeywordWithID struct {
	ID   string         `json:"ID,omitempty"`
	Data domain.Keyword `json:"Data,omitempty"`
}

// LanguageDBRepo is the object for the languages db handler
type LanguageDBRepo struct {
	Handler    DBHandler
	collection string
}

// FaqDBRepo is the object for the faq db handler
type FaqDBRepo struct {
	Handler    DBHandler
	collection string
}

// KeywordDBRepo is the object that manages the operations with the keywords
type KeywordDBRepo struct {
	Handler    DBHandler
	collection string
}

// KBHandler is the handler for the FAQ
type KBHandler FaqDBRepo

// KeywordHandler is the handler for the Keywords
type KeywordsHandler KeywordDBRepo

// LanguagesHandler is the handler for tha languages
type LanguagesHandler LanguageDBRepo

// NewLanguagesDBHandler creates a new handler for the languages
func NewLanguagesDBHandler(dbHandler DBHandler, collection string) *LanguagesHandler {

	languagesHandler := new(LanguagesHandler)
	languagesHandler.Handler = dbHandler
	languagesHandler.collection = collection
	return languagesHandler
}

// NewFaqDBHandler creates a new handler for the faq
func NewFaqDBHandler(dbHandler DBHandler, collection string) *KBHandler {

	kbHandler := new(KBHandler)
	kbHandler.Handler = dbHandler
	kbHandler.collection = collection
	return kbHandler
}

// NewKeywordsDBHandler creates a new handler for the keywords
func NewKeywordsDBHandler(dbHandler DBHandler, collection string) *KeywordsHandler {

	keywordsHandler := new(KeywordsHandler)
	keywordsHandler.Handler = dbHandler
	keywordsHandler.collection = collection
	return keywordsHandler
}