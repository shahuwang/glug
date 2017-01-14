package glug

import (
	"net/http"
)

type Connection struct {
	Response   http.ResponseWriter
	Request    *http.Request
	Handler    HandleFunc
	PathParams map[string]string      //譬如/test/:name/ 这种路径，会取得name的实际值，存放于这里
	Params     map[string]interface{} // 参数存放在这
}

func NewConnection(resp http.ResponseWriter, req *http.Request) *Connection {
	conn := Connection{
		Response:   resp,
		Request:    req,
		PathParams: make(map[string]string),
		Params:     make(map[string]interface{}),
	}
	return &conn
}
