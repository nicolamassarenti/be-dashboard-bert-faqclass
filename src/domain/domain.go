package domain

// FaqRepository is the interface
type FaqRepository interface {
	KnowledgeBase() ([]Faq, error)
	Faq(ID string) (Faq, error)
	ChangeTrainingStatus(ID string, newStatus bool) error
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
