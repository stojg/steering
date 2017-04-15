package steering

import (
	"github.com/stojg/vector"
)

func NewSeparation(m Body, targets []*vector.Vector3, threshold float64) *Separation {
	return &Separation{
		body:             m,
		targets:          targets,
		threshold:        threshold,
		decayCoefficient: 10000,
	}
}

type Separation struct {
	body             Body
	targets          []*vector.Vector3
	threshold        float64
	decayCoefficient float64
}

func (s *Separation) Get() *SteeringOutput {
	steering := NewSteeringOutput()
	for _, target := range s.targets {
		direction := s.body.Position().NewSub(target)
		distance := direction.Length()
		if distance > s.threshold {
			continue
		}

		// inverse square
		//strength := math.Min(s.decayCoefficient / (distance * distance), s.body.MaxAcceleration.Length())

		// linear
		strength := s.body.MaxAcceleration().Length() * ((s.threshold - distance) / s.threshold)

		steering.linear.Add(direction.Normalize().Scale(strength))
	}
	return steering
}
