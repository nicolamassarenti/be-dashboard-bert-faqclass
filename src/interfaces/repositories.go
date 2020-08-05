package interfaces

import (
	"github.com/NicolaMassarenti/be-dashboard-bert-faqclass/src/domain"
)

// FaqDBHandler is the interface for the handler of the DB for the FAQ
type FaqDBHandler interface {
	Store(collection string, faq *domain.Faq) error
	Get(collection string, ID string) (domain.Faq, error)
	GetAll(collection string) ([]domain.Faq, error)
	ChangeBool(collection string, ID, path string, value bool) error
	Delete(collection string, ID string) error
}

// FaqDBRepo is the object for the faq db handler
type FaqDBRepo struct {
	Handler    FaqDBHandler
	collection string
}

// KBHandler is the handler for the FAQ
type KBHandler FaqDBRepo

// NewFaqDBHandler creates a new handler for the faq
func NewFaqDBHandler(dbHandler FaqDBHandler, collection string) *KBHandler {

	kbHandler := new(KBHandler)
	kbHandler.Handler = dbHandler
	kbHandler.collection = collection
	return kbHandler
}

// KnowledgeBase is the implementation that returns all the faq of the knowledge base
func (repo *KBHandler) KnowledgeBase() ([]domain.Faq, error) {

	return repo.Handler.GetAll(repo.collection)
}

// Faq is the implementation that returns a specific ID
func (repo *KBHandler) Faq(ID string) (faq domain.Faq, err error) {

	return repo.Handler.Get(repo.collection, ID)
}

// ChangeTrainingStatus changes the "isTrained" bool of a Faq
func (repo *KBHandler) ChangeTrainingStatus(ID string, newStatus bool) error {
	path := "isTrained"
	return repo.Handler.ChangeBool(repo.collection, ID, path, newStatus)
}

// AddFaq adds a new faq
func (repo *KBHandler) AddFaq(faq domain.Faq) error {

	return repo.Handler.Store(repo.collection, &faq)

}

// DeleteFaq deletes a faq
func (repo *KBHandler) DeleteFaq(ID string) error {
	return repo.Handler.Delete(repo.collection, ID)
}
