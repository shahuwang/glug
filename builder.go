package glug

import ()

type GlugFunc func(*Connection) bool

type Builder struct {
	funcs []GlugFunc
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (this *Builder) Call(conn *Connection) bool {
	for _, callFunc := range this.funcs {
		if !callFunc(conn) {
			return false
		}
	}
	return true
}

func (this *Builder) Add(glug GlugFunc) {
	this.funcs = append(this.funcs, glug)
}
