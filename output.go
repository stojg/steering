package steering

import (
	. "github.com/stojg/vivere/lib/vector"
)

// NewSteeringOutput returns a new zero initialized SteeringOutput
func NewSteeringOutput() *SteeringOutput {
	return &SteeringOutput{
		linear:  &Vector3{},
		angular: &Vector3{},
	}
}

// SteeringOutput describes wished changes in velocity (linear) and rotation (angular)
type SteeringOutput struct {
	linear  *Vector3
	angular *Vector3
}

func (s *SteeringOutput) SetLinear(val *Vector3) {
	s.linear = val
}

func (s *SteeringOutput) Linear() *Vector3 {
	return s.linear
}

func (s *SteeringOutput) SetAngular(val *Vector3) {
	s.angular = val
}

func (s *SteeringOutput) Angular() *Vector3 {
	return s.angular
}




