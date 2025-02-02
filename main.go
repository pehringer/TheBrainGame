package main

import "github.com/pehringer/TheBrainGame/internal/grid"

func main() {
	matrices := []grid.Matrix{
		grid.New(2, 2, 2,
			1, 0,
			0, 0,
			0, 0,
			0, 0,
		),
		grid.New(2, 2, 2,
			1, 0,
			1, 0,
			0, 0,
			0, 0,
		),
		grid.New(2, 2, 2,
			1, 1,
			1, 0,
			0, 0,
			0, 0,
		),
		grid.New(2, 2, 2,
			1, 1,
			1, 0,
			1, 0,
			0, 0,
		),
	}
	grid.Json(matrices)
}
