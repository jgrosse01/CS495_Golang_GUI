package main

import (
	"flag"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// I did use the raycasting demo as a template, but I typed my own code for all but windowFrame.go which I copy pasted
// and slightly modified to fit my naming convention.

var fullscreen = false

// game window width!
var width = 1920

// game window height!
var height = 1080

// GUI element scale!
var scale = 3.0

// stores player position as a vector from coordinate (0,0)
var playerPos pixel.Vec

// stores player direction as a unit vector
var playerDir pixel.Vec

// initializes the flat surface that the "camera" will move and rotate on
var plane pixel.Vec

func gameSetup() {
	// initial location is 12 to the right and 14.5 down in the map
	playerPos = pixel.V(12.0, 14.5)
	// initial direction is West
	playerDir = pixel.V(-1.0, 0.0)
	// camera height is 0.7 out of the total tile height of 1.
	plane = pixel.V(0.0, 0.7)
}

func main() {
	// WOW FLAGS ARE COOL!!!!! I don't need to use os.Args[i]! FLAGS!
	// I referenced the demo here for the flags but there's no other way to make flags that I can find :)))))
	flag.BoolVar(&fullscreen, "f", fullscreen, "fullscreen")
	flag.IntVar(&width, "w", width, "width")
	flag.IntVar(&height, "h", height, "height")
	flag.Float64Var(&scale, "s", scale, "scale")
	flag.Parse()

	// call the game setup where we place our player
	gameSetup()

	// run the gui window and its associated main loop
	pixelgl.Run(run)
}
