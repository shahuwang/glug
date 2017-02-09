package glug

import (
	// "fmt"
	"net/http"
	"reflect"
)

type Glug interface {
	Call(conn *Connection) bool
}

type Router interface {
	Call(conn *Connection) bool
	Match(conn *Connection) bool
	Dispatch(conn *Connection) bool
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type HandleFunc func(*Connection)

type GlugRouter struct {
	Builder  *Builder
	GetTree  *PathTree
	PostTree *PathTree
}

func (this *GlugRouter) Init() {
	this.GetTree = NewPathTree()
	this.PostTree = NewPathTree()
	this.Builder = NewBuilder()
}

func (this *GlugRouter) Use(glug GlugFunc) {
	this.Builder.Add(glug)
}

func (this *GlugRouter) Call(conn *Connection) bool {
	return this.Builder.Call(conn)
}

func (this *GlugRouter) Match(conn *Connection) bool {
	//TODO
	path := conn.Request.URL.Path
	switch conn.Request.Method {
	case "GET":
		if !this.GetTree.Match(conn, path) {
			http.NotFound(conn.response, conn.Request)
			return false
		}
	case "POST":
		if !this.PostTree.Match(conn, path) {
			http.NotFound(conn.response, conn.Request)
			return false
		}
	}
	return true
}

func (this *GlugRouter) Dispatch(conn *Connection) bool {
	//TODO
	if conn.Handler == nil {
		http.NotFound(conn.response, conn.Request)
	}
	conn.Handler(conn)
	return true
}

func (this *GlugRouter) Get(path string, handle HandleFunc) {
	this.GetTree.Add(path, handle)
}

func (this *GlugRouter) Post(path string, handle HandleFunc) {
	this.PostTree.Add(path, handle)
}

func (this *GlugRouter) Forward(path string, router Router) {
	r := reflect.ValueOf(router).Elem()
	gtree := r.FieldByName("GetTree").Interface().(*PathTree)
	ptree := r.FieldByName("PostTree").Interface().(*PathTree)
	this.GetTree.Merge(path, gtree)
	this.PostTree.Merge(path, ptree)
}

func (this *GlugRouter) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	conn := NewConnection(resp, req)
	this.Call(conn)
}
