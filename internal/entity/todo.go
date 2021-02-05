package io

import (
	"encoding/json"

	"gopkg.in/mgo.v2/bson"
)

type Todo struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Title    string        `json:"title"  bson:"title"`
	Content  string        `json:"content"  bson:"content"`
	Complete bool          `json:"complete" bson:"complete"`
}

func (t Todo) String() string {
	b, err := json.Marshal(t)
	if err != nil {
		return "unsupported value type"
	}
	return string(b)
}
