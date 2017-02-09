package glug

import (
	"fmt"
	// "io/ioutil"
	"net/http"
	// "sort"
	"testing"
)

type TestRouter struct {
	GlugRouter
}

func TestInterface(t *testing.T) {
	tr := new(TestRouter)
	tr.Init()
	tr.Use(tr.Match)
	tr.Use(Logger)
	tr.Use(tr.Dispatch)
	tr.Get("/login", func(conn *Connection) {
		fmt.Println("login +++++++++++++")
		conn.Sendresp(500, conn.Request.Header, []byte("hello"))
	})
	tr.Get("/login/:name/", func(conn *Connection) {
		fmt.Println("login name +++++++++++++++++")
		name := conn.PathParams["name"]
		conn.Sendresp(200, conn.Request.Header, []byte(name))
	})
	tr2 := new(TestRouter)
	tr2.Init()
	tr2.Use(tr2.Match)
	tr2.Use(tr2.Dispatch)
	tr2.Get("/hello", func(conn *Connection) {
		fmt.Println("hello +++++++++++++++")
	})
	tr.Forward("/logout/", tr2)
	go http.ListenAndServe(":8083", tr)
	http.Get("http://localhost:8083/logout/hello")
	http.Get("http://localhost:8083/login/hello")
	http.Get("http://localhost:8083/login/")
}
