package repositories

import (
	"github.com/nicolamassarenti/be-dashboard-bert-faqclass/src/domain"
)

// DBHandler defines the interface of the db handler
type DBHandler interface {
	Add(collection string, faq *map[string]interface{}) error
	ChangeBool(collection string, ID, path string, value bool) error
	Delete(collection string, ID string) error
	Get(collection string, ID string) (map[string]interface{}, error)
	GetAll(collection string) ([]map[string]interface{}, error)
	Update(collection string, ID string, data map[string]interface{}) error
}

// repositoryFaq defines the Faq managed by the repository
type repositoryFaq struct {
	ID   string     `json:"ID,omitempty"`
	Data domain.Faq `json:"data,omitempty"`
}

// repositoryKeyword defines the Keyword managed by the repository
type repositoryKeyword struct {
	ID   string         `json:"ID,omitempty"`
	Data domain.Keyword `json:"Data,omitempty"`
}

// LanguageDBRepo defines the DB handler and the collection for the languages
type LanguageDBRepo struct {
	Handler    DBHandler
	collection string
}

// FaqDBRepo defines the DB handler and the collection for the faqs
type FaqDBRepo struct {
	Handler    DBHandler
	collection string
}

// KeywordDBRepo defines the DB handler and the collection for the keywords
type KeywordDBRepo struct {
	Handler    DBHandler
	collection string
}

// KBHandler is the handler for the faqs
type KBHandler FaqDBRepo

// KeywordsHandler is the handler for the keywords
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