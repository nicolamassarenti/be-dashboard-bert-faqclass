package usecases

import (
	"github.com/nicolamassarenti/be-dashboard-bert-faqclass/src/domain"
)

// faqFromDomainLayer transfroms a FAQ from the struct of domain to the struct of the usecases layer
func faqFromDomainLayer(faq domain.Faq) Faq {
	answers := make([]Answer, len(faq.Answers))
	for i, ans := range faq.Answers {
		answers[i] = Answer{ans.Lang, ans.Answer}
	}

	trainingExamples := make([]TrainingExample, len(faq.TrainingExamples))
	for i, example := range faq.TrainingExamples {
		trainingExamples[i] = TrainingExample{example.Language, example.Examples}
	}
	return Faq{faq.ID, faq.MainExample, answers, faq.IsTrained, trainingExamples}
}

// faqToDomainLayer transfroms a FAQ from the struct of usecases to the one of the domain layer
func faqToDomainLayer(faq Faq) domain.Faq {

	answers := make([]domain.Answer, len(faq.Answers))
	for i, ans := range faq.Answers {
		answers[i] = domain.Answer{ans.Lang, ans.Answer}
	}

	trainingExamples := make([]domain.TrainingExample, len(faq.TrainingExamples))
	for i, example := range faq.TrainingExamples {
		trainingExamples[i] = domain.TrainingExample{example.Language, example.Examples}
	}
	return domain.Faq{
		ID:               faq.ID,
		MainExample:      faq.MainExample,
		Answers:          answers,
		IsTrained:        faq.IsTrained,
		TrainingExamples: trainingExamples,
	}
}
