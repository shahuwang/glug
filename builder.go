package glug

import (
	"reflect"
)

type PlugFunc func(*Connection)

type Builder struct {
	funcs []PlugFunc
}

func (this Builder) Init(options ...interface{}) Glug {
	for _, op := range options {
		typ := reflect.TypeOf(op).Kind()
		if typ == reflect.Struct {
			// 说明是实现了Glug接口的结构体
			this.funcs = append(this.funcs, op.(Glug).Call)
		}
		if typ == reflect.Func {
			// 说明只是如PlugFunc一样的函数
			function := reflect.ValueOf(op)
			this.funcs = append(this.funcs, func(conn *Connection) {
				function.Call([]reflect.Value{reflect.ValueOf(conn)})
			})

		}
	}
	return this
}

func (this Builder) Call(conn *Connection) {
	for _, callFunc := range this.funcs {
		callFunc(conn)
	}
}
