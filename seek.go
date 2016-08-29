package steering

import (
	"github.com/stojg/vector"
	. "github.com/stojg/vivere/lib/components"
)

func NewSeek(model *Model, body *RigidBody, target *vector.Vector3) *Seek {
	s := &Seek{
		model:  model,
		body:   body,
		target: target,
	}
	return s
}

// Seek makes the character to go full speed against the target
type Seek struct {
	model  *Model
	body   *RigidBody
	target *vector.Vector3
}

// GetSteering returns a linear steering
func (s *Seek) Get() *SteeringOutput {
	steering := NewSteeringOutput()
	// Get the direction to the target
	steering.linear = s.target.NewSub(s.model.Position())
	// Go full speed ahead
	steering.linear.Normalize()
	steering.linear.HadamardProduct(s.body.MaxAcceleration)
	return steering
}
