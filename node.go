package gorouter

import "net/http"

type Node struct {
	idx      int
	slug     string
	nodes    []*Node
	handlers map[string]http.HandlerFunc
}

func newRootNode() *Node {
	return &Node{
		idx:      0,
		slug:     "/",
		nodes:    []*Node{},
		handlers: make(map[string]http.HandlerFunc),
	}
}

func newNode(slugs []string, idx int) *Node {
	return &Node{
		idx:      idx,
		slug:     slugs[idx-1],
		nodes:    []*Node{},
		handlers: make(map[string]http.HandlerFunc),
	}
}

func (n *Node) setHandler(method string, handler http.HandlerFunc) {
	n.handlers[method] = handler
}

func (n *Node) match(slugs []string) bool {
	if n.idx == 0 {
		return true
	}
	return slugs[n.idx-1] == n.slug
}

func (n *Node) find(slugs []string) *Node {
	if n.match(slugs) {
		if n.idx == len(slugs) {
			return n
		}
		for _, cn := range n.nodes {
			if found := cn.find(slugs); found != nil {
				return found
			}
		}
	}
	return nil
}

func (n *Node) add(slugs []string, method string, handler http.HandlerFunc) bool {
	if len(slugs) == n.idx {
		n.setHandler(method, handler)
		return true
	}
	for _, cn := range n.nodes {
		if cn.match(slugs) {
			return cn.add(slugs, method, handler)
		}
	}
	nn := newNode(slugs, n.idx+1)
	n.nodes = append(n.nodes, nn)
	if len(slugs) == nn.idx {
		nn.setHandler(method, handler)
		return true
	}
	return nn.add(slugs, method, handler)
}
