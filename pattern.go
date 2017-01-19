package glug

import (
	// "fmt"
	// "sort"
	"strings"
)

type Node struct {
	Segment  Segment
	Children []*Node
	origin   string // 存放原始的字符串
}

func (this *Node) Match(conn *Connection, segment string) bool {
	//TODO
	return this.Segment.Match(conn, segment)
}

type PathTree struct {
	Root *Node
	segs []SegMatchFunc
}

func NewPathTree() *PathTree {
	segs := []SegMatchFunc{VariantMatchFunc, NormalMatchFunc}
	node := &Node{Children: make([]*Node, 0), origin: ""}
	return &PathTree{Root: node, segs: segs}
}

func (this *PathTree) Merge(path string, tree *PathTree) {
	//在路径path之后将tree并入
	node := this.ensureNode(path)
	node.Children = tree.Root.Children
	this.segs = append(this.segs, tree.segs...)
}

func (this *PathTree) Add(path string, handle HandleFunc) {
	node := this.ensureNode(path)
	node.Segment.AddHandle(handle)
}

func (this *PathTree) ensureNode(path string) *Node {
	segments := strings.Split(path, "/")
	root := this.Root
	for _, segment := range segments {
		if segment == "" {
			continue
		}
		children := root.Children
		var node *Node
		matched := false
		for _, node = range children {
			if node.origin == segment {
				matched = true
				break
			}
		}
		if !matched {
			node = this.newNode(segment)
			root.Children = append(root.Children, node)
			root = node
		} else {
			root = node
		}
	}
	return root
}

func (this *PathTree) Match(conn *Connection, path string) bool {
	root := this.Root
	segments := strings.Split(path, "/")
	var final Segment
	finded := true
	for _, segment := range segments {
		if segment == "" {
			continue
		}
		children := root.Children
		var node *Node
		matched := false
		for _, node = range children {
			if node.Match(conn, segment) {
				matched = true
				break
			}
		}
		if !matched {
			// 确保每个segment都是匹配上了才call
			finded = false
			break
		}
		root = node
		final = node.Segment
	}
	if finded {
		final.Call(conn)
	}
	return finded
}

func (this *PathTree) newNode(segment string) *Node {
	for _, fun := range this.segs {
		match, seg := fun(segment)
		if match {
			return &Node{Segment: seg, Children: make([]*Node, 0), origin: segment}
		}
	}
	return nil
}
