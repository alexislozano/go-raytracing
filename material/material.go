package material

import (
	"github.com/alexislozano/go-raytracing/ray"
	"github.com/alexislozano/go-raytracing/vec3"
)

type Material interface {
	Scatter(r *ray.Ray, p vec3.Vec3, normal vec3.Vec3) (bool, vec3.Vec3, ray.Ray)
}
