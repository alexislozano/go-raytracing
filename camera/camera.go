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
	u               vec3.Vec3
	v               vec3.Vec3
	w               vec3.Vec3
	lensRadius      float64
}

func New(lookFrom vec3.Vec3, lookAt vec3.Vec3, vUp vec3.Vec3, vFov float64, aspect float64, aperture float64, focusDist float64) Camera {
	theta := vFov * math.Pi / 180
	halfHeight := math.Tan(theta / 2)
	halfWidth := aspect * halfHeight
	w := vec3.Sub(lookFrom, lookAt).Unit()
	u := vec3.Cross(vUp, w).Unit()
	v := vec3.Cross(w, u)
	return Camera{
		vec3.Add4(
			lookFrom,
			vec3.MulCoeff(u, -halfWidth*focusDist),
			vec3.MulCoeff(v, -halfHeight*focusDist),
			vec3.MulCoeff(w, -focusDist),
		),
		vec3.MulCoeff(u, 2*halfWidth*focusDist),
		vec3.MulCoeff(v, 2*halfHeight*focusDist),
		lookFrom,
		u,
		v,
		w,
		aperture / 2,
	}
}

func (c *Camera) GetRay(s float64, t float64) ray.Ray {
	rD := vec3.MulCoeff(vec3.RandomInUnitDisk(), c.lensRadius)
	offset := vec3.Add(vec3.MulCoeff(c.u, rD.X), vec3.MulCoeff(c.v, rD.Y))
	return ray.Ray{
		Origin: vec3.Add(c.origin, offset),
		Direction: vec3.Add4(
			c.lowerLeftCorner,
			vec3.MulCoeff(c.horizontal, s),
			vec3.MulCoeff(c.vertical, t),
			vec3.Neg(vec3.Add(c.origin, offset)),
		),
	}
}
