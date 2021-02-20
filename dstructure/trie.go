package dstructure

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"sync"
)

type Tree struct {
	method map[string]*Node
}

type Node struct {
	label    string
	handler  http.Handler
	children map[string]*Node
}

type Param struct {
	key   string
	value string
}

type Params []*Param

type Result struct {
	handler http.Handler
	params  Params
}

const (
	pathDelimiter     = "/"
	paramDelimiter    = ":"
	leftPtnDelimiter  = "["
	rightPtnDelimiter = "]"
	ptnWildCard       = "(.+)"
)

func NewTree() *Tree {
	return &Tree{
		method: map[string]*Node{
			http.MethodGet: {
				label:    "",
				handler:  nil,
				children: make(map[string]*Node),
			},
			http.MethodPost: {
				label:    "",
				handler:  nil,
				children: make(map[string]*Node),
			},
			http.MethodOptions: {
				label:    "",
				handler:  nil,
				children: make(map[string]*Node),
			},
			http.MethodPut: {
				label:    "",
				handler:  nil,
				children: make(map[string]*Node),
			},
			http.MethodDelete: {
				label:    "",
				handler:  nil,
				children: make(map[string]*Node),
			},
		},
	}
}

func (t *Tree) Insert(method, path string, handler http.Handler) error {

	currNode := t.method[method]
	if path == paramDelimiter {

		if len(currNode.children) != 0 && currNode.handler == nil {
			return errors.New("Root node exists")
		}

		currNode.label = path
		currNode.handler = handler
		return nil
	}

	for _, l := range t.removeEmpty(strings.Split(path, pathDelimiter)) {
		if nextNode, ok := currNode.children[l]; ok {
			currNode = nextNode
			continue
		}

		currNode.children[l] = &Node{
			label:    l,
			handler:  handler,
			children: make(map[string]*Node),
		}
		currNode = currNode.children[l]
	}

	return nil
}

func (t *Tree) removeEmpty(s []string) []string {
	var stripped []string
	for _, val := range s {
		if val != "" {
			stripped = append(stripped, val)
		}
	}
	return stripped
}

type regexCache struct {
	s sync.Map
}

func (rc *regexCache) Get(ptn string) (*regexp.Regexp, error) {
	v, ok := rc.s.Load(ptn)
	if ok {
		reg, ok := v.(*regexp.Regexp)
		if ok {
			return reg, nil
		}
		return nil, fmt.Errorf("the value of %q is invalid", ptn)
	}

	reg, err := regexp.Compile(ptn)
	if err != nil {
		return nil, err
	}

	rc.s.Store(ptn, reg)
	return reg, nil
}

var regC *regexCache

func (t *Tree) Search(method, path string) (*Result, error) {

	n, ok := t.method[method]
	if !ok || (len(n.label) == 0 && len(n.children) == 0) {
		return nil, errors.New("tree is empty")
	}

	// var params Params
	//TODO completion of this method
	return nil, nil

}
