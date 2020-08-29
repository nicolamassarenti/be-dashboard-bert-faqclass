package interfaces

import (
	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"

	"github.com/nicolamassarenti/be-dashboard-bert-faqclass/src/domain"
	"github.com/nicolamassarenti/be-dashboard-bert-faqclass/src/usecases"
)

// DBHandler is the interface for the handler of the DB for the FAQ
type DBHandler interface {
	Store(collection string, faq *map[string]interface{}) error
	Get(collection string, ID string) (map[string]interface{}, error)
	GetAll(collection string) ([]map[string]interface{}, error)
	ChangeBool(collection string, ID, path string, value bool) error
	Delete(collection string, ID string) error
}

type repositoryFaq struct {
	MainExample      string
	Answers          []answer
	IsTrained        bool
	TrainingExamples []trainingExample
}

type trainingExample struct {
	Language string
	Examples []string
}

type answer struct {
	Lang   string
	Answer string
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

// KBHandler is the handler for the FAQ
type KBHandler FaqDBRepo

// LanguagesHandler is the handler for tha languages
type LanguagesHandler LanguageDBRepo

// NewLanguagesDBHandler creates a new handler for the languages
func NewLanguagesDBHandler(dbHandler DBHandler, collection string) *LanguagesHandler {

	languagesHandler := new(LanguagesHandler)
	languagesHandler.Handler = dbHandler
	languagesHandler.collection = collection
	return languagesHandler
}

// GetAllLanguages returns all the languages
func (repo *LanguagesHandler) GetAllLanguages() ([]usecases.Language, error) {
	var langs []usecases.Language
	langsMap, err := repo.Handler.GetAll(repo.collection)
	if err != nil {
		return langs, err
	}

	mapstructure.Decode(langsMap, &langs)
	return langs, err
}

// NewFaqDBHandler creates a new handler for the faq
func NewFaqDBHandler(dbHandler DBHandler, collection, path string) *KBHandler {

	kbHandler := new(KBHandler)
	kbHandler.Handler = dbHandler
	kbHandler.collection = collection
	return kbHandler
}

// KnowledgeBase is the implementation that returns all the faq of the knowledge base
func (repo *KBHandler) KnowledgeBase() ([]domain.Faq, error) {
	var repFaqArray []repositoryFaqWithID
	faqs, err := repo.Handler.GetAll(repo.collection)
	if err != nil {
		return nil, err
	}

	// deconding the map to my type `repositoryFaqWithID`
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
	path := "isTrained"
	return repo.Handler.ChangeBool(repo.collection, ID, path, newStatus)
}

// AddFaq adds a new faq
func (repo *KBHandler) AddFaq(faq domain.Faq) error {
	faqMap := map[string]interface{}{
		"MainExample":      structs.Map(faq.MainExample),
		"Answers":          structs.Map(faq.Answers),
		"IsTrained":        structs.Map(faq.IsTrained),
		"TrainingExamples": structs.Map(faq.TrainingExamples),
	}
	return repo.Handler.Store(repo.collection, &faqMap)

}

// DeleteFaq deletes a faq
func (repo *KBHandler) DeleteFaq(ID string) error {
	return repo.Handler.Delete(repo.collection, ID)
}
