package vfs

import (
	"errors"
	"fmt"
	"strings"
)

type list struct {
	number int
	next   *list
	prev   *list
}

func newList(num int) *list {
	return &list{number: num}
}

func (l *list) pushNumber(n int) error {
	currNode := l
	for currNode.next != nil {
		currNode = currNode.next
	}
	currNode.next = newList(n)
	currNode.next.prev = currNode
	return nil
}

func (l *list) String() string {
	currNode := l
	var builder strings.Builder
	for currNode != nil {
		str := fmt.Sprintf("%+v", currNode.number)
		builder.WriteString(str)
		currNode = currNode.next
	}
	return builder.String()
}

func goBack(l *list) (*list, error) {
	if l.prev == nil {
		return nil, errors.New("Cannot traverse backwards")
	}
	return l.prev, nil
}

func goForward(l *list) (*list, error) {
	if l.next == nil {
		return nil, errors.New("Cannot traverse forward")
	}
	return l.next, nil
}
