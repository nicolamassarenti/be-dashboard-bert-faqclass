package domain

// FaqRepository is the interface
type FaqRepository interface {
	AddFaq(faq Faq) error
	ChangeTrainingStatus(ID string, newStatus bool) error
	DeleteFaq(ID string) error
	Faq(ID string) (Faq, error)
	KnowledgeBase() ([]Faq, error)
	Update(ID string, faq Faq) error
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
	Answer []string
}
