package domain

// FaqRepository is the interface
type FaqRepository interface {
	KnowledgeBase() []Faq
	Faq(ID string) Faq
	ChangeTrainingFaq(ID string)
	AddFaq(Faq)
	DeleteFaq(ID string)
}

// Faq contains the data that define a F.A.Q.
type Faq struct {
	ID               string
	MainExample      string
	answers          []Answer
	IsTrained        bool
	TrainingExamples []TrainingExample
}

// TrainingExample contain the training examples of a specific language
type TrainingExample struct {
	Language string
	Examples []string
}

// Answer contains the answer in a language
type Answer struct {
	lang   string
	answer string
}

// Answer returns the answer in a specific language and an "ok" boolean that is false if doesn't exists a language
func (faq *Faq) Answer(lang string) (response string, ok bool) {
	for _, answer := range faq.answers {
		if answer.lang == lang {
			return answer.answer, true
		}
	}
	return "", false
}

// ChangeTrainingFaq changes the training status of a Faq
func (faq *Faq) ChangeTrainingFaq() {
	faq.IsTrained = !faq.IsTrained
}
