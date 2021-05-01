package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/alexislozano/go-raytracing/ray"
	"github.com/alexislozano/go-raytracing/vec3"
)

func hitSphere(center vec3.Vec3, radius float64, r *ray.Ray) bool {
	oc := vec3.Sub(r.Origin, center)
	a := vec3.Dot(r.Direction, r.Direction)
	b := 2.0 * vec3.Dot(oc, r.Direction)
	c := vec3.Dot(oc, oc) - radius*radius
	discriminant := b*b - 4*a*c
	return discriminant > 0
}

func color(r *ray.Ray) vec3.Vec3 {
	if (hitSphere(vec3.Vec3{X: 0.0, Y: 0.0, Z: -1.0}, 0.5, r)) {
		return vec3.Vec3{X: 1.0, Y: 0.0, Z: 0.0}
	}
	unitDirection := r.Direction.Unit()
	t := 0.5 * (unitDirection.Y + 1.0)
	return vec3.Add(
		vec3.MulCoeff(vec3.Vec3{X: 1.0, Y: 1.0, Z: 1.0}, 1.0-t),
		vec3.MulCoeff(vec3.Vec3{X: 0.5, Y: 0.7, Z: 1.0}, t),
	)
}

func main() {
	const imageWidth = 200
	const imageHeight = 100

	var image strings.Builder
	image.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", imageWidth, imageHeight))

	lower_left := vec3.Vec3{X: -2.0, Y: -1.0, Z: -1.0}
	horizontal := vec3.Vec3{X: 4.0, Y: 0.0, Z: 0.0}
	vertical := vec3.Vec3{X: 0.0, Y: 2.0, Z: 0.0}
	origin := vec3.Vec3{X: 0.0, Y: 0.0, Z: 0.0}

	for j := imageHeight - 1; j >= 0; j -= 1 {
		for i := 0; i < imageWidth; i += 1 {
			u := float64(i) / (imageWidth)
			v := float64(j) / (imageHeight)

			r := ray.Ray{Origin: origin, Direction: vec3.Add3(
				lower_left,
				vec3.MulCoeff(horizontal, u),
				vec3.MulCoeff(vertical, v),
			)}
			color := color(&r)

			ir := int(255.99 * color.X)
			ig := int(255.99 * color.Y)
			ib := int(255.99 * color.Z)

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
