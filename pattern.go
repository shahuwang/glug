package glug

import (
	"sort"
	"strings"
)

type Segment interface {
	// 路径的 一段，即/path/的path部分
	Match(segment string) bool
	Call(conn *Connection) bool
	AddHandle(handle HandleFunc)
}

type NormalSegment struct {
	segment string
	handle  HandleFunc
}

func (this *NormalSegment) Match(segment string) bool {
	return this.segment == segment
}

func (this *NormalSegment) Call(conn *Connection) bool {
	conn.Handler = this.handle
	return true
}

func (this *NormalSegment) AddHandle(handle HandleFunc) {
	this.handle = handle
}

type Node struct {
	Segment  Segment
	Children []*Node
}

func (this *Node) Match(segment string) bool {
	//TODO
	return this.Segment.Match(segment)
}

func NewNode(segment string) *Node {
	//TODO
	seg := NormalSegment{segment: segment}
	node := Node{Segment: &seg, Children: make([]*Node, 0)}
	return &node
}

type PathTree struct {
	Root *Node
}

func NewPathTree() *PathTree {
	node := NewNode("")
	return &PathTree{Root: node}
}

func (this *PathTree) Add(path string, handle HandleFunc) {
	segments := strings.Split(path, "/")
	root := this.Root
	for _, segment := range segments {
		children := root.Children
		index := sort.Search(len(children), func(i int) bool {
			node := children[i]
			return node.Match(segment)
		})
		if index == len(children) {
			//没有找到
			node := NewNode(segment)
			root.Children = append(root.Children, node)
			root = node
		} else {
			root = children[index]
		}
	}
	root.Segment.AddHandle(handle)
}

func (this *PathTree) Match(conn *Connection, path string) bool {
	root := this.Root
	segments := strings.Split(path, "/")
	finded := true
	allowed := make([]Segment, 0)
	for _, segment := range segments {
		children := root.Children
		index := sort.Search(len(children), func(i int) bool {
			node := children[i]
			return node.Match(segment)
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
