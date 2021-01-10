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

// KeywordRepository is the interface
type KeywordRepository interface {
	Add(keyword Keyword) error
	Delete(ID string) error
	Keywords() ([]Keyword, error)
	Update(ID string, keyword Keyword) error
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
	Language   string
	Answers []string
}

// Keyword is a keyword
type Keyword struct {
	ID string
	Name string
}
