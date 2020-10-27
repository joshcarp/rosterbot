package database

import (
	"context"

	"github.com/joshcarp/rosterbot/command"

	"cloud.google.com/go/firestore"
)

type Firestore struct {
	*firestore.Client
}

func NewFirestore(projectID string) (Firestore, error) {
	a, err := firestore.NewClient(context.Background(), projectID)
	if err != nil {
		return Firestore{}, err
	}
	return Firestore{Client: a}, nil
}

func (f Firestore) Get(collection, name string, i interface{}) error {
	snap, err := f.Collection(collection).Doc(name).Get(context.Background())
	if err != nil {
		return err
	}
	return snap.DataTo(i)
}

func (f Firestore) Set(collection, name string, i interface{}) error {
	_, err := f.Collection(collection).Doc(name).Set(context.Background(), i)
	return err
}

func (f Firestore) Filter(collection string, op, prefix string, filters map[string]interface{}) ([]command.RosterPayload, error) {
	col := f.Collection(collection)
	q := col.Query
	for filter, val := range filters {
		q = q.Where(prefix+filter, op, val)
	}
	iter := q.Documents(context.Background())
	if iter == nil {
		return nil, nil
	}
	docs, err := iter.GetAll()
	if err != nil {
		return nil, err
	}
	ret := make([]command.RosterPayload, 0, len(docs))
	for _, b := range docs {
		var payload command.RosterPayload
		b.DataTo(&payload)
		ret = append(ret, payload)
	}
	return ret, nil
}

func (f Firestore) Delete(collection, name string) error {
	_, err := f.Client.Collection(collection).Doc(name).Delete(context.Background())
	return err
}
