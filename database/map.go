package database

import (
	"encoding/json"

	"github.com/joshcarp/rosterbot/command"
)

type Map map[string][]byte

func NewMap() (Map, error) {
	return make(map[string][]byte, 0), nil
}

func (f Map) Get(collection, name string, i interface{}) error {
	return json.Unmarshal(f[collection+"-"+name], i)
}

func (f Map) Set(collection, name string, i interface{}) error {
	b, err := json.Marshal(i)
	if err != nil {
		return err
	}
	f[collection+"-"+name] = b
	return nil
}

func (f Map) Filter(collection string, op, prefix string, filters map[string]interface{}) ([]command.RosterPayload, error) {
	a := make([]command.RosterPayload, 0)
	for _, e := range f {
		var w command.RosterPayload
		err := json.Unmarshal(e, &w)
		if err != nil {
			continue
		}
		var docontinue = false
		for filt, val := range filters {
			if w.Time.Complete[filt] != val {
				docontinue = true
				break
			}
		}
		if docontinue {
			continue
		}
		a = append(a, w)
	}
	return a, nil
}

func (f Map) Delete(collection, name string) error {
	delete(f, collection+"-"+name)
	return nil
}
