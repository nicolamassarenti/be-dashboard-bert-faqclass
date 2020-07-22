package usecases

import (
	"github.com/NicolaMassarenti/be-dashboard-bert-faqclass/domain"
)

//Logger is the interface that manages the logs
type Logger interface {
	Log(message string) error
}

//KnowledgeBaseInteractor is the object that manages the interactions
type KnowledgeBaseInteractor struct {
	FaqRepository domain.FaqRepository
}
