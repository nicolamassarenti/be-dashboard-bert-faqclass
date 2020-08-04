package interfaces

import (
	"github.com/NicolaMassarenti/be-dashboard-bert-faqclass/src/domain"
)

// FaqDBHandler is the interface for the handler of the DB for the FAQ
type FaqDBHandler interface {
	Store(faq *domain.Faq) error
	Get(ID string) (domain.Faq, error)
	GetAll() ([]domain.Faq, error)
	ChangeBool(ID, path string, value bool) error
	Delete(ID string) error
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

	return repo.Handler.GetAll()
}

// Faq is the implementation that returns a specific ID
func (repo *KBHandler) Faq(ID string) (faq domain.Faq, err error) {

	return repo.Handler.Get(ID)
}

// ChangeTrainingStatus changes the "isTrained" bool of a Faq
func (repo *KBHandler) ChangeTrainingStatus(ID string, newStatus bool) error {
	path := "isTrained"
	return repo.Handler.ChangeBool(ID, path, newStatus)
}

// AddFaq adds a new faq
func (repo *KBHandler) AddFaq(faq domain.Faq) error {

	return repo.Handler.Store(&faq)

}

// DeleteFaq deletes a faq
func (repo *KBHandler) DeleteFaq(ID string) error {
	return repo.Handler.Delete(ID)
}
