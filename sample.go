package main

import (
	"fmt"
	"math/rand"
)

type Cell [][]int


func main(){
	var initialCell Cell
	var numRound  int
	numRound, initialCell = initialize()
	runGame(numRound, initialCell)
}

func initialize()(int, Cell){
	var initialCell Cell
	var numRound  int
	var isRandomCell string
	fmt.Println("Would you like to use a pre-prepared cell? yes:y, no:any other key")
	fmt.Scan(&isRandomCell)
	fmt.Println("how many round do you want to play? enter any non-negative integer")
	fmt.Scan(&numRound)
	if isRandomCell == "y"{
		initialCell = makeRandomCell()
	}else{
		initialCell = readPrebuiltCell()
	}
	return numRound, initialCell
}

func runGame(numRound int, initialCell Cell){
	lifeNow := MakeCell(len(initialCell), len(initialCell[0]))
	DeepCopy(&lifeNow, &initialCell)
	for i := 0; i<numRound;i++{
		var lifeNext Cell
		lifeNext = runRound(lifeNow)
		PrintResult(lifeNext, i)
		DeepCopy(&lifeNow, &lifeNext)
	}
}
                                                    
func makeRandomCell()Cell{
	var width, height int
	fmt.Println("enter width as int")
	fmt.Scan(&width)
	fmt.Println(width)
	fmt.Println("enter height as int")
	fmt.Scan(&height)
	fmt.Println(height)
	var life = MakeCell(height, width)
	//initialize cells
	for i:=0; i<len(life); i++{
		for j := 0; j < len(life[i]); j++{
			if rand.Float64()<0.5{
				life[i][j] = 1
			}
		}
	}
	return life
}

func readPrebuiltCell()Cell{
	//todo: not implemented
	panic("not implemented")
	return Cell{}
}

func runRound(lifeNow Cell)Cell{
	height:=len(lifeNow)
	width:=len(lifeNow[0])
	lifeNext := MakeCell(len(lifeNow), len(lifeNow[0]))
	for y:=0; y<height; y++{
		for x:=0; x<width; x++{
			liveNeighbor := countLiveNeighbor(lifeNow, x, y, width, height)
			updateOneCell(&lifeNext[y][x],lifeNow[y][x], liveNeighbor)
		}
	}
	return lifeNext
}

func countLiveNeighbor(list Cell, x int, y int, width int, height int) int {
	count :=0

	//遷移のパターンを全列挙、枠外は除外
	for i := -1; i<2; i++{
		for j := -1; j<2; j++{
			if i ==0 && j ==0 {
				continue
			}
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

func convertLifeCell(life Cell)[]string{
	res:=make([]string, len(life))
	converter:=map[int]string{
		0:"□ ",
		1:"■ ",
	}
	for i:=0;i<len(life);i++{
		str := ""
		for j:=0;j<len(life[i]);j++{
			str+=converter[life[i][j]]
		}
		res[i] = str
	}
	return res
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
