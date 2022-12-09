package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	_ "sort"
)

type file struct{
	name string
	size int
}

type directory struct{
	name string
	size int
	parentDir *directory
	subdirectories []*directory
	files []file
	isRootDir bool
}

func main() {
	readFile, err := os.Open("input7.txt")
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
	
	rootDir := parseInput(fileLines)
	fmt.Println(rootDir)
	usedDiskSpace := (calcDirSize(&rootDir))
	unusedSpace  := 70000000 - usedDiskSpace
	spaceToFree := 30000000 - unusedSpace
	fmt.Println(rootDir)
	minSizeDir := getSmallestToDelete(&rootDir, &rootDir, spaceToFree)
	fmt.Println("Deleting ", minSizeDir.name, "with size ", minSizeDir.size)
	
	fmt.Printf("Part 1: %v\n", SolveDay7Part1(fileLines))
	fmt.Printf("Part 2: %v\n", SolveDay7Part2(fileLines))
}

func getSmallestToDelete(workingDir *directory, currentDir *directory,  minSize int)*directory{
	if workingDir.size >= minSize && workingDir.size < currentDir.size{
		currentDir = workingDir
	}
	for _, subDir := range workingDir.subdirectories{
		 currentDir = getSmallestToDelete(subDir, currentDir, minSize)
	}
	return currentDir		
}

func parseInput(input []string) directory{
	var currentDir *directory
	var rootDir directory
	rootDir.name = "/"
	rootDir.isRootDir = true
	currentDir = &rootDir
	for _, line := range input{
		fmt.Println(rootDir)
		fmt.Println("Processing line ", line)
		fmt.Println("Current Directory", currentDir.name)
		if line[0] == '$'{
			command := line[2:4]
			if command == "cd"{
					if line[5] == '/' {
						fmt.Println("Switching to root directory")
						currentDir = &rootDir
					}else if line[5] == '.'{
						currentDir = currentDir.parentDir
					}else{
						fmt.Println("Creating directory")
						tmpDir := new(directory)
						tmpDir.parentDir = currentDir
						tmpDir.name = line[5:]
						tmpDir.isRootDir = false
						currentDir.subdirectories = append(currentDir.subdirectories, tmpDir)
						currentDir = tmpDir
					}
				}
		}else{
			if line[:3] != "dir"{
				fmt.Println("processing listing entry")
				listEntry := strings.Fields(line)
				tmpFile := new(file)
				tmpFile.name = listEntry[1]
				tmpFile.size, _ = strconv.Atoi(listEntry[0])
				currentDir.files = append(currentDir.files, *tmpFile)
				currentDir.size += tmpFile.size
			}else{
				fmt.Println("Skipping directory listing")
			}
		}
	}
	return rootDir
}

func calcDirSize(workingDir *directory) int{
var size int
for _, directory := range workingDir.subdirectories{
	size += calcDirSize(directory)
}
for _, file := range workingDir.files{
	size += file.size
}
workingDir.size=size
return size
}

func calcSum(workingDir *directory) int{
	var size int
	if workingDir.size <=100000{
		size += workingDir.size
	}
	for _, subDir := range workingDir.subdirectories{
		size += calcSum(subDir)
	}
	return size		
}

//SolveDay7Part1
func SolveDay7Part1(input []string) int {
	var result  int
	return result
}

//SolveDay7Part2
func SolveDay7Part2(input []string) int {
	var result  int
	return result
}