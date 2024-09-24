package main

/*any live cell with fewer than two live neighbours dies
any live cell with two or three live neighbours is unaffected
any live cell with more than three live neighbours dies
any dead cell with exactly three live neighbours becomes alive*/

func calculateNextState(p golParams, world [][]byte) [][]byte {
	// Initialize the next state (result) as a new 2D slice with the same dimensions as world
	nextWorld := make([][]byte, p.imageHeight)
	for i := range nextWorld {
		nextWorld[i] = make([]byte, p.imageWidth)
	}

	for y := 0; y < p.imageHeight; y++ {
		for x := 0; x < p.imageWidth; x++ {
			// Sum up the alive neighbors
			sum := int(world[(y+p.imageHeight-1)%p.imageHeight][(x+p.imageWidth-1)%p.imageWidth])/255 +
				int(world[(y+p.imageHeight-1)%p.imageHeight][x%p.imageWidth])/255 +
				int(world[(y+p.imageHeight-1)%p.imageHeight][(x+p.imageWidth+1)%p.imageWidth])/255 +
				int(world[y%p.imageHeight][(x+p.imageWidth-1)%p.imageWidth])/255 +
				int(world[y%p.imageHeight][(x+p.imageWidth+1)%p.imageWidth])/255 +
				int(world[(y+p.imageHeight+1)%p.imageHeight][(x+p.imageWidth-1)%p.imageWidth])/255 +
				int(world[(y+p.imageHeight+1)%p.imageHeight][x%p.imageWidth])/255 +
				int(world[(y+p.imageHeight+1)%p.imageHeight][(x+p.imageWidth+1)%p.imageWidth])/255

			if world[y][x] == 255 {
				if sum < 2 || sum > 3 {
					nextWorld[y][x] = 0 // Cell dies
				} else {
					nextWorld[y][x] = 255 // Cell stays alive
				}
			} else {
				if sum == 3 {
					nextWorld[y][x] = 255 // Dead cell becomes alive
				} else {
					nextWorld[y][x] = 0 // Dead cell stays dead
				}
			}
		}
	}
	return nextWorld
}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	var liveCells []cell
	for y := 0; y < p.imageHeight; y++ {
		for x := 0; x < p.imageWidth; x++ {
			if world[y][x] == 255 {
				liveCells = append(liveCells, cell{x: x, y: y})
			}
		}
	}
	return liveCells
}
