package webservice

import "github.com/nicolamassarenti/be-dashboard-bert-faqclass/src/usecases"

func webserviceKeywordToUsecaseKeyword(keyword Keyword) usecases.Keyword {
	return usecases.Keyword{
		DisplayText: keyword.DisplayText,
	}
}

func webserviceFaqToUsecaseFaq(faq Faq) usecases.Faq {
	var answers []usecases.Answer
	for k, v := range faq.Answers {
		answers = append(answers, usecases.Answer{Language: k, Answers: v})
	}

	var trainingExamples []usecases.TrainingExample
	for k, v := range faq.Examples {
		trainingExamples = append(trainingExamples, usecases.TrainingExample{Language: k, Examples: v})
	}

	return usecases.Faq{
		MainExample:      faq.MainQuestion,
		Answers:          answers,
		IsTrained:        faq.Trained,
		TrainingExamples: trainingExamples,
	}
}

func usecaseFaqToWebserviceFaq(faq usecases.Faq) Faq {

	answers := make(map[string][]string)
	for _, ans := range faq.Answers {
		answers[ans.Language] = ans.Answers
	}

	trainingExamples := make(map[string][]string)
	for _, example := range faq.TrainingExamples {
		trainingExamples[example.Language] = example.Examples
	}


	return Faq{
		MainQuestion: faq.MainExample,
		Answers: answers,
		Trained: faq.IsTrained,
		Examples: trainingExamples,
	}
}

func usecaseLanguageToMapStringString(langs []usecases.Language) map[string]string {
	result := make(map[string]string)
	for _, lang := range langs {
		result[lang.IsoName] = lang.DisplayName
	}
	return result
}