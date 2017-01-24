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
	Context       map[string]interface{} // 用于存放每个glug函数插入的数据
	registerFuncs []RegisterFunc
}

func NewConnection(resp http.ResponseWriter, req *http.Request) *Connection {
	conn := Connection{
		response:      resp,
		Request:       req,
		registerFuncs: make([]RegisterFunc, 0),
		PathParams:    make(map[string]string),
		Context:       make(map[string]interface{}),
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

func (conn *Connection) Register(fun RegisterFunc) {
	conn.registerFuncs = append(conn.registerFuncs, fun)
}

func (conn *Connection) runBeforeSend(status int, header http.Header, body []byte) {
	resp := &Resp{
		Conn:   conn,
		Status: status,
		Header: header,
		Body:   body,
	}
	for _, fun := range conn.registerFuncs {
		fun(resp)
	}
}
