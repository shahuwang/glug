package glug

import (
	"fmt"
	"net/http"
)

type Glug interface {
	Init(options ...interface{})
	Call(conn *Connection)
}

type Router interface {
	BuildGlug(glugs ...interface{})
	Init(options ...interface{}) Glug
	Call(conn *Connection)
	Match(conn *Connection)
	Dispatch(conn *Connection)
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type HandlerFunc func(params ...interface{})

type GlugRouter struct {
	Builder Builder
}

func (this *GlugRouter) BuildGlug(glugs ...interface{}) {
	builder := Builder{}
	builder.Init(glugs...)
	this.Builder = builder
}

func (this *GlugRouter) Init(options ...interface{}) {
	//TODO
}

func (this *GlugRouter) Call(conn *Connection) {
	this.Builder.Call(conn)
}

func (this *GlugRouter) Match(conn *Connection) {
	fmt.Println("============")
	//TODO
}

func (this *GlugRouter) Dispatch(conn *Connection) {
	//TODO
	conn.Response.Write([]byte("hello world"))
}

func (this *GlugRouter) Get(path string, fn HandlerFunc) {

}

func (this *GlugRouter) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	conn := NewConnection(resp, req)
	fmt.Println("xxxxxxx")
	this.Call(conn)
}
