package main

import "math"

// I did write these functions, I used the example as a reference, but I did write these, and I understand the math
// At least I'm pretty sure I do
// but there isn't really another way to do this math properly so there isn't much leeway in variety here.
func moveForward(speed float64) {
	// if the position the player is trying to move to is empty in the world
	// rounded to fit into the indexing system
	if world[int(playerPos.X+playerDir.X*speed)][int(playerPos.Y)] == 0 {
		// move that way in the X factor of the direction vector
		playerPos.X += playerDir.X * speed
	}

	// if the position the player is trying to move to is empty in the world
	if world[int(playerPos.X)][int(playerPos.Y+playerDir.Y*speed)] == 0 {
		// move that way in the Y factor of the direction vector
		playerPos.Y += playerDir.Y * speed
	}
}

func moveBack(speed float64) {
	// if the position the player is trying to move to is empty in the world
	if world[int(playerPos.X-playerDir.X*speed)][int(playerPos.Y)] == 0 {
		// move that way in the X factor of the direction vector
		playerPos.X -= playerDir.X * speed
	}

	// if the position the player is trying to move to is empty in the world
	if world[int(playerPos.X)][int(playerPos.Y-playerDir.Y*speed)] == 0 {
		// move that way in the Y factor of the direction vector
		playerPos.Y -= playerDir.Y * speed
	}
}

func turnRight(speed float64) {
	// store the old player direction value in X
	oldDirX := playerDir.X

	// this gets back to the unit vector representation of the player direction
	// negative values go to the right on the unit circle; hence negative speed instead of speed
	// replace the x direction with the one step transformation to the right along the circumference of the unit circle
	playerDir.X = playerDir.X*math.Cos(-speed) - playerDir.Y*math.Sin(-speed)
	// replace the x direction with the one step transformation to the right along the circumference of the unit circle
	playerDir.Y = playerDir.Y*math.Cos(-speed) + oldDirX*math.Sin(-speed)

	// do the same thing to the plane (rendering of the whole map) to scale tiles properly in relation to distance
	oldPlaneX := plane.X

	plane.X = plane.X*math.Cos(-speed) - plane.Y*math.Sin(-speed)
	plane.Y = plane.Y*math.Cos(-speed) + oldPlaneX*math.Sin(-speed)
}

func turnLeft(speed float64) {
	// store old player X direction value
	oldDirX := playerDir.X

	// this gets back to the unit vector representation of the player direction
	// positive values go to the left on the unit circle; hence just speed
	// replace the x direction with the one step transformation to the left along the circumference of the unit circle
	playerDir.X = playerDir.X*math.Cos(speed) - playerDir.Y*math.Sin(speed)
	// replace the x direction with the one step transformation to the left along the circumference of the unit circle
	playerDir.Y = playerDir.Y*math.Cos(speed) + oldDirX*math.Sin(speed)

	// do the same thing to the plane (rendering of the whole map) to scale tiles properly in relation to distance
	oldPlaneX := plane.X

	plane.X = plane.X*math.Cos(speed) - plane.Y*math.Sin(speed)
	plane.Y = plane.Y*math.Cos(speed) + oldPlaneX*math.Sin(speed)
}
