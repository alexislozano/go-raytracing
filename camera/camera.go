package camera

import (
	"math"

	"github.com/alexislozano/go-raytracing/ray"
	"github.com/alexislozano/go-raytracing/vec3"
)

type Camera struct {
	lowerLeftCorner vec3.Vec3
	horizontal      vec3.Vec3
	vertical        vec3.Vec3
	origin          vec3.Vec3
}

func New(lookFrom vec3.Vec3, lookAt vec3.Vec3, vUp vec3.Vec3, vFov float64, aspect float64) Camera {
	theta := vFov * math.Pi / 180
	halfHeight := math.Tan(theta / 2)
	halfWidth := aspect * halfHeight
	w := vec3.Sub(lookFrom, lookAt).Unit()
	u := vec3.Cross(vUp, w).Unit()
	v := vec3.Cross(w, u)
	return Camera{
		vec3.Add4(
			lookFrom,
			vec3.Neg(vec3.MulCoeff(u, halfWidth)),
			vec3.Neg(vec3.MulCoeff(v, halfHeight)),
			vec3.Neg(w),
		),
		vec3.MulCoeff(u, 2*halfWidth),
		vec3.MulCoeff(v, 2*halfHeight),
		lookFrom,
	}
}

func (c *Camera) GetRay(u float64, v float64) ray.Ray {
	return ray.Ray{
		Origin: c.origin,
		Direction: vec3.Add4(
			c.lowerLeftCorner,
			vec3.MulCoeff(c.horizontal, u),
			vec3.MulCoeff(c.vertical, v),
			vec3.Neg(c.origin),
		),
	}
}
