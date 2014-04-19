package main

import (
	rethink "github.com/dancannon/gorethink"
	//"github.com/martini-contrib/sessionauth"
	"time"
)

type Todo struct {
	Id        string    `form:"id" gorethink:"id,omitempty" json:"id"`
	UserId    string    `form:"userid" gorethink:"user_id" json: "user_id"`
	Body      string    `form:"body" gorethink:"body" json:"body"`
	Completed bool      `form:"completed" gorethink:"completed" json:"completed"`
	Created   time.Time `json:"created_at"`
}

func (t *Todo) isCompleted() bool {
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
		if err := row.Scan(&t); err != nil {
			return err
		}
	}
	return nil
}

type TodoList struct {
}
