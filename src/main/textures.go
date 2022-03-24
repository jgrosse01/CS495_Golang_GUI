package main

import (
	"bufio"
	"bytes"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
	"strconv"
	"strings"
)

// used textures from demo, I don't know how to make pretty graphics, and I also don't know how to convert an image to
// a string of bytes to use in this program since the demo taught me how to render with a byte_array rather than indexing
// an image to figure out how to select pieces of a PNG to map to one texture
var textures = textureLoader()

// define each texture as a 64x64 cube
const textureSize = 64

// loads the texture map from the byte array to be used in rendering the game
func textureLoader() *image.RGBA {
	// textures used from the pixel raycasting demo because I'm not graphically inclined
	// read the byte array from the file
	textureBytes := readByteArray("data/texture_byteslice")
	//if err != nil {
	//	fmt.Println("Error reading file containing byte array.")
	//}

	// convert the byte array containing a texture map to a png to use and render
	tempImage, err := png.Decode(bytes.NewReader(textureBytes))
	// if it errors, print an error in the console for debug
	if err != nil {
		fmt.Println("Error reading byte in png byte array.")
	}

	// create the texture map of the bytemap generated png by its bounds
	textureMap := image.NewRGBA(tempImage.Bounds())

	// overwrite textureMap as a rectangle containing the temporary image so that we can reference this as a map
	// starts draw at 0,0 (empty point)
	draw.Draw(textureMap, textureMap.Bounds(), tempImage, image.Point{}, draw.Src)

	// returns texture map as a PNG image
	return textureMap
}

// this is the only way to type this that makes sense
// same as the demo, but not copying an algorithm
func getTexNum(x, y int) int {
	return world[x][y]
}

func readByteArray(filepath string) []byte {
	// list to return
	var retList []byte
	// open file and store any potential errors
	file, err := os.Open(filepath)
	// print potential error
	if err != nil {
		fmt.Println("Error in opening texture_byteslice file.")
	}
	// when the method exits, do this.
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
	var stringArr []string
	for fileScanner.Scan() {
		stringArr = strings.Split(strings.ReplaceAll(fileScanner.Text(), " ", ""), ",")
	}

	// convert to integers (since strings cannot be converted to a byte properly)
	// make an int slice of the size of stringArr
	var intArray = make([]int, len(stringArr))
	for i, _ := range stringArr {
		// convert each string in the list stringArr to an integer value then put it at that index of intArray
		intVar, err := strconv.Atoi(stringArr[i])
		if err != nil {
			fmt.Println("Error converting string element to int in the process of reading texture bytes from file.")
		}
		intArray[i] = intVar
	}

	// finally convert all of the integers to bytes to be used to make a PNG texture.
	for _, element := range intArray {
		retList = append(retList, byte(element))
	}

	// return the texture byte slice
	return retList
}
