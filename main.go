package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	const imageWidth = 200
	const imageHeight = 100

	var image strings.Builder
	image.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", imageWidth, imageHeight))

	for j := imageHeight - 1; j >= 0; j -= 1 {
		for i := 0; i < imageWidth; i += 1 {
			r := float64(i) / (imageWidth)
			g := float64(j) / (imageHeight)
			b := 0.2

			ir := int(255.99 * r)
			ig := int(255.99 * g)
			ib := int(255.99 * b)

			image.WriteString(fmt.Sprintf("%d %d %d\n", ir, ig, ib))
		}
	}

	file, err := os.Create("out.ppm")
	if err != nil {
		panic("Could not open the file")
	}
	defer file.Close()

	_, err = file.WriteString(image.String())
	if err != nil {
		panic("Could not write in file")
	}
}
