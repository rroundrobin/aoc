package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	_ "sort"
)

func main() {
	readFile, err := os.Open("input5.txt")
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
  
	fmt.Printf("Part 1: %v\n", SolveDay5Part1(fileLines))
	fmt.Printf("Part 2: %v\n", SolveDay5Part2(fileLines))
}

//SolveDay5Part1
func SolveDay5Part1(input []string) string {
	var result  string
	var highestPile int
	
	for i, line := range input{
		if len(line) == 0 {
			highestPile = i - 1
			break
		}
	}
	numOfStacks := (len(input[highestPile])+1)/4
	stacks := make([]string, numOfStacks)
	for i:=highestPile-1; i>=0;i--{
		
		for j:=0;j<numOfStacks;j++{
			positionToRead:= (j+1)*4-3
			if positionToRead>len(input[i]){break}
			charToAdd := string(input[i][positionToRead])
			if charToAdd != " "{
				stacks[j] += charToAdd
			}
			
			}
		}


	for i:=highestPile+2;i<len(input);i++{
		instructions := strings.Fields(input[i])
		numItems, _ := strconv.Atoi(instructions[1])
		src, _ := strconv.Atoi(instructions[3])
		dest, _ := strconv.Atoi(instructions[5])
		src--
		dest--
		for j:=0;j<numItems;j++{
			stacks[dest] += string(stacks[src][len(stacks[src])-1])
			stacks[src] = stacks[src][:len(stacks[src])-1]
		}
	}
	for _, line  := range stacks{
		result += string(line[len(line)-1])
	}

	return result
}

//SolveDay5Part2
func SolveDay5Part2(input []string) string {
	var result  string
	var highestPile int
	
	for i, line := range input{
		if len(line) == 0 {
			highestPile = i - 1
			break
		}
	}
	numOfStacks := (len(input[highestPile])+1)/4
	stacks := make([]string, numOfStacks)
	for i:=highestPile-1; i>=0;i--{
		
		for j:=0;j<numOfStacks;j++{
			positionToRead:= (j+1)*4-3
			if positionToRead>len(input[i]){break}
			charToAdd := string(input[i][positionToRead])
			if charToAdd != " "{
				stacks[j] += charToAdd
			}
			
			}
		}


	for i:=highestPile+2;i<len(input);i++{
 		instructions := strings.Fields(input[i])
		numItems, _ := strconv.Atoi(instructions[1])
		src, _ := strconv.Atoi(instructions[3])
		dest, _ := strconv.Atoi(instructions[5])
		src--
		dest--
		stacks[dest] += string(stacks[src][len(stacks[src])-numItems:len(stacks[src])])
		stacks[src] = stacks[src][:len(stacks[src])-numItems]
	}
	for _, line  := range stacks{
		result += string(line[len(line)-1])
	}
	return result
		
}

