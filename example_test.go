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
		fmt.Println("================")
		conn.Response.Write([]byte("loginxxxxxxx15"))
	})
	go http.ListenAndServe(":8083", tr)
	resp, err := http.Get("http://localhost:8083/login")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("read body is %s, err is %s, %s\n", body, err, resp.Status)
}
