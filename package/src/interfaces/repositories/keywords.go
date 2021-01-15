package repositories

import (
	"github.com/mitchellh/mapstructure"
	"github.com/nicolamassarenti/be-dashboard-bert-faqclass/src/domain"
)

// Add creates a new keyword
func (repo *KeywordsHandler) Add(keyword domain.Keyword) error {
	keywordMapStringInterface := domainKeywordToMapStringInterface(keyword)
	return repo.Handler.Add(repo.collection, &keywordMapStringInterface)
}

// Delete deletes a keyword
func (repo *KeywordsHandler) Delete(ID string) error {
	return repo.Handler.Delete(repo.collection, ID)
}

// Keywords returns all the keywords
func (repo *KeywordsHandler) Keywords() (keywords []domain.Keyword, err error) {
	var repositoryKeywords []repositoryKeyword
	keywordsList, err := repo.Handler.GetAll(repo.collection)
	if err != nil {
		return nil, err
	}

	// decoding the map to my type `repositoryFaq`
	err = mapstructure.Decode(keywordsList, &repositoryKeywords)
	if err != nil{
		return nil, err
	}

	for _, repKeyword := range repositoryKeywords {
		keywords = append(
			keywords,
			domain.Keyword{
				ID:          repKeyword.ID,
				DisplayText: repKeyword.Data.DisplayText,
			},
		)
	}

	return keywords, nil
}

func (repo *KeywordsHandler) Update(ID string, keyword domain.Keyword) error {
	keywordMap := domainKeywordToMapStringInterface(keyword)
	return repo.Handler.Update(repo.collection, ID, keywordMap)
}
