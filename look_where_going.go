package steering

import (
)


func NewLookWhereYoureGoing(character Body) *LookWhereYoureGoing {
	return &LookWhereYoureGoing{
		char: character,
	}
}

// LookWhereYoureGoing turns the character so it faces the direction the character is moving
type LookWhereYoureGoing struct {
	char Body
}

// GetSteering returns a angular steering
func (s *LookWhereYoureGoing) Get() *SteeringOutput {
	if s.char.Velocity().Length() == 0 {
		return NewSteeringOutput()
	}
	face := NewFace(s.char, s.char.Velocity().NewAdd(s.char.Position()))
	return face.Get()
}
