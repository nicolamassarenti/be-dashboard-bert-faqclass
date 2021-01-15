package repositories

import (
	"github.com/nicolamassarenti/be-dashboard-bert-faqclass/src/domain"
	"time"
)

// Returns the map[string]interface in the format of the database
func domainFaqToMapStringInterface(faq domain.Faq) map[string]interface{}{
	return map[string]interface{}{
		"MainExample":      faq.MainExample,
		"Answers":          faq.Answers,
		"IsTrained":        faq.IsTrained,
		"TrainingExamples": faq.TrainingExamples,
		"UpdateDate": 		time.Now().Format(time.RFC3339),
	}
}

// Returns the map[string]interface formatted as requested by the database
func domainKeywordToMapStringInterface(keyword domain.Keyword) map[string]interface{}{
	return map[string]interface{}{
		"DisplayText": keyword.DisplayText,
		"UpdateDate":  time.Now().Format(time.RFC3339),
	}
}