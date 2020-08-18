package interfaces

import (
	"github.com/fatih/structs"

	"github.com/nicolamassarenti/be-dashboard-bert-faqclass/src/usecases"
)

func usecaseFaqToWebserviceFaq(faq usecases.Faq) Faq {
	answers := make([]Answer, len(faq.Answers))
	for i, ans := range faq.Answers {
		answers[i] = Answer{ans.Lang, ans.Answer}
	}

	trainingExamples := make(map[string][]string)
	for _, example := range faq.TrainingExamples {
		trainingExamples[example.Language] = example.Examples
	}
	return Faq{faq.MainExample, answers, faq.IsTrained, trainingExamples}
}

func webserviceFaqToUsecaseFaq(faq Faq) usecases.Faq {

	answers := make([]usecases.Answer, len(faq.Answers))
	for i, ans := range faq.Answers {
		answers[i] = usecases.Answer{ans.lang, ans.answer}
	}

	var trainingExamples []usecases.TrainingExample
	for k, v := range faq.Examples {
		trainingExamples = append(trainingExamples, usecases.TrainingExample{k, v})
	}

	return usecases.Faq{
		"",
		faq.MainQuestion,
		answers,
		faq.Trained,
		trainingExamples,
	}
}

func structToMap(item interface{}) map[string]interface{} {

	return structs.Map(item)
}

func languagesToFrontEnd(langs []usecases.Language) map[string]string {
	result := make(map[string]string)
	for _, lang := range langs {
		result[lang.IsoName] = lang.DisplayName
	}
	return result
}
