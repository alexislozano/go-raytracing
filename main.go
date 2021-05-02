package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"

	"github.com/alexislozano/go-raytracing/camera"
	"github.com/alexislozano/go-raytracing/hitable"
	"github.com/alexislozano/go-raytracing/material"
	"github.com/alexislozano/go-raytracing/ray"
	"github.com/alexislozano/go-raytracing/vec3"
)

func color(r *ray.Ray, world hitable.Hitable, depth int) vec3.Vec3 {
	hit, rec, mat := world.Hit(r, 0.001, math.MaxFloat64)
	if hit {
		if depth < 50 {
			isScattered, attenuation, scattered := mat.Scatter(r, rec.P, rec.Normal)
			if isScattered {
				return vec3.Mul(attenuation, color(&scattered, world, depth+1))
			}
		}
		return vec3.Vec3{X: 0, Y: 0, Z: 0}
	} else {
		unitDirection := r.Direction.Unit()
		t := 0.5 * (unitDirection.Y + 1.0)
		return vec3.Add(
			vec3.MulCoeff(vec3.Vec3{X: 1.0, Y: 1.0, Z: 1.0}, 1.0-t),
			vec3.MulCoeff(vec3.Vec3{X: 0.5, Y: 0.7, Z: 1.0}, t),
		)
	}
}

func main() {
	const imageWidth = 200
	const imageHeight = 100
	const raysPerPixel = 100

	var image strings.Builder
	image.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", imageWidth, imageHeight))

	list := []hitable.Hitable{
		&hitable.Sphere{
			Center:   vec3.Vec3{X: 0.0, Y: 0.0, Z: -1.0},
			Radius:   0.5,
			Material: &material.Lambertian{Albedo: vec3.Vec3{X: 0.8, Y: 0.3, Z: 0.3}},
		},
		&hitable.Sphere{
			Center:   vec3.Vec3{X: 0.0, Y: -100.5, Z: -1.0},
			Radius:   100,
			Material: &material.Lambertian{Albedo: vec3.Vec3{X: 0.8, Y: 0.8, Z: 0.0}},
		},
		&hitable.Sphere{
			Center: vec3.Vec3{X: 1, Y: 0, Z: -1},
			Radius: 0.5,
			Material: &material.Metal{
				Albedo: vec3.Vec3{X: 0.8, Y: 0.6, Z: 0.2},
				Fuzz:   0.3,
			},
		},
		&hitable.Sphere{
			Center: vec3.Vec3{X: -1, Y: 0, Z: -1},
			Radius: 0.5,
			Material: &material.Metal{
				Albedo: vec3.Vec3{X: 0.8, Y: 0.8, Z: 0.8},
				Fuzz:   1,
			},
		},
	}
	world := hitable.HitableList{List: list}
	cam := camera.New()

	for j := imageHeight - 1; j >= 0; j-- {
		for i := 0; i < imageWidth; i++ {
			col := vec3.Vec3{X: 0, Y: 0, Z: 0}
			for s := 0; s < raysPerPixel; s++ {
				u := (float64(i) + rand.Float64()) / imageWidth
				v := (float64(j) + rand.Float64()) / imageHeight
				r := cam.GetRay(u, v)
				col = vec3.Add(col, color(&r, &world, 0.0))
			}

			col = vec3.DivCoeff(col, raysPerPixel)
			col = vec3.Vec3{
				X: math.Sqrt(col.X),
				Y: math.Sqrt(col.Y),
				Z: math.Sqrt(col.Z),
			}

			ir := int(255.99 * col.X)
			ig := int(255.99 * col.Y)
			ib := int(255.99 * col.Z)

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
