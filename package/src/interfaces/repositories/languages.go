package repositories

import "github.com/nicolamassarenti/be-dashboard-bert-faqclass/src/usecases"

// Languages returns all the languages
func (repo *LanguagesHandler) Languages() ([]usecases.Language, error) {
	langsMap, err := repo.Handler.GetAll(repo.collection)

	languages := make([]usecases.Language, len(langsMap))
	for idx, lang := range langsMap{
		data, _ := lang["faq"].(map[string]interface{})
		languages[idx] = usecases.Language{
			IsoName: data["IsoName"].(string),
			DisplayName: data["DisplayName"].(string),
		}
	}

	return languages, err
}
