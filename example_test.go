package glug

import (
	"fmt"
	// "reflect"
	"net/http"
	"testing"
)

type TestRouter struct {
	GlugRouter
}

func TestInterface(t *testing.T) {
	tr := NewRouter()
	fmt.Printf("%+v\n", tr)
	tr.Use(tr.Match)
	tr.Use(tr.Dispatch)
	tr.Get("/login", func(conn *Connection) {
		conn.Response.Write([]byte("loginxxxxxxx"))
	})
	er := http.ListenAndServe(":9083", tr)
	fmt.Println(er)
	resp, err := http.Get("http://localhost:9083")
	if err != nil {
		fmt.Println(err)
	}
	var body []byte
	_, err = resp.Body.Read(body)
	fmt.Println(err)
	fmt.Println(string(body))
}
