package glug

import (
	"sort"
	"strings"
)

type Segment interface {
	// 路径的 一段，即/path/的path部分
	Match(segement string) bool
	Init(options ...interface{})
	Call(conn *Connection)
}

type NormalSegment struct {
	segement string
}

func (this *NormalSegment) Match(segment string) bool {
	return this.segment == segment
}

func (this *NormalSegment) Init(options ...interface{}) {

}

func (this *NormalSegment) Call(conn *Connection) {

}

type Node struct {
	Segment  *Segment
	Children []*Node
}

func (this *Node) Match(segement string) bool {
	//TODO
	return true
}

func NewNode(segement string) *Node {
	//TODO
	seg := NormalSegement{segement: segement}
	node := Node{Segment: &seg, Children: make([]*Node, 0)}
	return &node
}

type PathTree struct {
	Root *Node
}

func (this *PathTree) Add(path string) {
	segements := strings.Split(path, "/")
	root := this.Root
	for _, segement := range segements {
		children := root.Children
		index := sort.Search(len(children), func(i int) bool {
			node := children[i]
			return node.Match(segement)
		})
		if index == len(children) {
			//没有找到
			node := NewNode(segement)
			root.Children = append(root.Children, node)
			root = node
		} else {
			root = children[index]
		}
	}
}

func (this *PathTree) Match(conn *Connection, path string) bool {
	root := this.Root
	segements := strings.Split(path, "/")
	finded := true
	allowed := make([]*Segment, 0)
	for _, segement := range segements {
		children := root.Children
		index := sort.Search(len(children), func(i int) bool {
			node := children[i]
			return node.Match(segement)
		})
		if index == len(children) {
			finded = false
			break
		}
		root = children[index]
		allowed = append(allowed, root.Segment)
	}
	if finded {
		for _, seg := range allowed {
			seg.Call(conn)
		}
	}
	return finded
}
