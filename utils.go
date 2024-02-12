package main

import (
	"fmt"
	"time"
)


func PrintList(list Cell){
	for i:=0; i<len(list); i++{
		fmt.Println(list[i])
	}
}

func MakeCell(height int, width int)Cell{
	var cell Cell
	cell = make(Cell, height)
	for i := 0; i < height; i++ {
		cell[i] = make([]int, width)
	}
	return cell
}
func DeepCopy(dst *Cell, src *Cell){
	//Didnt work, src cannnot copy into dst
	for i := 0; i < len(*dst); i++{
		copy((*dst)[i],(*src)[i])
	}
}

func PrintResult(life Cell, ith int){
	// printList(lifeCount)
	fmt.Println(ith+1,"th generation")
	convertedLife :=convertLifeCell(life)
	for i:=0; i<len(life);i++{
		println(convertedLife[i])
	}
	fmt.Println("")
	time.Sleep(time.Second*1)
	fmt.Print("\033[H\033[2J")
}