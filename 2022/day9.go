package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	_ "sort"
)

func main() {
	readFile, err := os.Open("input9.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
    var fileLines []string
  
    for fileScanner.Scan() {
        fileLines = append(fileLines, fileScanner.Text())
    }
  
    readFile.Close()
  
	fmt.Printf("Part 1: %v\n", SolveDay9Part1(fileLines))
	fmt.Printf("Part 2: %v\n", SolveDay9Part2(fileLines))
}

type coordinate struct{
	x, y int
}

//SolveDay9Part1
func SolveDay9Part1(input []string) int {
	var result  int
	usedMap := make(map[coordinate]bool)
	headcoordinate := coordinate{x:0, y:0}
	tailcoordinate := coordinate{x:0, y:0}
	usedMap[tailcoordinate] = true

	for _, line := range(input){
		//fmt.Println(usedMap)
		//fmt.Println("processing ", line)
		direction := line[0]
		steps, _ := strconv.Atoi(string(line[2:]))
		for i:=0;i<steps;i++{
			curx := headcoordinate.x
			cury := headcoordinate.y
		switch direction{
		case 'U':{
			headcoordinate = coordinate{curx,cury-1}
		}
		case 'D':{
			headcoordinate = coordinate{curx,cury+1}
		}
		case 'L':{
			headcoordinate = coordinate{curx-1,cury}
		}
		case 'R':{
			headcoordinate = coordinate{curx+1,cury}
		}
		default:{panic("unknown case")}
		}
		//fmt.Println("new head ", headcoordinate)
		if moveTail(&headcoordinate, &tailcoordinate){
		//fmt.Println(line, "head is now ", headcoordinate)
		//fmt.Println("adding ", tailcoordinate)
		usedMap[tailcoordinate] = true
		}
		}
	}
	result = len(usedMap)
	//fmt.Println(usedMap)
	return result
}

func moveTail(head, tail *coordinate) bool {
	deltax := head.x - tail.x
	deltay := head.y - tail.y
	moved := false
	//Bedarf geradliniges Movement
/*	if deltax == 0{
		if deltay > 1{
			tail.y = tail.y +1
			moved = true
		}
		if deltay < -1{
			tail.y = tail.y -1
			moved = true
		}
	}
	if deltay == 0{
		if deltax > 1{
			tail.x = tail.x +1
			moved = true
		}
		if deltax < -1{
			tail.x = tail.x -1
			moved = true
		}
	}
	//diagonale Movements
	if (deltax == -2 && deltay == -1) || (deltax == -1 && deltay == -2) || (deltax == -2 && deltay == -2) {
			tail.x = tail.x - 1
			tail.y = tail.y - 1
			moved = true
	}
	if (deltax == 1 && deltay == -2) || (deltax == 2 && deltay == -1) || (deltax == 2 && deltay == -2){
		tail.x = tail.x + 1
		tail.y = tail.y - 1
		moved = true
	}
	if (deltax == 1 && deltay == 2) || (deltax == 2 && deltay == 1) || (deltax == 2 && deltay == 2){
		tail.x = tail.x + 1
		tail.y = tail.y + 1
		moved = true
	}
	if (deltax == -1 && deltay == 2) || (deltax == -2 && deltay == 1) || (deltax == -2 && deltay == 2){
		tail.x = tail.x - 1
		tail.y = tail.y + 1
		moved = true
	}	
	*/

	if (Abs(deltax) > 1 || Abs(deltay) > 1){
		//Tail muss sich bewegen
		moved = true
		tail.x += moveCount(deltax)
		tail.y += moveCount(deltay)
		
	}	
	//fmt.Println("delta", deltax, deltay)
	//fmt.Println("head", head.x, head.y)
	//fmt.Println("tail", tail.x, tail.y)
	return moved
}

func Abs(x int) int{
	if x<0{
		return x * -1
	}
	return x
}

func moveCount(x int)int{
	if x < 0 {
		return -1
	}
	if x == 0 {
		return 0
	}
	return 1
}


//SolveDay9Part2
func SolveDay9Part2(input []string) int {
	var result  int
	
	usedMap := make(map[coordinate]bool)
	headcoordinate := coordinate{x:0, y:0}
	knot1 := coordinate{x:0, y:0}
	knot2 := coordinate{x:0, y:0}
	knot3 := coordinate{x:0, y:0}
	knot4 := coordinate{x:0, y:0}
	knot5 := coordinate{x:0, y:0}
	knot6 := coordinate{x:0, y:0}
	knot7 := coordinate{x:0, y:0}
	knot8 := coordinate{x:0, y:0}
	tailcoordinate := coordinate{x:0, y:0}
	knots := []*coordinate{&headcoordinate, &knot1, &knot2, &knot3, &knot4, &knot5, &knot6, &knot7, &knot8, &tailcoordinate}
	usedMap[tailcoordinate] = true

	for _, line := range(input){
		//fmt.Println(usedMap)
		//fmt.Println("processing ", line)
		direction := line[0]
		steps, _ := strconv.Atoi(string(line[2:]))
		for i:=0;i<steps;i++{
			curx := headcoordinate.x
			cury := headcoordinate.y
		switch direction{
		case 'U':{
			headcoordinate = coordinate{curx,cury-1}
		}
		case 'D':{
			headcoordinate = coordinate{curx,cury+1}
		}
		case 'L':{
			headcoordinate = coordinate{curx-1,cury}
		}
		case 'R':{
			headcoordinate = coordinate{curx+1,cury}
		}
		default:{panic("unknown case")}
		}
		for i:=1;i<len(knots);i++{
			if moveTail(knots[i-1], knots[i]) && (i == (len(knots)-1)){
				usedMap[tailcoordinate] = true		
			}
		}
		}
		fmt.Println(knots[0], knots[1], knots[2], knots[3], knots[4],knots[5], knots[6], knots[7], knots[8], knots[9])
	}
	result = len(usedMap)
	//fmt.Println(usedMap)
	return result
}

