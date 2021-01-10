package repositories

import (
	"github.com/mitchellh/mapstructure"
	"github.com/nicolamassarenti/be-dashboard-bert-faqclass/src/domain"
)

// Add creates a new keyword
func (repo *KeywordsHandler) Add(keyword domain.Keyword) error {
	keywordMap := getKeywordMapToAdd(keyword)
	return repo.Handler.Add(repo.collection, &keywordMap)
}

// Delete deletes a keyword
func (repo *KeywordsHandler) Delete(ID string) error {
	return repo.Handler.Delete(repo.collection, ID)
}

func (repo *KeywordsHandler) Keyword(ID string) (keyword domain.Keyword, err error) {

	keywordMap, err := repo.Handler.Get(repo.collection, ID)
	if err != nil {
		return
	}

	err = mapstructure.Decode(keywordMap, &keyword)
	return
}

func (repo *KeywordsHandler) Keywords() (keywords []domain.Keyword, err error) {
	var repFaqArray []repositoryKeywordWithID
	faqs, err := repo.Handler.GetAll(repo.collection)
	if err != nil {
		return nil, err
	}

	// decoding the map to my type `repositoryFaqWithID`
	err = mapstructure.Decode(faqs, &repFaqArray)
	if err != nil{
		return nil, err
	}

	for _, repKeyword := range repFaqArray {
		keywords = append(
			keywords,
			domain.Keyword{
				ID:               repKeyword.ID,
				Name:      			repKeyword.Keyword.Name,
			},
		)
	}

	return keywords, nil
}

func (repo *KeywordsHandler) Update(ID string, keyword domain.Keyword) error {
	keywordMap := getKeywordMapToAdd(keyword)
	return repo.Handler.Update(repo.collection, ID, keywordMap)
}
