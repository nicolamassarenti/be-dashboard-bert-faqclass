package domain

// FaqRepository is the interface
type FaqRepository interface {
	KnowledgeBase() ([]Faq, error)
	Faq(ID string) (Faq, error)
	ChangeTrainingStatus(ID string) error
	AddFaq(Faq) error
	DeleteFaq(ID string) error
}

// Faq contains the data that define a F.A.Q.
type Faq struct {
	ID               string
	MainExample      string
	Answers          []Answer
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
	Lang   string
	Answer string
}

// Answer returns the answer in a specific language and an "ok" boolean that is false if doesn't exists a language
func (faq *Faq) Answer(lang string) (response string, ok bool) {
	for _, answer := range faq.Answers {
		if answer.Lang == lang {
			return answer.Answer, true
		}
	}
	return "", false
}

// ChangeTrainingFaq changes the training status of a Faq
func (faq *Faq) ChangeTrainingFaq() {
	faq.IsTrained = !faq.IsTrained
}
