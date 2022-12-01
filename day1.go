package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"sort"
)

func main() {
	readFile, err := os.Open("input1.txt")
	if err != nil {
		os.Exit(1)
	}
	
	fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
    var fileLines []string
  
    for fileScanner.Scan() {
        fileLines = append(fileLines, fileScanner.Text())
    }
  
    readFile.Close()
  
	fmt.Printf("Part 1: %v\n", SolveDay1Part1(fileLines))
	fmt.Printf("Part 2: %v\n", SolveDay1Part2(fileLines))
}

//SolveDay1Part1
func SolveDay1Part1(input []string) int {
	var highval, curval int
	for _, line := range input {
		if line != ""{
			tmpval, _ := strconv.Atoi(line)
			curval += tmpval
			if curval > highval {highval = curval}
		}else{
			curval = 0
		}
	}
	return highval
}

//SolveDay1Part2
func SolveDay1Part2(input []string) int {
	var curval, numElves int
	var elfCalories [] int
	for _, line := range input {
		if line != ""{
			tmpval, _ := strconv.Atoi(line)
			curval += tmpval
		}else{
			elfCalories = append(elfCalories, curval)
			numElves++
			curval = 0
		}
	}
	sort.Ints(elfCalories)
	fmt.Println(elfCalories[numElves-1])
	fmt.Println(elfCalories[numElves-2])
	fmt.Println(elfCalories[numElves-3])
	
	return (elfCalories[numElves-1] + elfCalories[numElves-2] + elfCalories[numElves-3])
}

