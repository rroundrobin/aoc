package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
	_ "sort"
)

func main() {
	readFile, err := os.Open("input4.txt")
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
  
	fmt.Printf("Part 1: %v\n", SolveDay4Part1(fileLines))
	fmt.Printf("Part 2: %v\n", SolveDay4Part2(fileLines))
}

//SolveDay4Part1
func SolveDay4Part1(input []string) int {
	var result  int
	for _, line := range input{
		pair := strings.Split(line, ",")
		pair1 := strings.Split(pair[0], "-")
		pair2 := strings.Split(pair[1], "-")
		low1, _ := strconv.Atoi(pair1[0])
		low2, _ := strconv.Atoi(pair2[0])
		high1, _ := strconv.Atoi(pair1[1])
		high2, _ := strconv.Atoi(pair2[1])
		if ((low1<=low2&&high1>=high2)||(low2<=low1&&high2>=high1)){
			result++
		}
	}
	return result
}

//SolveDay4Part2
func SolveDay4Part2(input []string) int {
	var result  int
	for _, line := range input{
		pair := strings.Split(line, ",")
		pair1 := strings.Split(pair[0], "-")
		pair2 := strings.Split(pair[1], "-")
		low1, _ := strconv.Atoi(pair1[0])
		low2, _ := strconv.Atoi(pair2[0])
		high1, _ := strconv.Atoi(pair1[1])
		high2, _ := strconv.Atoi(pair2[1])
		if ((low2>=low1&&low2<=high1)||(high2<=high1&&high2>=low1)||(low2<=low1&&high2>=high1)){
			result++
		}
	}
	return result
}

