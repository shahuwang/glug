package glug

import (
	"fmt"
	// "reflect"
	"testing"
)

type TestRouter struct {
	GlugRouter
}

func TestInterface(t *testing.T) {
	tr := TestRouter{}
	fmt.Printf("%+v\n", tr)
	tr.BuildGlug(
		tr.Match,
		tr.Dispatch,
	)
}
