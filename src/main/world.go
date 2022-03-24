package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// NYI make a random map get read from a randomly chosen file into the [][]int format to have a different world
// each time the game is launched
// this is currently just the demo world
var world = readWorld("data/world0")

func readWorld(filepath string) [][]int {
	// variable to store the world that we will return
	var retWorld [][]int

	// open file
	file, err := os.Open(filepath)

	if err != nil {
		fmt.Println("Error opening world file for reading.")
	}

	// closes the file that was opened when the method is done
	defer func(file *os.File) {
		// get the error value from the closure of the file if there is one
		err := file.Close()
		// if there is then print that
		if err != nil {
			fmt.Println("[ERROR] Failed to close file! This is unexpected. " +
				"Memory will be freed upon exiting program")
		}
	}(file)

	// make a file reader instance for the file that is currently open
	fileScanner := bufio.NewScanner(file)

	// make a string array to read the texture bytes from a file
	var stringArr = make([][]string, len(retWorld))
	for fileScanner.Scan() {
		// append to string array the slice of the values on the line in the file separated by commas with all spaces removed
		stringArr = append(stringArr, strings.Split(strings.ReplaceAll(fileScanner.Text(), " ", ""), ","))
	}

	// for every row of the stringArr
	for i, _ := range stringArr {
		// make a new integer list that is the same length of the list at location stringArr[i]
		var intList = make([]int, len(stringArr[i]))
		// for every item in that sublist
		for j, _ := range stringArr[i] {
			// convert the value to an integer from a string
			intVar, err := strconv.Atoi(stringArr[i][j])
			// error check
			if err != nil {
				fmt.Println("Error converting world [][]string to [][]int during the loading process.")
			}
			// add that integer value to the intList list at the proper location
			intList[j] = intVar
		}
		// add the intList to the world return double indexed list
		retWorld = append(retWorld, intList)
	}

	// return the world as a [][]int
	return retWorld
}
