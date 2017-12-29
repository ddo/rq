package rq

import (
	"encoding/json"
)

// NewFromJSON create Rq object from json
func NewFromJSON(data []byte) (r *Rq, err error) {
	r = &Rq{}
	err = json.Unmarshal(data, r)
	if r.URL == "" || r.Method == "" {
		r = nil
	}
	return
}

// JSONify returns Rq object in json
func (r *Rq) JSONify() ([]byte, error) {
	return json.Marshal(r)
}
