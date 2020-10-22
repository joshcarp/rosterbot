package database

import (
	"context"

	"cloud.google.com/go/firestore"
)

type Database interface {
	Get(collection, name string, i interface{}) error
	Set(collection, name string, i interface{}) error
	Filter(collection, name string) ([]byte, error)
}

type Firestore struct {
	firestore.Client
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

func (f Firestore) Filter(collection string, op string, filters map[string]string) ([]*firestore.DocumentSnapshot, error) {
	col := f.Collection(collection)
	q := col.Query
	for filter, val := range filters {
		q = q.Where(filter, op, val)
	}
	iter := q.Documents(context.Background())
	if iter == nil {
		return nil, nil
	}
	return iter.GetAll()
}
