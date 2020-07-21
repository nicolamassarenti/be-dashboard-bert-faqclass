package domain

// FaqRepository is the interface
type FaqRepository interface {
	GetKnowledgeBase() []Faq
}

// Faq contains the data that define a F.A.Q.
type Faq struct {
	ID               string
	MainExample      string
	IsTrained        bool
	TrainingExamples []TrainingExample
}

// TrainingExample contain the training examples of a specific language
type TrainingExample struct {
	Language string
	Examples []string
}
