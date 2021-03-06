package usecases

import (
	"github.com/nicolamassarenti/be-dashboard-bert-faqclass/src/domain"
)

// domainFaqToUsecaseFaq transforms a FAQ from the struct of domain to the struct of the usecases layer
func domainFaqToUsecaseFaq(faq domain.Faq) Faq {
	answers := make([]Answer, len(faq.Answers))
	for i, ans := range faq.Answers {
		answers[i] = Answer{ans.Language, ans.Answers}
	}

	trainingExamples := make([]TrainingExamples, len(faq.TrainingExamples))
	for i, example := range faq.TrainingExamples {
		trainingExamples[i] = TrainingExamples{example.Language, example.Examples}
	}
	return Faq{faq.ID, faq.MainExample, answers, faq.IsTrained, trainingExamples}
}

// usecaseFaqToDomainFaq transforms a FAQ from the struct of usecases to the one of the domain layer
func usecaseFaqToDomainFaq(faq Faq) domain.Faq {

	answers := make([]domain.Answers, len(faq.Answers))
	for i, ans := range faq.Answers {
		answers[i] = domain.Answers{Language: ans.Language, Answers: ans.Answers}
	}

	trainingExamples := make([]domain.TrainingExamples, len(faq.TrainingExamples))
	for i, example := range faq.TrainingExamples {
		trainingExamples[i] = domain.TrainingExamples{Language: example.Language, Examples: example.Examples}
	}
	return domain.Faq{
		MainExample:      faq.MainExample,
		Answers:          answers,
		IsTrained:        faq.IsTrained,
		TrainingExamples: trainingExamples,
	}
}
