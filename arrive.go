package steering

import (
	. "github.com/stojg/vivere/lib/components"
	"github.com/stojg/vector"
)

func NewArrive(m *Model, b *RigidBody, target *vector.Vector3, maxSpeed, targetRadius, slowRadius float64) *Arrive {
	return &Arrive{
		model:        m,
		body:         b,
		target:       target,
		targetRadius: targetRadius,
		slowRadius:   slowRadius,
		maxSpeed:     maxSpeed,
	}
}

// Arrive tries to get the character to arrive slowly at a target
type Arrive struct {
	model        *Model
	body         *RigidBody
	target       *vector.Vector3
	targetRadius float64
	slowRadius   float64
	maxSpeed     float64
}

func (s *Arrive) Get() *SteeringOutput {

	const timeToTarget = 0.1

	steering := NewSteeringOutput()

	direction := s.target.NewSub(s.model.Position())
	distance := direction.Length()

	// We have arrived, no output
	if distance < s.targetRadius {
		return steering
	}

	// We are outside the slow radius, so full speed ahead
	var targetSpeed float64
	if distance > s.slowRadius {
		targetSpeed = s.maxSpeed
	} else {
		targetSpeed = s.maxSpeed * distance / s.slowRadius
	}

	// The target velocity combines speed and direction
	targetVelocity := direction
	targetVelocity.Normalize()
	targetVelocity.Scale(targetSpeed)

	// Acceleration tries to get to the target velocity
	steering.linear = targetVelocity.NewSub(s.body.Velocity)
	steering.linear.Scale(1 / timeToTarget)

	// check if acceleration is to fast
	if steering.linear.SquareLength() > s.body.MaxAcceleration.SquareLength() {
		steering.linear.Normalize()
		steering.linear.Scale(s.body.MaxAcceleration.Length())
	}

	return steering
}
