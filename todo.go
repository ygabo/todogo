package main

import (
	rethink "github.com/dancannon/gorethink"
	"github.com/martini-contrib/sessionauth"
	"time"
)

type Todo struct {
	Id            string `form:"id" gorethink:"id,omitempty"`
	UserId        string `form:"userid" gorethink:"user_id"`
	Body          string `form:"body" gorethink:"body"`
	Completed     string `form:"completed" gorethink:"completed"`
	Created       time.Time
	authenticated bool `form:"-" gorethink:"-"`
}

func (t *Todo) isCompleted() {
	return t.Completed
}

func (t *Todo) toggleCompleted() {
	t.Completed = !t.Completed
}

func (t *Todo) GetById(id interface{}) error {

	row, err := rethink.Table("todo").Get(id).RunRow(dbSession)
	if err != nil {
		return err
	}
	if !row.IsNil() {
		if err := row.Scan(&u); err != nil {
			return err
		}
	}
	return nil
}
