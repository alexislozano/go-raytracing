package material

import (
	"github.com/alexislozano/go-raytracing/ray"
	"github.com/alexislozano/go-raytracing/vec3"
)

type Metal struct {
	Albedo vec3.Vec3
	Fuzz   float64
}

func (l *Metal) Scatter(r *ray.Ray, p vec3.Vec3, normal vec3.Vec3) (bool, vec3.Vec3, ray.Ray) {
	reflected := reflect(r.Direction.Unit(), normal)
	scattered := ray.Ray{
		Origin:    p,
		Direction: vec3.Add(reflected, vec3.MulCoeff(vec3.RandomInUnitSphere(), l.Fuzz)),
	}
	attenuation := l.Albedo
	return vec3.Dot(scattered.Direction, normal) > 0, attenuation, scattered
}
