package material

import (
	"github.com/alexislozano/go-raytracing/ray"
	"github.com/alexislozano/go-raytracing/vec3"
)

type Lambertian struct {
	Albedo vec3.Vec3
}

func (l *Lambertian) Scatter(r *ray.Ray, p vec3.Vec3, normal vec3.Vec3) (bool, vec3.Vec3, ray.Ray) {
	target := vec3.Add3(p, normal, vec3.RandomInUnitSphere())
	scattered := ray.Ray{Origin: p, Direction: vec3.Sub(target, p)}
	attenuation := l.Albedo
	return true, attenuation, scattered
}
