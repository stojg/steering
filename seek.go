package steering

import (
	"github.com/stojg/vector"
)

func NewSeek(body Body, target *vector.Vector3) *Seek {
	s := &Seek{
		body:  body,
		target: target,
	}
	return s
}

// Seek makes the character to go full speed against the target
type Seek struct {
	body   Body
	target *vector.Vector3
}

// GetSteering returns a linear steering
func (s *Seek) Get() *SteeringOutput {
	steering := NewSteeringOutput()
	// Get the direction to the target
	steering.linear = s.target.NewSub(s.body.Position())
	// Go full speed ahead
	steering.linear.Normalize()
	steering.linear.HadamardProduct(s.body.MaxAcceleration())
	return steering
}
