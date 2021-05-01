package camera

import (
	"github.com/alexislozano/go-raytracing/ray"
	"github.com/alexislozano/go-raytracing/vec3"
)

type Camera struct {
	lowerLeftCorner vec3.Vec3
	horizontal      vec3.Vec3
	vertical        vec3.Vec3
	origin          vec3.Vec3
}

func New() Camera {
	return Camera{
		vec3.Vec3{X: -2.0, Y: -1.0, Z: -1.0},
		vec3.Vec3{X: 4.0, Y: 0.0, Z: 0.0},
		vec3.Vec3{X: 0.0, Y: 2.0, Z: 0.0},
		vec3.Vec3{X: 0.0, Y: 0.0, Z: 0.0},
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
