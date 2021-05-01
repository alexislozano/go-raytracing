package hitable

import "github.com/alexislozano/go-raytracing/vec3"

type HitRecord struct {
	t      float64
	p      vec3.Vec3
	Normal vec3.Vec3
}
