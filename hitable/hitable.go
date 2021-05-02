package hitable

import (
	"github.com/alexislozano/go-raytracing/material"
	"github.com/alexislozano/go-raytracing/ray"
)

type Hitable interface {
	Hit(r *ray.Ray, tMin float64, tMax float64) (bool, HitRecord, material.Material)
}
