package repositories

import (
	"github.com/nicolamassarenti/be-dashboard-bert-faqclass/src/domain"
	"time"
)

// domainFaqToMapStringInterface returns a faq as map[string]interface
func domainFaqToMapStringInterface(faq domain.Faq) map[string]interface{}{
	return map[string]interface{}{
		"MainExample":      faq.MainExample,
		"Answers":          faq.Answers,
		"IsTrained":        faq.IsTrained,
		"TrainingExamples": faq.TrainingExamples,
		"UpdateDate": 		time.Now().Format(time.RFC3339),
	}
}

// domainKeywordToMapStringInterface returns a keyword as map[string]interface
func domainKeywordToMapStringInterface(keyword domain.Keyword) map[string]interface{}{
	return map[string]interface{}{
		"DisplayText": keyword.DisplayText,
		"UpdateDate":  time.Now().Format(time.RFC3339),
	}
}