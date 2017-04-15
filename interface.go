package steering

import "github.com/stojg/vector"

type Body interface {
	Position() *vector.Vector3
	Velocity() *vector.Vector3
	MaxVelocity() float64
	MaxAcceleration() *vector.Vector3
	Orientation() *vector.Quaternion
	MaxRotation() float64
	Rotation() *vector.Vector3
}
