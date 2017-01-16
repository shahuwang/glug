package glug

import (
	"net/http"
)

type RegisterFunc func(*Resp)

type Resp struct {
	Conn   *Connection
	Status int
	Header http.Header
	Body   []byte
}

type Connection struct {
	response      http.ResponseWriter
	Request       *http.Request
	Handler       HandleFunc
	PathParams    map[string]string      //譬如/test/:name/ 这种路径，会取得name的实际值，存放于这里
	Params        map[string]interface{} // 参数存放在这
	registerFuncs []RegisterFunc
}

func NewConnection(resp http.ResponseWriter, req *http.Request) *Connection {
	conn := Connection{
		response:      resp,
		Request:       req,
		registerFuncs: make([]RegisterFunc, 0),
		PathParams:    make(map[string]string),
		Params:        make(map[string]interface{}),
	}
	return &conn
}

func (conn *Connection) Sendresp(status int, header http.Header, body []byte) {
	conn.runBeforeSend(status, header, body)
	for k, v := range header {
		for _, i := range v {
			conn.response.Header().Add(k, i)
		}
	}
	conn.response.WriteHeader(status)
	conn.response.Write(body)
}

func (conn *Connection) runBeforeSend(status int, header http.Header, body []byte) {
	resp := &Resp{
		conn:   conn,
		status: status,
		header: header,
		body:   body,
	}
	for _, fun := range conn.registerFuncs {
		fun(resp)
	}
}
