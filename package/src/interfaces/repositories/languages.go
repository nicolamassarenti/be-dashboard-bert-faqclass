package repositories

import "github.com/nicolamassarenti/be-dashboard-bert-faqclass/src/usecases"

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
