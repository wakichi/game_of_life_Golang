package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestUpdateOneCellBirthRule(t *testing.T){
	actualNextCell :=0
	nowCell:=0
	expectedNextCell :=1
	liveNeighbor:=3
	updateOneCell(&actualNextCell, nowCell, liveNeighbor)
	assert.Equal(t, actualNextCell, expectedNextCell)
}