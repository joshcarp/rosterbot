package database

import "github.com/joshcarp/rosterbot/command"

type Database interface {
	Get(collection, name string, i interface{}) error
	Set(collection, name string, i interface{}) error
	Filter(collection string, op, prefix string, filters map[string]interface{}) ([]command.RosterPayload, error)
	Delete(collection, name string) error
}
