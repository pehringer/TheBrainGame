package grid

import (
	"encoding/json"
	"fmt"
	"os"
)

func Json(matrices []Matrix) {
	output := make([]Coordinates, len(matrices))
	for m := 0; m < len(matrices); m++ {
		for z := 0; z < matrices[m].Layers; z++ {
			for y := 0; y < matrices[m].Rows; y++ {
				for x := 0; x < matrices[m].Columns; x++ {
					if *matrices[m].Index(z, y, x) == 0 {
						continue
					}
					output[m] = append(output[m], Coordinate{
						X: x,
						Y: y,
						Z: z,
					})
				}
			}
		}
	}

	if b, err := json.MarshalIndent(output, "\n", "\t"); err != nil {
		fmt.Println(err)
		return
	} else if f, err := os.Create("output.json"); err != nil {
		fmt.Println(err)
		return
	} else {
		f.Write(b)
		f.Close()
	}
}
