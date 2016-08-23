package steering

import (
	. "github.com/stojg/vivere/lib/components"
)

func NewLookWhereYoureGoing(character *RigidBody, cbody *Model) *LookWhereYoureGoing {
	return &LookWhereYoureGoing{
		model: character,
		cbody: cbody,
	}
}

// LookWhereYoureGoing turns the character so it faces the direction the character is moving
type LookWhereYoureGoing struct {
	model *RigidBody
	cbody *Model
}

// GetSteering returns a angular steering
func (s *LookWhereYoureGoing) Get() *SteeringOutput {
	if s.model.Velocity.Length() == 0 {
		return NewSteeringOutput()
	}
	face := NewFace(s.cbody, s.model, s.model.Velocity.NewAdd(s.cbody.Position()))
	return face.Get()
}
