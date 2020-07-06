package models

type Faq struct {
	id           string
	mainQuestion string
	trained      bool
	examples     []TrainingExample
}
