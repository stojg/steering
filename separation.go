package steering

import (
	. "github.com/stojg/vivere/lib/components"
	. "github.com/stojg/vivere/lib/vector"
	//"math"
)

func NewSeparation(m *Model, b *RigidBody, targets []*Vector3 , threshold float64) *Separation {
	return &Separation{
		model: m,
		body: b,
		targets: targets,
		threshold: threshold,
		decayCoefficient: 10000,
	}
}

type Separation struct {
	model        *Model
	body         *RigidBody
	targets     []*Vector3
	threshold float64
	decayCoefficient float64
}

func (s *Separation) Get() *SteeringOutput {
	steering := NewSteeringOutput()
	for _, target := range s.targets {
		direction := s.model.Position().NewSub(target)
		distance := direction.Length()
		if distance > s.threshold {
			continue
		}

		// inverse square
		//strength := math.Min(s.decayCoefficient / (distance * distance), s.body.MaxAcceleration.Length())

		// linear
		strength := s.body.MaxAcceleration.Length() * ((s.threshold - distance) / s.threshold)

		steering.linear.Add(direction.Normalize().Scale(strength))
	}
	return steering
}
