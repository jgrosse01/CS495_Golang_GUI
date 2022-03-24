package main

import (
	"flag"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// I did use the raycasting demo as a template, but I typed my own code.
// I broke down what the raycasting demo was doing and figured out what the code meant then typed it out myself.

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
	plane = pixel.V(0.0, 0.66)
}

func main() {
	// WOW FLAGS ARE COOL!!!!! I don't need to use os.Args[i]! FLAGS!
	// I copy pasted the flags but there's no other way to make flags :)))))
	flag.BoolVar(&fullscreen, "f", fullscreen, "fullscreen")
	flag.IntVar(&width, "w", width, "width")
	flag.IntVar(&height, "h", height, "height")
	flag.Float64Var(&scale, "s", scale, "scale")
	flag.Parse()

	gameSetup()

	pixelgl.Run(run)
}
