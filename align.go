package steering

import (
	"github.com/stojg/vector"
	"math"
)

type Body interface{
	Position() *vector.Vector3
	Velocity() *vector.Vector3
	MaxAcceleration() *vector.Vector3
	Orientation() *vector.Quaternion
	MaxRotation() float64
	Rotation() *vector.Vector3
}

func NewAlign(c Body, target *vector.Quaternion, targetRadius, slowRadius float64) *Align {
	return &Align{
		char:        c,
		target:       target,
		targetRadius: targetRadius,
		slowRadius:   slowRadius,
	}
}

// Align ensures that the character have the same orientation as the target
type Align struct {
	char         Body
	target       *vector.Quaternion
	targetRadius float64 // 0.02
	slowRadius   float64 // 0.1
}

// GetSteering returns the angular steering to mimic the targets orientation
func (align *Align) Get() *SteeringOutput {

	const timeToTarget = 0.1

	steering := NewSteeringOutput()

	invInitial := &vector.Quaternion{
		R: align.char.Orientation().R,
		I: -align.char.Orientation().I,
		J: -align.char.Orientation().J,
		K: -align.char.Orientation().K,
	}

	q := align.target.NewMultiply(invInitial)

	// protect the ArcCos from numerical instabilities
	if q.R > 1.0 {
		q.R = 1.0
	} else if q.R < -1.0 {
		q.R = -1.0
	}

	theta := 2 * math.Acos(q.R)

	sin := 1 / math.Sin(theta/2)
	axis := &vector.Vector3{
		sin * q.I,
		sin * q.J,
		sin * q.K,
	}

	theta = align.mapToRange(theta)
	thetaNoSign := math.Abs(theta)
	// Check if we are there, return no steering
	if (thetaNoSign) < align.targetRadius {
		return steering
	}

	var targetRotation float64
	if thetaNoSign > align.slowRadius {
		targetRotation = align.char.MaxRotation()
	} else {
		targetRotation = align.char.MaxRotation() * (thetaNoSign / align.slowRadius)
	}

	targetRotation *= theta / thetaNoSign

	axis.Normalize()
	axis.Scale(targetRotation)
	axis.Sub(align.char.Rotation())
	axis.Scale(1 / timeToTarget)

	steering.angular = axis
	return steering
}

func (align *Align) mapToRange(rotation float64) float64 {
	for rotation < -math.Pi {
		rotation += math.Pi * 2
	}
	for rotation > math.Pi {
		rotation -= math.Pi * 2
	}
	return rotation
}
