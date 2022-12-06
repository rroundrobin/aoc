package main

import (
	"fmt"
	"os"
	"bufio"
	_ "strconv"
	_ "sort"
)

func main() {
	readFile, err := os.Open("input6.txt")
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
  
	fmt.Printf("Part 1: %v\n", SolveDay6Part1(fileLines))
	fmt.Printf("Part 2: %v\n", SolveDay6Part2(fileLines))
}

//SolveDay6Part1
func SolveDay6Part1(input []string) int {
	var result int
	for i:=0;i<(len(input[0])-3);i++{
		var a,b,c,d byte
		a=input[0][i]
		b=input[0][i+1]
		c=input[0][i+2]
		d=input[0][i+3]
		if (a!=b&&a!=c&&a!=d&&b!=c&&b!=d&&c!=d){
			return i+4
		}
	}
	return result
}

//SolveDay6Part2
func SolveDay6Part2(input []string) int {
	var result  int
	for i:=0;i<(len(input[0])-13);i++{
		testString := input[0][i:i+14]
		isResult := true
		for j, charToTest := range testString{
			for k, secondChar := range testString{
				if j!=k&& charToTest==secondChar{
					isResult=false
				}
			}
		}
		if isResult {
			return i+14
			}
		}
		return result
	}


