package glug

import (
	"net/http"
)

type Connection struct {
	Response http.ResponseWriter
	Request  *http.Request
	Handler  HandleFunc
}

func NewConnection(resp http.ResponseWriter, req *http.Request) *Connection {
	conn := Connection{
		Response: resp,
		Request:  req,
	}
	return &conn
}
