package material

import (
	"math"

	"github.com/alexislozano/go-raytracing/ray"
	"github.com/alexislozano/go-raytracing/vec3"
)

type Material interface {
	Scatter(r *ray.Ray, p vec3.Vec3, normal vec3.Vec3) (bool, vec3.Vec3, ray.Ray)
}

func reflect(v vec3.Vec3, n vec3.Vec3) vec3.Vec3 {
	return vec3.Sub(v, vec3.MulCoeff(n, 2*vec3.Dot(v, n)))
}

func refract(v vec3.Vec3, n vec3.Vec3, niOverNt float64) (bool, vec3.Vec3) {
	uv := v.Unit()
	dt := vec3.Dot(uv, n)
	discriminant := 1.0 - niOverNt*niOverNt*(1-dt*dt)
	if discriminant > 0 {
		refracted := vec3.Sub(
			vec3.MulCoeff(vec3.Sub(uv, vec3.MulCoeff(n, dt)), niOverNt),
			vec3.MulCoeff(n, math.Sqrt(discriminant)),
		)
		return true, refracted
	} else {
		return false, vec3.Vec3{X: 0, Y: 0, Z: 0}
	}
}

func schlick(cosine float64, refIdx float64) float64 {
	r0 := (1 - refIdx) / (1 + refIdx)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow(1-cosine, 5)
}
