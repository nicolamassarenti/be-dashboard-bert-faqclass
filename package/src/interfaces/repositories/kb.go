package repositories

import (
	"github.com/mitchellh/mapstructure"
	"github.com/nicolamassarenti/be-dashboard-bert-faqclass/src/domain"
)

// KnowledgeBase is the implementation that returns all the faq of the knowledge base
func (repo *KBHandler) KnowledgeBase() ([]domain.Faq, error) {
	var repFaqArray []repositoryFaqWithID
	faqs, err := repo.Handler.GetAll(repo.collection)
	if err != nil {
		return nil, err
	}

	// decoding the map to my type `repositoryFaqWithID`
	err = mapstructure.Decode(faqs, &repFaqArray)
	if err != nil {
		return nil, err
	}

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

	err = mapstructure.Decode(faqMap, &faq)
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
