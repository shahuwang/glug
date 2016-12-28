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
	tr := &TestRouter{}
	fmt.Printf("%+v\n", tr)
	tr.BuildGlug(
		tr.Match,
		tr.Dispatch,
	)
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
