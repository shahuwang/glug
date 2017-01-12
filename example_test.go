package glug

import (
	"fmt"
	// "reflect"
	"io/ioutil"
	"net/http"
	"testing"
)

type TestRouter struct {
	GlugRouter
}

func TestInterface(t *testing.T) {
	tr := NewRouter()
	tr.Use(tr.Match)
	tr.Use(tr.Dispatch)
	tr.Get("/login", func(conn *Connection) {
		conn.Response.Write([]byte("loginxxxxxxx15\n\r"))
	})
	go http.ListenAndServe(":9083", tr)
	resp, err := http.Get("http://localhost:9083/login")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
