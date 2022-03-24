package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
	"os"
)

func run() {
	// set the bounds to start at 0,0, make the window width by height times a scale factor
	// make sure there is a taskbar
	// lock the frame rate to the maximum refresh rate allowed on the monitor (smoother)
	configuration := pixelgl.WindowConfig{
		Title:       "Potentially A Future Doom Remake v0.1 alpha",
		Bounds:      pixel.R(0, 0, float64(width), float64(height)),
		Undecorated: false,
		VSync:       true,
	}

	// if its fullscreen just slap it on the OS designated primary monitor.
	if fullscreen {
		configuration.Monitor = pixelgl.PrimaryMonitor()
	}

	// create a window with the prior configuration values, error if that fails
	window, err := pixelgl.NewWindow(configuration)
	if err != nil {
		fmt.Println("Error creating game window.")
	}

	// storing the center of the window
	screenCenter := window.Bounds().Center()

	for !window.Closed() {
		// if you press escape, exit the program (intended)
		if window.JustPressed(pixelgl.KeyEscape) {
			os.Exit(0)
		}

		window.Clear(color.Black)

		// movement commands
		// using just pressed for movement to ensure constant inputs work properly.
		// shift conditions are added to forward and turning to increase speed of each when "sprinting"
		// I did that to turning for smoothness so that you don't get stuck running into walls
		if window.Pressed(pixelgl.KeyW) {
			if window.Pressed(pixelgl.KeyLeftShift) {
				moveForward(.44)
			} else {
				moveForward(.24)
			}
		}

		if window.Pressed(pixelgl.KeyS) {
			moveBack(.24)
		}

		if window.Pressed(pixelgl.KeyA) {
			if window.Pressed(pixelgl.KeyLeftShift) {
				turnLeft(.13)
			} else {
				turnLeft(.07)
			}
		}

		if window.Pressed(pixelgl.KeyD) {
			if window.Pressed(pixelgl.KeyLeftShift) {
				turnRight(.13)
			} else {
				turnRight(.07)
			}
		}

		// gets the image on the frame
		p := pixel.PictureDataFromImage(frame())

		// draws the newly rendered screen sprite and puts it on the window scaled around the center
		// this line could have been done other ways, and I did not write it myself since I couldn't get it working
		// in other ways
		// I did play with the scale multiplier here. The lower the value the wider your field of view. Kinda cool!
		pixel.NewSprite(p, p.Bounds()).
			Draw(window, pixel.IM.Moved(screenCenter).Scaled(screenCenter, scale*0.35))

		// update the display
		window.Update()
	}
}
