package infrastructure

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/NicolaMassarenti/be-dashboard-bert-faqclass/src/domain"
	"google.golang.org/api/iterator"
)

// FirestoreHandler is the struct that has the client to firestore
type FirestoreHandler struct {
	Client  *firestore.Client
	Context context.Context
}

// NewFirestoreHandler creates a new firestore handler, with a client
func NewFirestoreHandler(projectID string) *FirestoreHandler {
	// Get a Firestore client.
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// defer client.Close()

	firestoreHandler := new(FirestoreHandler)
	firestoreHandler.Client = client
	firestoreHandler.Context = ctx

	return firestoreHandler
}

// GetAll returns all the documents of a collection
func (handler *FirestoreHandler) GetAll() ([]domain.Faq, error) {
	iter := handler.Client.Collection("Faq").Documents(handler.Context)

	var faqs []domain.Faq

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var faq domain.Faq

		err = doc.DataTo(&faq)
		if err != nil {
			return nil, err
		}

		faqs = append(faqs, faq)
	}

	return faqs, nil
}

// Get returns a specific faq
func (handler *FirestoreHandler) Get(ID string) (domain.Faq, error) {

	iter := handler.Client.Collection("Faq").Where("ID", "==", ID).Documents(handler.Context)
	defer iter.Stop()

	var faq domain.Faq

	doc, err := iter.Next()
	if err != nil {
		return domain.Faq{}, err
	}

	err = doc.DataTo(&faq)
	if err != nil {
		return domain.Faq{}, err
	}

	return faq, nil
}

// ChangeBool changes the bool value of a document
func (handler *FirestoreHandler) ChangeBool(ID, path string, value bool) error {
	iter := handler.Client.Collection("Faq").Where("ID", "==", ID).Documents(handler.Context)
	defer iter.Stop()

	doc := handler.Client.Doc(ID)

	_, err := doc.Update(handler.Context, []firestore.Update{
		{Path: path, Value: value},
	})

	return err
}

// Store adds a new faq
func (handler *FirestoreHandler) Store(faq *domain.Faq) error {

	_, _, err := handler.Client.Collection("Faq").Add(handler.Context, faq)

	return err
}

// Delete deletes an Faq
func (handler *FirestoreHandler) Delete(ID string) error {
	_, err := handler.Client.Doc(ID).Delete(handler.Context)

	return err
}
