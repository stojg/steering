package steering

import (
	. "github.com/stojg/vivere/lib/components"
	. "github.com/stojg/vivere/lib/vector"
	"math"
)

//func NewArrive(m *Model, b *RigidBody, target *Vector3, maxSpeed, targetRadius, slowRadius float64) *Arrive {
func NewFace(m *Model, b *RigidBody, target *Vector3) *Face {
	return &Face{
		model:           m,
		body:            b,
		target:          target,
		baseOrientation: NewQuaternion(1, 0, 0, 0),
	}
}

// Face turns the character so it 'looks' at the target
type Face struct {
	model  *Model
	body   *RigidBody
	target *Vector3
	// @todo fix
	baseOrientation *Quaternion
}

// GetSteering returns a angular steering
func (face *Face) Get() *SteeringOutput {

	// 1. Calculate the target to delegate to align

	// Work out the direction to target
	direction := face.target.NewSub(face.model.Position())

	// Check for zero direction
	if direction.SquareLength() == 0 {
		return NewSteeringOutput()
	}

	orientation := face.calculateOrientation(direction)
	align := NewAlign(face.model, face.body, orientation, 0.1, 1)
	return align.Get()
}

func (face *Face) calculateOrientation(vector *Vector3) *Quaternion {
	vector.Normalize()

	baseZVector := VectorX().Rotate(face.baseOrientation)

	if baseZVector.Equals(vector) {
		return face.baseOrientation.Clone()
	}
	if baseZVector.Equals(vector.NewInverse()) {
		// @todo need to fix this is the base orientation isn't 1,0,0,0?
		return NewQuaternion(0, 0, 1, 0)
	}

	// find the minimal rotation from the base to the target
	angle := math.Acos(baseZVector.Dot(vector))
	axis := baseZVector.NewCross(vector).Normalize()

	return QuaternionFromAxisAngle(axis, angle)
}
