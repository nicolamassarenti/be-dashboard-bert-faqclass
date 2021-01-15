package repositories

import (
	"github.com/mitchellh/mapstructure"
	"github.com/nicolamassarenti/be-dashboard-bert-faqclass/src/domain"
)

// KnowledgeBase returns a list the knowledge base, so all faqs
func (repo *KBHandler) KnowledgeBase() ([]domain.Faq, error) {
	var repositoryFaqs []repositoryFaq
	faqs, err := repo.Handler.GetAll(repo.collection)
	if err != nil {
		return nil, err
	}

	// decoding the map to my type `repositoryFaq`
	err = mapstructure.Decode(faqs, &repositoryFaqs)
	if err != nil {
		return nil, err
	}

	// Returning data as array of domain faqs
	var kb []domain.Faq
	for _, repFaq := range repositoryFaqs {
		kb = append(
			kb,
			domain.Faq{
				ID:               repFaq.ID,
				MainExample:      repFaq.Data.MainExample,
				Answers:          repFaq.Data.Answers,
				IsTrained:        repFaq.Data.IsTrained,
				TrainingExamples: repFaq.Data.TrainingExamples,
			},
		)
	}

	return kb, nil
}

// Faq returns the faq that matches `ID`
func (repo *KBHandler) Faq(ID string) (faq domain.Faq, err error) {

	faqMap, err := repo.Handler.Get(repo.collection, ID)
	if err != nil {
		return
	}

	err = mapstructure.Decode(faqMap, &faq)
	return
}

// ChangeTrainingStatus changes the "isTrained" bool of a faq
func (repo *KBHandler) ChangeTrainingStatus(ID string, newStatus bool) error {
	fieldToChange := "IsTrained"
	return repo.Handler.ChangeBool(repo.collection, ID, fieldToChange, newStatus)
}

// AddFaq adds a new faq
func (repo *KBHandler) AddFaq(faq domain.Faq) error {
	faqMapStringInterface := domainFaqToMapStringInterface(faq)
	return repo.Handler.Add(repo.collection, &faqMapStringInterface)

}

// DeleteFaq deletes a faq
func (repo *KBHandler) DeleteFaq(ID string) error {
	return repo.Handler.Delete(repo.collection, ID)
}

// Update updates a faq
func (repo *KBHandler) Update(ID string, faq domain.Faq) error {
	faqMapStringInterface := domainFaqToMapStringInterface(faq)
	return repo.Handler.Update(repo.collection, ID, faqMapStringInterface)
}
