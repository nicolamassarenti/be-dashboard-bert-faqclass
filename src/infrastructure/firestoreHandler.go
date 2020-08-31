package infrastructure

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
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

	firestoreHandler := new(FirestoreHandler)
	firestoreHandler.Client = client
	firestoreHandler.Context = ctx

	defer firestoreHandler.Client.Close()

	return firestoreHandler
}

// Add adds a new faq
func (handler *FirestoreHandler) Add(collection string, data *map[string]interface{}) error {

	_, _, err := handler.Client.Collection(collection).Add(handler.Context, data)

	return err
}

// ChangeBool changes the bool value of a document
func (handler *FirestoreHandler) ChangeBool(collection string, ID, path string, value bool) error {
	_, err := handler.Client.Doc(collection+"/"+ID).Update(handler.Context, []firestore.Update{
		{Path: path, Value: value},
	})

	return err
}

// Delete deletes an Faq
func (handler *FirestoreHandler) Delete(collection string, ID string) error {
	_, err := handler.Client.Doc(collection + "/" + ID).Delete(handler.Context)

	return err
}

// Get returns a specific faq
func (handler *FirestoreHandler) Get(collection string, ID string) (map[string]interface{}, error) {

	doc, err := handler.Client.Doc(collection + "/" + ID).Get(handler.Context)

	if err != nil {
		return nil, err
	}

	return doc.Data(), err
}

// GetAll returns all the documents of a collection
func (handler *FirestoreHandler) GetAll(collection string) ([]map[string]interface{}, error) {
	iter := handler.Client.Collection(collection).Documents(handler.Context)

	var faqs []map[string]interface{}

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		faqs = append(faqs, map[string]interface{}{
			"ID":  doc.Ref.ID,
			"faq": doc.Data(),
		})
	}

	return faqs, nil
}

// Update updates a Faq
func (handler *FirestoreHandler) Update(collection string, ID string, data map[string]interface{}) error {
	_, err := handler.Client.Doc(collection + "/" + ID).Set(handler.Context, data)

	return err
}