/**
 *@Author luojunying
 *@Date 2022-01-16 22:19
 */
package Xy

import (
	"path"
	"strings"
)

//实现路由前缀树
type Tree struct {
	root *node
}

type node struct {
	isLast  bool
	segment string
	handler ControllerHandler
	Childs  []*node
}

func NewTree() *Tree {
	return &Tree{newNode()}
}

func newNode() *node {
	return &node{
		isLast:  false,
		segment: "",
		Childs:  []*node{},
	}
}

//url格式
func uriFormat(uri string) string {
	absoluteUri := path.Clean(uri)
	upperUri := strings.ToUpper(absoluteUri)
	return upperUri
}

func (tree *Tree) AddRouter(uri string, handler ControllerHandler) {
	root := tree.root
	uri = uriFormat(uri)
	segments := strings.Split(uri, "/")
	for index, segment := range segments {
		if !isWildSegment(segment) {
			segment = strings.ToUpper(segment)
		}
		isLast := index == len(segments)-1
		var objNode *node
		childNodes := root.filterChildNodes(segment)
		if len(childNodes) > 0 {
			for _, cnode := range childNodes {
				if cnode.segment == segment {
					objNode = cnode
					break
				}
			}
		}
		if objNode == nil {
			cnode := newNode()
			cnode.segment = segment
			if isLast {
				cnode.isLast = true
				cnode.handler = handler
			}
			root.Childs = append(root.Childs, cnode)
			objNode = cnode
		}
		root = objNode
	}
}

//是否 ":" 开头
func isWildSegment(segment string) bool {
	return strings.HasPrefix(segment, ":")
}

//过滤下一层满足segment规则的子节点
func (n *node) filterChildNodes(segment string) []*node {
	if len(n.Childs) == 0 {
		return nil
	}
	//如果segment是通配符, 所有下一层都满足
	if isWildSegment(segment) {
		return n.Childs
	}
	//
	nodes := make([]*node, 0, len(n.Childs))
	for _, cnode := range n.Childs {
		if isWildSegment(cnode.segment) {
			nodes = append(nodes, cnode)
		} else if cnode.segment == segment {
			nodes = append(nodes, cnode)
		}
	}
	return nodes
}

func (n *node) matchNode(uri string) *node {
	segments := strings.SplitN(uri, "/", 2)
	segment := segments[0]
	if !isWildSegment(segment) {
		segment = strings.ToUpper(segment)
	}
	//匹配下一层
	childs := n.filterChildNodes(segment)
	if childs == nil || len(childs) == 0 {
		return nil
	}
	//如果只有一个
	if len(segments) == 1 {
		for _, tn := range childs {
			if tn.isLast && tn.segment == segment {
				return tn
			}
		}
		return nil
	}
	//如果多个
	for _, tn := range childs {
		tnMatch := tn.matchNode(segments[1])
		if tnMatch != nil {
			return tnMatch
		}
	}
	return nil
}

func (tree *Tree) FindHandler(uri string) ControllerHandler {
	matchNode := tree.root.matchNode(uri)
	if matchNode != nil {
		return matchNode.handler
	}
	return nil
}
