package main

import (
	"fmt"
	"os"
	"bufio"
	_ "strconv"
	_ "sort"
)

func main() {
	readFile, err := os.Open("input2.txt")
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
  
	fmt.Printf("Part 1: %v\n", SolveDay2Part1(fileLines))
	fmt.Printf("Part 2: %v\n", SolveDay2Part2(fileLines))
}

//A rock, paper, scissors

//SolveDay2Part1
func SolveDay2Part1(input []string) int {
	var result  int
	for _, line := range input{
		var elfsBet, myBet byte
		elfsBet = line[0]
		myBet = line[2]
		switch myBet {
		case 'X':{
			result += 1
			switch elfsBet{
			case 'A': result += 3
			case 'B': result += 0
			case 'C': result += 6
			}
		}
		case 'Y':{
			result += 2
			switch elfsBet{
				case 'A': result += 6
				case 'B': result += 3
				case 'C': result += 0
				}
		}
		case 'Z':{
			result += 3
			switch elfsBet{
				case 'A': result += 0
				case 'B': result += 6
				case 'C': result += 3
				}
		}
			
		}
	}
	return result
}

//SolveDay2Part2
func SolveDay2Part2(input []string) int {
	var result  int
	for _, line := range input{
		var elfsBet, myBet byte
		elfsBet = line[0]
		myBet = line[2]
		switch myBet {
		case 'X':{
			result += 0
			switch elfsBet{
			case 'A': result += 3
			case 'B': result += 1
			case 'C': result += 2
			}
		}
		case 'Y':{
			result += 3
			switch elfsBet{
				case 'A': result += 1
				case 'B': result += 2
				case 'C': result += 3
				}
		}
		case 'Z':{
			result += 6
			switch elfsBet{
				case 'A': result += 2
				case 'B': result += 3
				case 'C': result += 1
				}
		}
			
		}
	}

	return result
}

