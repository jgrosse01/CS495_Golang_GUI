package main

import (
	"github.com/faiface/pixel"
	"image"
	"math"
)

// I DID NOT WRITE FRAME, JUST MODIFIED IT. THERE APPEARS TO BE MUCH COMPLEXITY IN THIS MATH.
// I did however do my best to comment it and some variable names are changed.
// I am trying to make sense of this since the demo code was all condensed and came with no comments.
func frame() *image.RGBA {
	// here we make a new empty rectangle of the proper window size into an empty image
	m := image.NewRGBA(image.Rect(0, 0, width, height))

	// for the entire width of the empty image
	for x := 0; x < width; x++ {
		var (
			// not sure about step
			step image.Point
			// distance from the nearest visible side of each tile?
			sideDist pixel.Vec
			// perpetual distance from the wall?
			perpWallDist float64
			// if the player is hitting a wall. and if the hitting is happening to the relevant side of the tile?
			hit, side bool

			// position of the object to cast rays and the location in the world
			rayPos, worldX, worldY = playerPos, int(playerPos.X), int(playerPos.Y)

			// sets the screen focus point for the camera so it stays where it is in relation to the monitor
			cameraX = 2*float64(x)/float64(width) - 1

			// the direction that rays should be cast (unit vector facing the direction of the rays being cast)
			rayDir = pixel.V(
				playerDir.X+plane.X*cameraX,
				playerDir.Y+plane.Y*cameraX,
			)

			// the change in the distance that the rays are to be cast in?
			deltaDist = pixel.V(
				math.Sqrt(1.0+(rayDir.Y*rayDir.Y)/(rayDir.X*rayDir.X)),
				math.Sqrt(1.0+(rayDir.X*rayDir.X)/(rayDir.Y*rayDir.Y)),
			)
		)

		// no clue what these if statements are doing math wise...
		if rayDir.X < 0 {
			step.X = -1
			sideDist.X = (rayPos.X - float64(worldX)) * deltaDist.X
		} else {
			step.X = 1
			sideDist.X = (float64(worldX) + 1.0 - rayPos.X) * deltaDist.X
		}

		if rayDir.Y < 0 {
			step.Y = -1
			sideDist.Y = (rayPos.Y - float64(worldY)) * deltaDist.Y
		} else {
			step.Y = 1
			sideDist.Y = (float64(worldY) + 1.0 - rayPos.Y) * deltaDist.Y
		}

		for !hit {
			if sideDist.X < sideDist.Y {
				sideDist.X += deltaDist.X
				worldX += step.X
				side = false
			} else {
				sideDist.Y += deltaDist.Y
				worldY += step.Y
				side = true
			}

			if world[worldX][worldY] > 0 {
				hit = true
			}
		}

		var wallX float64

		if side {
			perpWallDist = (float64(worldY) - rayPos.Y + (1-float64(step.Y))/2) / rayDir.Y
			wallX = rayPos.X + perpWallDist*rayDir.X
		} else {
			perpWallDist = (float64(worldX) - rayPos.X + (1-float64(step.X))/2) / rayDir.X
			wallX = rayPos.Y + perpWallDist*rayDir.Y
		}

		// get width of a wall tile?
		wallX -= math.Floor(wallX)

		// note the texture size of the tile
		texX := int(wallX * float64(textureSize))

		// not sure
		lineHeight := int(float64(height) / perpWallDist)

		// no idea here
		if lineHeight < 1 {
			lineHeight = 1
		}

		drawStart := -lineHeight/2 + height/2
		if drawStart < 0 {
			drawStart = 0
		}

		drawEnd := lineHeight/2 + height/2
		if drawEnd >= height {
			drawEnd = height - 1
		}

		if !side && rayDir.X > 0 {
			texX = textureSize - texX - 1
		}

		if side && rayDir.Y < 0 {
			texX = textureSize - texX - 1
		}

		texNum := getTexNum(worldX, worldY)

		for y := drawStart; y < drawEnd+1; y++ {
			d := y*256 - height*128 + lineHeight*128
			texY := ((d * textureSize) / lineHeight) / 256

			c := textures.RGBAAt(
				texX+textureSize*(texNum),
				texY%textureSize,
			)

			if side {
				c.R = c.R / 2
				c.G = c.G / 2
				c.B = c.B / 2
			}

			m.Set(x, y, c)
		}

		// just a reminder that I still have no clue here
		var floorWall pixel.Vec

		if !side && rayDir.X > 0 {
			floorWall.X = float64(worldX)
			floorWall.Y = float64(worldY) + wallX
		} else if !side && rayDir.X < 0 {
			floorWall.X = float64(worldX) + 1.0
			floorWall.Y = float64(worldY) + wallX
		} else if side && rayDir.Y > 0 {
			floorWall.X = float64(worldX) + wallX
			floorWall.Y = float64(worldY)
		} else {
			floorWall.X = float64(worldX) + wallX
			floorWall.Y = float64(worldY) + 1.0
		}

		distWall, distPlayer := perpWallDist, 0.0

		// still nope
		for y := drawEnd + 1; y < height; y++ {
			currentDist := float64(height) / (2.0*float64(y) - float64(height))

			weight := (currentDist - distPlayer) / (distWall - distPlayer)

			currentFloor := pixel.V(
				weight*floorWall.X+(1.0-weight)*playerPos.X,
				weight*floorWall.Y+(1.0-weight)*playerPos.Y,
			)

			fx := int(currentFloor.X*float64(textureSize)) % textureSize
			fy := int(currentFloor.Y*float64(textureSize)) % textureSize

			m.Set(x, y, textures.At(fx, fy))

			m.Set(x, height-y-1, textures.At(fx+(4*textureSize), fy))
			m.Set(x, height-y, textures.At(fx+(4*textureSize), fy))
		}
	}

	// returns the built png
	return m
}
