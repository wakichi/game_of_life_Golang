package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeepCopyTrue(t *testing.T){
	width, height :=3,3
	var src =Cell{{1,2,3}, {6,5,4}, {7,8,9}}
	var dst Cell
	dst = make(Cell, height)
	for i := 0; i < height; i++ {
		dst[i] = make([]int, width)
	}
	DeepCopy(&dst, &src)
	assert.Equal(t,src, dst)
}

func TestMakeCell(t *testing.T) {
	height, width := 3,4
	actual := MakeCell(height, width)
	expected:=Cell{{0,0,0,0},{0,0,0,0},{0,0,0,0}} //height:3, width:4
	assert.Equal(t, actual, expected)
}

func TestDeepCopyWithMakeCell(t *testing.T) {
	height, width := 3,4
	dst := MakeCell(height, width)
	src :=Cell{{1,2,3,4},{5,6,7,8},{2,4,6,8}} //height:3, width:4
	DeepCopy(&dst, &src)
	assert.Equal(t,src, dst)
}
