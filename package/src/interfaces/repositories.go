package interfaces

import (
	"github.com/mitchellh/mapstructure"
	"time"

	"github.com/nicolamassarenti/be-dashboard-bert-faqclass/src/domain"
	"github.com/nicolamassarenti/be-dashboard-bert-faqclass/src/usecases"
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

// repositoryFaqWithID is the Faq retrieved by the repository
type repositoryFaqWithID struct {
	ID  string     `json:"ID,omitempty"`
	Faq domain.Faq `json:"Faq,omitempty"`
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
	Handler DBHandler
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

// GetAllLanguages returns all the languages
func (repo *LanguagesHandler) GetAllLanguages() ([]usecases.Language, error) {
	langsMap, err := repo.Handler.GetAll(repo.collection)

	langs := make([]usecases.Language, len(langsMap))
	for idx, lang := range langsMap{
		data, _ := lang["faq"].(map[string]interface{})
		langs[idx] = mapStringInterfaceToUsecasesLang(data)
	}

	return langs, err
}

// Returns the map[string]interface in the format of the database
func getFaqMapToAdd(faq domain.Faq) map[string]interface{}{
	return map[string]interface{}{
		"MainExample":      faq.MainExample,
		"Answers":          faq.Answers,
		"IsTrained":        faq.IsTrained,
		"TrainingExamples": faq.TrainingExamples,
		"UpdateDate": 		time.Now().Format(time.RFC3339),
	}
}

// Retruns the map[string]interface formatted as requested by the database
func getKeywordMapToAdd(keyword domain.Keyword) map[string]interface{}{
	return map[string]interface{}{
		"Name": keyword.Name,
		"UpdateDate": 		time.Now().Format(time.RFC3339),
	}
}

// KnowledgeBase is the implementation that returns all the faq of the knowledge base
func (repo *KBHandler) KnowledgeBase() ([]domain.Faq, error) {
	var repFaqArray []repositoryFaqWithID
	faqs, err := repo.Handler.GetAll(repo.collection)
	if err != nil {
		return nil, err
	}

	// decoding the map to my type `repositoryFaqWithID`
	mapstructure.Decode(faqs, &repFaqArray)

	var kb []domain.Faq
	for _, repFaq := range repFaqArray {
		kb = append(
			kb,
			domain.Faq{
				ID:               repFaq.ID,
				MainExample:      repFaq.Faq.MainExample,
				Answers:          repFaq.Faq.Answers,
				IsTrained:        repFaq.Faq.IsTrained,
				TrainingExamples: repFaq.Faq.TrainingExamples,
			},
		)
	}

	return kb, nil
}

// Faq is the implementation that returns a specific ID
func (repo *KBHandler) Faq(ID string) (faq domain.Faq, err error) {

	faqMap, err := repo.Handler.Get(repo.collection, ID)
	if err != nil {
		return
	}

	mapstructure.Decode(faqMap, &faq)
	return
}

// ChangeTrainingStatus changes the "isTrained" bool of a Faq
func (repo *KBHandler) ChangeTrainingStatus(ID string, newStatus bool) error {
	path := "IsTrained"
	return repo.Handler.ChangeBool(repo.collection, ID, path, newStatus)
}

// AddFaq adds a new faq
func (repo *KBHandler) AddFaq(faq domain.Faq) error {
	faqMap := getFaqMapToAdd(faq)
	return repo.Handler.Add(repo.collection, &faqMap)

}

// DeleteFaq deletes a faq
func (repo *KBHandler) DeleteFaq(ID string) error {
	return repo.Handler.Delete(repo.collection, ID)
}

// AddFaq adds a new faq
func (repo *KBHandler) Update(ID string, faq domain.Faq) error {
	faqMap := getFaqMapToAdd(faq)
	return repo.Handler.Update(repo.collection, ID, faqMap)

}

// AddKeyword adds a new keyword
func (repo *KeywordsHandler) AddKeyword(keyword domain.Keyword) error {
	keywordMap := getKeywordMapToAdd(keyword)
	return repo.Handler.Add(repo.collection, &keywordMap)
}