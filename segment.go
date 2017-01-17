package glug

import (
	"strings"
)

type Segment interface {
	// 路径的 一段，即/path/的path部分
	Match(conn *Connection, segment string) bool
	Call(conn *Connection) bool
	AddHandle(handle HandleFunc)
}

// 给定一个path小段，进行判断匹配，譬如VariantSegment还是NormalSegment
// 匹配上，返回true和生成的对象
// 否则，返回false和nil
type SegMatchFunc func(segment string) (bool, Segment)

func NormalMatchFunc(segment string) (bool, Segment) {
	if segment == "" {
		return false, nil
	}
	seg := NormalSegment{segment: segment}
	return true, &seg
}

func VariantMatchFunc(segment string) (bool, Segment) {
	if strings.HasPrefix(segment, ":") {
		seg := VariantSegment{segment: segment[1:]}
		return true, &seg
	}
	return false, nil
}

type NormalSegment struct {
	// 普通的路径
	segment string
	handle  HandleFunc
}

func (this *NormalSegment) Match(conn *Connection, segment string) bool {
	return this.segment == segment
}

func (this *NormalSegment) Call(conn *Connection) bool {
	//TODO 这里的处理逻辑值得商榷。。。。
	conn.Handler = this.handle
	return true
}

func (this *NormalSegment) AddHandle(handle HandleFunc) {
	this.handle = handle
}

type VariantSegment struct {
	// 变量路径，如/test/:name/
	segment string
	handle  HandleFunc
}

func (this *VariantSegment) Match(conn *Connection, segment string) bool {
	conn.PathParams[this.segment] = segment
	return true
}

func (this *VariantSegment) Call(conn *Connection) bool {
	conn.Handler = this.handle
	return true
}

func (this *VariantSegment) AddHandle(handle HandleFunc) {
	this.handle = handle
}
