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
		conn.Sendresp(500, conn.Request.Header, []byte("hello"))
	})
	go http.ListenAndServe(":8083", tr)
	resp, err := http.Get("http://localhost:8083/login?a=1")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%+v\n", resp.Header)
	fmt.Printf("read body is %s, err is %s, %s\n", body, err, resp.Status)
}
