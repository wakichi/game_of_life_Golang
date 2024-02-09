package main

import (
	"fmt"
	"math/rand"
	"time"
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
func main2(){
	initialCell:= [][]int
	numRound := int
	initialCell, numRound = initialize()
	runGame(initialCell, numRound)
}

func initialize()(int, [][]int){
	initialCell := [][]int
	numRound := int
	isRandomCell:=string
	fmt.Println("Would you like to use a pre-prepared cell? yes:y, no:any other key")
	fmt.Scan(&isRandomCell)
	fmt.Println("how many round do you want to play? enter any non-negative integer")
	if isRandomCell == "y"{
		initialCell = makeRandomCell()
	}else{
		initialCell = readPrebuiltCell()
	}
	return numRound, initialCell
}

func runGame(numRound int, initialCell [][]int){
	lifeNow:=[][]int
	deepCopy(lifeNow, initialCell)
	for i := 0; i<numRound;i++{
		lifeNext :=[][]int
		lifeNext = runRound(lifeNow)
		printResult(lifeNext)
		deepCopy(lifeNow, lifeNext)
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
	for i:=0; i<100; i++{
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
		convertedLife :=convertLifeCell(life)
		for i:=0; i<height;i++{
			println(convertedLife[i])
		}
		fmt.Println("")
		time.Sleep(time.Second*1)
		fmt.Print("\033[H\033[2J")
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
func convertLifeCell(life [][]int)[]string{
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
