// Building on top of the auth example is the todo app example.

// This from the martini-contrib sessionauth example,
// but this is using RethinkDB instead of sqlite3. For personal learning purposes only.
package main

import (
	"code.google.com/p/go.crypto/bcrypt"
	"fmt"
	rethink "github.com/dancannon/gorethink"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessionauth"
	"github.com/martini-contrib/sessions"
	"log"
	"net/http"
)

var (
	dbSession *rethink.Session
)

func init() {

	// Assumes there's a rethinkdb instance running locally with db called 'todo'
	// Db has table called "user" with.
	// "yelnil@example.com" with password "qwe"
	var dbError error
	dbSession, dbError = rethink.Connect(rethink.ConnectOpts{
		Address:  "localhost:28015",
		Database: "todo"})
	if dbError != nil {
		log.Fatalln(dbError.Error())
	}
}

func main() {
	store := sessions.NewCookieStore([]byte("secret123"))
	m := martini.Classic()
	m.Use(render.Renderer())

	store.Options(sessions.Options{MaxAge: 0})
	m.Use(sessions.Sessions("my_session", store))

	// Every request is bound with empty user. If there's a session,
	// that empty user is filled with appopriate data
	m.Use(sessionauth.SessionUser(GenerateAnonymousUser))
	sessionauth.RedirectUrl = "/login"
	sessionauth.RedirectParam = "next"

	m.Get("/", func(r render.Render) {
		r.HTML(200, "index", nil)
	})
	m.Get("/login", func(r render.Render) {
		r.HTML(200, "login", nil)
	})
	m.Get("/register", func(session sessions.Session, r render.Render) {
		if session.Get(sessionauth.SessionKey) != nil {
			r.HTML(200, "index", nil)
			return
		}
		r.HTML(200, "register", nil)
	})

	m.Post("/register", binding.Bind(MyUserModel{}), registerHandler)

	m.Post("/login", binding.Bind(MyUserModel{}), loginHandler)

	m.Get("/logout", sessionauth.LoginRequired, func(session sessions.Session, user sessionauth.User, r render.Render) {
		sessionauth.Logout(session, user)
		r.Redirect("/")
	})

	m.Run()
}
