package main

import (
	"fmt"
	"math/rand"
)
func printList(list [][]int){
	for i:=0; i<len(list); i++{
		fmt.Println(list[i])
	}
}
func main(){
	var width, height int
	fmt.Println("enter width as int")
	fmt.Scan(&width)
	fmt.Println(width)
	fmt.Println("enter height as int")
	fmt.Scan(&height)
	fmt.Println(height)
	var life = make([][]int, height)
	for i := 0; i < height; i++{
		life[i] = make([]int, width)
	}
	//initialize cells
	for i:=0; i<len(life); i++{
		for j := 0; j < len(life[i]); j++{
			if rand.Float64()<0.5{
				life[i][j] = 1
			}
		}
	}
	printList(life)

	// turn
	lifeNext := make([][]int, height)
	copy(lifeNext, life)
	for i:=0; i<width; i++{
		for j:=0; j<width; j++{
			liveNeighbor := countLiveNeighbor(life, i, j)
			//lifeNextのijを更新する
			//lifeNextのijのポインタ、liveNeighborを渡せば良さそう？
			updateOneCell(&life[i][j], liveNeighbor)
		}
	}
	life = lifeNext
	printList(lifeNext)
}

func countLiveNeighbor(list [][]int, x int, y int)int{
	return 0
}
func updateOneCell(cell *int, liveNeighbor int){
	
}
