package db

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

// DBHandler is the object that defines the db handler
type DBHandler struct {
	Client  *firestore.Client
	Context context.Context
}

// Handler creates a new firestore handler, with a client
func Handler(projectID string) *DBHandler {
	// Get a Firestore client.
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	firestoreHandler := new(DBHandler)
	firestoreHandler.Client = client
	firestoreHandler.Context = ctx

	return firestoreHandler
}

// Add adds a new document
func (handler *DBHandler) Add(collection string, data *map[string]interface{}) error {

	_, _, err := handler.Client.Collection(collection).Add(handler.Context, data)

	return err
}

// ChangeBool changes a bool value of a document
func (handler *DBHandler) ChangeBool(collection string, ID, path string, value bool) error {
	_, err := handler.Client.Doc(collection+"/"+ID).Update(handler.Context, []firestore.Update{
		{Path: path, Value: value},
	})

	return err
}

// Delete deletes a document
func (handler *DBHandler) Delete(collection string, ID string) error {
	_, err := handler.Client.Doc(collection + "/" + ID).Delete(handler.Context)

	return err
}

// Get returns a specific document
func (handler *DBHandler) Get(collection string, ID string) (map[string]interface{}, error) {

	doc, err := handler.Client.Doc(collection + "/" + ID).Get(handler.Context)

	if err != nil {
		return nil, err
	}

	return doc.Data(), err
}

// GetAll returns all the documents of a collection
func (handler *DBHandler) GetAll(collection string) ([]map[string]interface{}, error) {
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
			"data": doc.Data(),
		})
	}

	return faqs, nil
}

// Update updates a document
func (handler *DBHandler) Update(collection string, ID string, data map[string]interface{}) error {
	_, err := handler.Client.Doc(collection + "/" + ID).Set(handler.Context, data)

	return err
}