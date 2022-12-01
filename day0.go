package main

import (
	"fmt"
	"os"
	"bufio"
	_ "strconv"
	_ "sort"
)

func main() {
	readFile, err := os.Open("input0.txt")
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
  
	fmt.Printf("Part 1: %v\n", SolveDay0Part1(fileLines))
	fmt.Printf("Part 2: %v\n", SolveDay0Part2(fileLines))
}

//SolveDay0Part1
func SolveDay0Part1(input []string) int {
	var result  int
	return result
}

//SolveDay0Part2
func SolveDay0Part2(input []string) int {
	var result  int
	return result
}

