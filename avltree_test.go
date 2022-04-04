package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAVLTree_Append(t *testing.T) {
	avlTree := NewAVLTree(1, 2, 3, 4)
	assert.Equal(t, V(2), avlTree.root.value)
}
