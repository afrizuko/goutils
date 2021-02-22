package vfs

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GenerateVFS(t *testing.T) {

	startAt := 1
	lst := newList(startAt)
	var str strings.Builder
	str.WriteString(strconv.Itoa(startAt))

	for i := 2; i <= 5; i++ {
		assert.NoError(t, lst.pushNumber(i))
		str.WriteString(strconv.Itoa(i))
	}
	res, err := goForward(lst)
	assert.NoError(t, err)
	assert.Equal(t, 2, res.number)

	val, err := goBack(res)
	assert.NoError(t, err)
	assert.Equal(t, 1, val.number)
	assert.Equal(t, str.String(), lst.String())
}

func Test_ForwardBackTraversal(t *testing.T) {
	lst := newList(1)
	for i := 2; i <= 5; i++ {
		assert.NoError(t, lst.pushNumber(i))
	}
}

func Test_ForwardTraversal(t *testing.T) {
	lst := newList(1)
	for i := 2; i <= 5; i++ {
		assert.NoError(t, lst.pushNumber(i))
	}

	lst, err := goForward(lst)
	assert.NoError(t, err)
	assert.Equal(t, 2, lst.number)
}

func Test_ErrorOnBackTraversal(t *testing.T) {
	lst := newList(1)
	for i := 2; i <= 5; i++ {
		assert.NoError(t, lst.pushNumber(i))
	}
	_, err := goBack(lst)
	assert.Error(t, err)
}

func Test_ErrorOnForwardTraversal(t *testing.T) {
	lst := newList(1)

	_, err := goForward(lst)
	assert.Error(t, err)
}
