package main

import (
	"fmt"
	"os"
	"bufio"
	"bytes"
	_ "strconv"
	_ "sort"
)

func main() {
	readFile, err := os.Open("input3.txt")
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
  
	fmt.Printf("Part 1: %v\n", SolveDay3Part1(fileLines))
	fmt.Printf("Part 2: %v\n", SolveDay3Part2(fileLines))
}

//SolveDay3Part1
func SolveDay3Part1(input []string) int {
	var result  int
	// zwei Teilstrings machen
	// suche für jedes Zeichen in 1 nach Übereinstimmung
	// byteValues für Großbuchstaben sind zwischen 65und 90, Kleinbuchstabenzwischen 97 und 122
	for _, bagpack := range input{
		firstCompartment := bagpack[:len(bagpack)/2]
		secondCompartment := bagpack[len(bagpack)/2:]
		checkBagpack:
		for _, item1 := range []byte(firstCompartment){
			for _, item2 := range []byte(secondCompartment){
				if item1 == item2 {
					if item1 >=97 {
						result += int(item1 - 96)
					}else{
						result += int(item1 - 38)
					}
					break checkBagpack
				}
			}
		}
	}
	return result
}

//SolveDay3Part2
func SolveDay3Part2(input []string) int {
	var result  int
	numberOfGroups := len(input) /3
	for i := 0; i<numberOfGroups;i++{
	//	fmt.Println("Analyzing Group" , i)
		var tmpItems string
		buffer := bytes.NewBufferString(tmpItems)
		for _, item1 := range []byte(input[i*3]){
			for _, item2 := range []byte(input[(i*3)+1]){
				if item1 == item2 {
					buffer.WriteByte(item1)
				}
			}
		}
		tmpItems = buffer.String()
	//	fmt.Println("Übereinstimmung 1/2: " + tmpItems)
		check2to3:
		for _, item1 := range []byte(tmpItems){
			for _, item2 := range []byte(input[(i*3+2)]){
				if item1 == item2 {
					if item1 >=97 {
						result += int(item1 - 96)
					}else{
						result += int(item1 - 38)
					}
					break check2to3
				}
			}
		}
	}
	return result
}

