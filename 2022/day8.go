package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	_ "sort"
)

func main() {
	readFile, err := os.Open("input8.txt")
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
	var visibilityCount  int
	heightMap := make([][]int, len(input))
	for i, line := range(input){
		heightMap[i] = make([]int, len(line))
		for j, _ := range(line){
			heightMap[i][j], _ = strconv.Atoi(string(input[i][j]))
		}
	}

	for i, line := range(heightMap){
		for j, height := range(line){
			var visa, visb, visc, visd bool
			visa=true
			visb=true
			visc=true
			visd=true
//			fmt.Println("Checking i, j", i, j, " value", height)
			for z:=0;z<j;z++{
				if line[z]>=height{
					visa=false
//					fmt.Println("1checking against", i, z, visa)
					break
				}
			}
			for z:=j+1;z<len(line);z++{
				if line[z]>=height{
					visb=false
//					fmt.Println("2checking against", i, z, visb)
					break
				}
			}
			for z:=0;z<i;z++{
				if heightMap[z][j]>=height{
					visc=false
//					fmt.Println("3checking against", z, j, visc)
					break
				}
			}
			for z:=i+1;z<len(heightMap);z++{
				if heightMap[z][j]>=height{
					visd=false
//					fmt.Println("4checking against", z, j, visd)
					break
				}
			}
			visible := visa||visb||visc||visd
//			fmt.Println(visible)
			if visible{
				visibilityCount++
			}
		}

	}
	return visibilityCount
}

//SolveDay0Part2
func SolveDay0Part2(input []string) int {
	var scenicscore int
	heightMap := make([][]int, len(input))
	for i, line := range(input){
		heightMap[i] = make([]int, len(line))
		for j, _ := range(line){
			heightMap[i][j], _ = strconv.Atoi(string(input[i][j]))
		}
	}

	for i, line := range(heightMap){
		for j, height := range(line){
			var visa, visb, visc, visd int
			for z:=j-1;z>=0;z--{
				visa++
				if line[z]>=height{
					break
				}
			}
			for z:=j+1;z<len(line);z++{
				visb++
				if line[z]>=height{
					break
				}
			}
			for z:=i-1;z>=0;z--{
				visc++
				if heightMap[z][j]>=height{					
					break
				}
			}
			for z:=i+1;z<len(heightMap);z++{
				visd++
				if heightMap[z][j]>=height{
					break
				}
			}
			tmpscenicscore := visa*visb*visc*visd			
			if tmpscenicscore > scenicscore{
			scenicscore = tmpscenicscore
			}
		}
		
	}
	return scenicscore
}