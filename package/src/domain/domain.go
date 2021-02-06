package domain

// FaqRepository defines the interface for the Faqs repository
type FaqRepository interface {
	AddFaq(faq Faq) error
	ChangeTrainingStatus(ID string, newStatus bool) error
	DeleteFaq(ID string) error
	Faq(ID string) (Faq, error)
	KnowledgeBase() ([]Faq, error)
	Update(ID string, faq Faq) error
}

// KeywordRepository defines the interface for the Keywords repository
type KeywordRepository interface {
	Add(keyword Keyword) error
	Delete(ID string) error
	Keywords() ([]Keyword, error)
	Update(ID string, keyword Keyword) error
}

// Faq is a F.A.Q
type Faq struct {
	ID               string
	MainExample      string
	Answers          []Answers
	IsTrained        bool
	TrainingExamples []TrainingExamples
}

// TrainingExamples defines Faq examples in a specific language
type TrainingExamples struct {
	Language string
	Examples []string
}

// Answers defines Faq answers in a specific language
type Answers struct {
	Language   string
	Answers []string
}

// Keyword is a keyword
type Keyword struct {
	ID          string
	DisplayText string
}
