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
func deepCopy(dst [][]int, src [][]int){
	for i := 0; i < len(dst); i++{
		copy(dst[i],src[i])
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
	fmt.Println("first generation")
	printList(life)
	fmt.Println("")

	// turn
	for i:=0; i<10; i++{
		lifeNext := make([][]int, height)
		lifeCount:=make([][]int, height)
		for i := 0; i < height; i++{
			lifeCount[i] = make([]int, width)
			lifeNext[i] = make([]int, width)
			copy(lifeNext[i],life[i])
		}
		for y:=0; y<height; y++{
			for x:=0; x<width; x++{
				liveNeighbor := countLiveNeighbor(life, x, y, width, height)
				// for debug
				lifeCount[y][x] = liveNeighbor
				//lifeNextのijを更新する
				//lifeNextのijのポインタ、liveNeighborを渡せば良さそう？
				//TODO: 下の行が元凶。たぶんリストのコピーが参照渡しになってるかも
				updateOneCell(&lifeNext[y][x],life[y][x], liveNeighbor)
			}
		}
		// printList(lifeCount)
		life = lifeNext
		fmt.Println(i+1,"th generation")
		printList(lifeNext)
		fmt.Println("")
	}
}

func countLiveNeighbor(list [][]int, x int, y int, width int, height int) int {
	count :=0

	//遷移のパターンを全列挙、枠外は除外
	for i := -1; i<2; i++{
		for j := -1; j<2; j++{
			if i ==0 && j ==0 {
				continue
			}
			//debug nx, nyの列挙は良さそう
			nx:=x+j
			ny:=y+i
			//ngパターンを除外
			if !(nx<width && nx>=0 && ny<height && ny>=0) {
				continue
			}
			count+=list[ny][nx]
		}
	}
	return count
}

func updateOneCell(nextCell *int, nowCell int, liveNeighbor int){
	//birth rule
	if nowCell == 0 && liveNeighbor == 3{
		*nextCell = 1
	
	}else if nowCell == 1 && (liveNeighbor ==2 ||liveNeighbor ==3){
	//live rule
		*nextCell = 1
	}else{
		// death rule
		*nextCell = 0
	}
}

