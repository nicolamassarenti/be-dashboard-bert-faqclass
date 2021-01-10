package repositories

import "github.com/nicolamassarenti/be-dashboard-bert-faqclass/src/domain"

// AddKeyword adds a new keyword
func (repo *KeywordsHandler) AddKeyword(keyword domain.Keyword) error {
	keywordMap := getKeywordMapToAdd(keyword)
	return repo.Handler.Add(repo.collection, &keywordMap)
}

// DeleteKeyword deletes a keyword
func (repo *KeywordsHandler) DeleteKeyword(ID string) error {
	return repo.Handler.Delete(repo.collection, ID)
}
