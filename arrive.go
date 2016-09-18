package steering

import (
	"github.com/stojg/vector"
)

func NewArrive(source Body, target *vector.Vector3, maxSpeed, targetRadius, slowRadius float64) *Arrive {
	return &Arrive{
		object:         source,
		target:       target,
		targetRadius: targetRadius,
		slowRadius:   slowRadius,
		maxSpeed:     maxSpeed,
	}
}

// Arrive tries to get the character to arrive slowly at a target
type Arrive struct {
	object       Body
	target       *vector.Vector3
	targetRadius float64
	slowRadius   float64
	maxSpeed     float64
}

func (s *Arrive) Get() *SteeringOutput {

	const timeToTarget = 0.1

	steering := NewSteeringOutput()

	direction := s.target.NewSub(s.object.Position())
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
	direction.Normalize()
	direction.Scale(targetSpeed)

	// Acceleration tries to get to the target velocity
	direction.NewSub(s.object.Velocity())
	direction.Scale(1 / timeToTarget)

	// check if acceleration is to fast
	if direction.SquareLength() > s.object.MaxAcceleration().SquareLength() {
		direction.Normalize()
		direction.Scale(s.object.MaxAcceleration().Length())
	}

	steering.linear = direction
	return steering
}
