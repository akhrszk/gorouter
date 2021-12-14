package gorouter

import (
	"regexp"
	"strings"
)

type Node struct {
	idx      int
	slug     string
	rgxp     *regexp.Regexp
	nodes    []*Node
	handlers map[string]Handler
}

func newRootNode() *Node {
	return &Node{
		idx:      0,
		slug:     "/",
		nodes:    []*Node{},
		handlers: make(map[string]Handler), // { [http.Method]: Handler }
	}
}

func newNode(slugs []string, idx int) *Node {
	slug := slugs[idx-1]
	if slug[:1] == ":" {
		if i := strings.Index(slug, "("); i != -1 {
			regex := slug[i+1 : len(slug)-1]
			return &Node{
				idx:      idx,
				slug:     slug[:i],
				rgxp:     regexp.MustCompile("^" + regex + "$"),
				nodes:    []*Node{},
				handlers: make(map[string]Handler),
			}
		}
	}
	return &Node{
		idx:      idx,
		slug:     slug,
		nodes:    []*Node{},
		handlers: make(map[string]Handler),
	}
}

func (n *Node) setHandler(method string, handler Handler) {
	n.handlers[method] = handler
}

// Returns matching result and path paramter in [key, value]
func (n *Node) match(slugs []string) (bool, []string) {
	if n.idx == 0 {
		return true, nil
	}
	slug := slugs[n.idx-1]
	if n.slug[:1] == ":" {
		b := n.rgxp == nil || n.rgxp.MatchString(slug)
		if !b {
			return false, nil
		}
		return true, []string{n.slug[1:], slug}
	}
	return slugs[n.idx-1] == n.slug, nil
}

func (n *Node) find(slugs []string, params Params) (*Node, Params) {
	if b, param := n.match(slugs); b {
		if len(param) > 0 {
			params[param[0]] = param[1]
		}
		if n.idx == len(slugs) {
			return n, params
		}
		for _, cn := range n.nodes {
			if node, params := cn.find(slugs, params); node != nil {
				return node, params
			}
		}
	}
	return nil, nil
}

// Add child Node
func (n *Node) add(slugs []string, method string, handler Handler) bool {
	if len(slugs) == n.idx {
		n.setHandler(method, handler)
		return true
	}
	for _, cn := range n.nodes {
		if cn.slug == slugs[n.idx] {
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
