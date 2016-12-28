package glug

import (
	"sort"
	"strings"
)

type Segment interface {
	// 路径的一段，即/path/的path部分
	Match(segement string) bool
}

type Node struct {
	Segement *Segment
	Children []*Node
}

func (this *Node) Match(segement string) bool {
	//TODO
	return true
}

func NewNode(segement string) *Node {
	//TODO
	return new(Node)
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

func (this *PathTree) Match(path string) *Segment {
	root := this.Root
	segements := strings.Split(path, "/")
	finded := true
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
	}
	if finded {
		return root.Segement
	} else {
		return nil
	}
}
