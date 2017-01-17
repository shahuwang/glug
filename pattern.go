package glug

import (
	// "fmt"
	"sort"
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
		index := sort.Search(len(children), func(i int) bool {
			node := children[i]
			return node.origin == segment
		})
		if index >= len(children) {
			//没有找到
			node := this.newNode(segment)
			root.Children = append(root.Children, node)
			root = node
		} else {
			root = children[index]
		}
	}
	return root
}

func (this *PathTree) Match(conn *Connection, path string) bool {
	root := this.Root
	segments := strings.Split(path, "/")
	finded := true
	allowed := make([]Segment, 0)
	for _, segment := range segments {
		if segment == "" {
			continue
		}
		children := root.Children
		index := sort.Search(len(children), func(i int) bool {
			node := children[i]
			return node.Match(conn, segment)
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

func (this *PathTree) newNode(segment string) *Node {
	for _, fun := range this.segs {
		match, seg := fun(segment)
		if match {
			return &Node{Segment: seg, Children: make([]*Node, 0), origin: segment}
		}
	}
	return nil
}
