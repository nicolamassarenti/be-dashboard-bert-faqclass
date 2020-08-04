package interfaces

import "github.com/NicolaMassarenti/be-dashboard-bert-faqclass/src/usecases"

func usecaseFaqToWebserviceFaq(faq usecases.Faq) Faq {
	answers := make([]Answer, len(faq.Answers))
	for i, ans := range faq.Answers {
		answers[i] = Answer{ans.Lang, ans.Answer}
	}

	trainingExamples := make(map[string][]string)
	for _, example := range faq.TrainingExamples {
		trainingExamples[example.Language] = example.Examples
	}
	return Faq{faq.ID, faq.MainExample, answers, faq.IsTrained, trainingExamples}
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
		faq.ID,
		faq.MainQuestion,
		answers,
		faq.Trained,
		trainingExamples,
	}
}
