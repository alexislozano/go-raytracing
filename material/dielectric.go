package material

import (
	"math/rand"

	"github.com/alexislozano/go-raytracing/ray"
	"github.com/alexislozano/go-raytracing/vec3"
)

type Dielectric struct {
	RefIdx float64
}

func (d *Dielectric) Scatter(r *ray.Ray, p vec3.Vec3, normal vec3.Vec3) (bool, vec3.Vec3, ray.Ray) {
	var outwardNormal vec3.Vec3
	var niOverNt float64
	var cosine float64
	reflected := reflect(r.Direction, normal)
	if vec3.Dot(r.Direction, normal) > 0 {
		outwardNormal = vec3.Neg(normal)
		niOverNt = d.RefIdx
		cosine = d.RefIdx * vec3.Dot(r.Direction, normal) / r.Direction.Length()
	} else {
		outwardNormal = normal
		niOverNt = 1 / d.RefIdx
		cosine = -vec3.Dot(r.Direction, normal) / r.Direction.Length()
	}
	attenuation := vec3.Vec3{X: 1, Y: 1, Z: 1}
	var reflectProb float64
	var scattered ray.Ray
	isRefracted, refracted := refract(r.Direction, outwardNormal, niOverNt)
	if isRefracted {
		reflectProb = schlick(cosine, d.RefIdx)
	} else {
		reflectProb = 1
	}
	if rand.Float64() < reflectProb {
		scattered = ray.Ray{Origin: p, Direction: reflected}
	} else {
		scattered = ray.Ray{Origin: p, Direction: refracted}
	}
	return true, attenuation, scattered
}
