package hitable

import (
	"math"

	"github.com/alexislozano/go-raytracing/ray"
	"github.com/alexislozano/go-raytracing/vec3"
)

type Sphere struct {
	Center vec3.Vec3
	Radius float64
}

func (s *Sphere) Hit(r *ray.Ray, tMin float64, tMax float64) (bool, HitRecord) {
	rec := HitRecord{
		tMax,
		vec3.Vec3{X: 0.0, Y: 0.0, Z: 0.0},
		vec3.Vec3{X: 0.0, Y: 0.0, Z: 0.0},
	}
	oc := vec3.Sub(r.Origin, s.Center)
	a := vec3.Dot(r.Direction, r.Direction)
	b := vec3.Dot(oc, r.Direction)
	c := vec3.Dot(oc, oc) - s.Radius*s.Radius
	discriminant := b*b - a*c
	if discriminant > 0 {
		temp := (-b - math.Sqrt(b*b-a*c)) / a
		if temp < tMax && temp > tMin {
			rec.t = temp
			rec.P = r.Point(temp)
			rec.Normal = vec3.DivCoeff(vec3.Sub(rec.P, s.Center), s.Radius)
			return true, rec
		}
		temp = (-b + math.Sqrt(b*b-a*c)) / a
		if temp < tMax && temp > tMin {
			rec.t = temp
			rec.P = r.Point(temp)
			rec.Normal = vec3.DivCoeff(vec3.Sub(rec.P, s.Center), s.Radius)
			return true, rec
		}
	}
	return false, rec
}
