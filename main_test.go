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

func TestCountLiveNeighbor(t *testing.T){
	cell := Cell{
		{0,1,0,1},
		{1,1,1,1},
		{1,1,0,0},
		{0,0,1,1},
	}
	// case 1
	expected1:= 6
	actual1 := countLiveNeighbor(cell, 2,2,4,4)
	assert.Equal(t, actual1, expected1)

	//case 2
	expected2 := 2
	actual2 := countLiveNeighbor(cell, 3,0,4,4)
	assert.Equal(t,expected2, actual2)

	//case3
	expected3:=2
	actual3 := countLiveNeighbor(cell, 2,3,4,4)
	assert.Equal(t, expected3, actual3)
}

func TestRunRound(t *testing.T){
	previousCell :=Cell{
		{0,0,1},
		{1,1,1},
		{1,0,0},
		{1,0,1},
	}

	actualLaterCell := runRound(previousCell)
	expectedLaterCell:=Cell{
		{0,0,1},
		{1,0,1},
		{1,0,1},
		{0,1,0},
	}

	assert.Equal(t, expectedLaterCell, actualLaterCell)
}