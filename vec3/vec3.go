package vec3

import "math"

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func (v Vec3) SquaredLength() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.SquaredLength())
}

func Add(v1 Vec3, v2 Vec3) Vec3 {
	return Vec3{v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z}
}

func Sub(v1 Vec3, v2 Vec3) Vec3 {
	return Vec3{v1.X - v2.X, v1.Y - v2.Y, v1.Z - v2.Z}
}

func Mul(v1 Vec3, v2 Vec3) Vec3 {
	return Vec3{v1.X * v2.X, v1.Y * v2.Y, v1.Z * v2.Z}
}

func Div(v1 Vec3, v2 Vec3) Vec3 {
	return Vec3{v1.X / v2.X, v1.Y / v2.Y, v1.Z / v2.Z}
}

func MulCoeff(v1 Vec3, t float64) Vec3 {
	return Vec3{v1.X * t, v1.Y * t, v1.Z * t}
}

func DivCoeff(v1 Vec3, t float64) Vec3 {
	return Vec3{v1.X / t, v1.Y / t, v1.Z / t}
}

func (v Vec3) Unit() Vec3 {
	return DivCoeff(v, v.Length())
}

func Dot(v1 Vec3, v2 Vec3) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

func Cross(v1 Vec3, v2 Vec3) Vec3 {
	return Vec3{
		v1.Y*v2.Z - v1.Z*v2.Y,
		v1.Z*v2.X - v1.X*v2.Z,
		v1.X*v2.Y - v1.Y*v2.X,
	}
}
