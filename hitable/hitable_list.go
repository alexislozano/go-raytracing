package hitable

import (
	"github.com/alexislozano/go-raytracing/ray"
	"github.com/alexislozano/go-raytracing/vec3"
)

type HitableList struct {
	List []Hitable
}

func (hl *HitableList) Hit(r *ray.Ray, tMin float64, tMax float64) (bool, HitRecord) {
	hitAnything := false
	closestSoFar := tMax
	rec := HitRecord{
		closestSoFar,
		vec3.Vec3{X: 0.0, Y: 0.0, Z: 0.0},
		vec3.Vec3{X: 0.0, Y: 0.0, Z: 0.0},
	}
	for _, hitable := range hl.List {
		hit, newRec := hitable.Hit(r, tMin, closestSoFar)
		if hit {
			hitAnything = true
			closestSoFar = newRec.t
			rec = newRec
		}
	}
	return hitAnything, rec
}
