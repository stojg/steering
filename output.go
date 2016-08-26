package steering

import (
	"github.com/stojg/vector"
)

// NewSteeringOutput returns a new zero initialized SteeringOutput
func NewSteeringOutput() *SteeringOutput {
	return &SteeringOutput{
		linear:  &vector.Vector3{},
		angular: &vector.Vector3{},
	}
}

// SteeringOutput describes wished changes in velocity (linear) and rotation (angular)
type SteeringOutput struct {
	linear  *vector.Vector3
	angular *vector.Vector3
}

func (s *SteeringOutput) SetLinear(val *vector.Vector3) {
	s.linear = val
}

func (s *SteeringOutput) Linear() *vector.Vector3 {
	return s.linear
}

func (s *SteeringOutput) SetAngular(val *vector.Vector3) {
	s.angular = val
}

func (s *SteeringOutput) Angular() *vector.Vector3 {
	return s.angular
}




