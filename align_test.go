package steering

import (
	"github.com/stojg/vector"
	"testing"
)

type testBody struct {
	position    *vector.Vector3
	velocity    *vector.Vector3
	maxVelocity float64
	maxAcc      *vector.Vector3
	orientation *vector.Quaternion
	maxRotation float64
	rotation    *vector.Vector3
}

func (b *testBody) Position() *vector.Vector3 {
	return b.position
}
func (b *testBody) Velocity() *vector.Vector3 {
	return b.velocity
}
func (b *testBody) MaxVelocity() float64 {
	return b.maxVelocity
}
func (b *testBody) MaxAcceleration() *vector.Vector3 {
	return b.maxAcc
}
func (b *testBody) Orientation() *vector.Quaternion {
	return b.orientation
}
func (b *testBody) MaxRotation() float64 {
	return b.maxRotation
}
func (b *testBody) Rotation() *vector.Vector3 {
	return b.rotation
}

func newBody() *testBody {
	return &testBody{
		position:    vector.NewVector3(0, 0, 0),
		velocity:    vector.NewVector3(0, 0, 0),
		maxAcc:      vector.NewVector3(1, 1, 1),
		maxVelocity: 1,
		maxRotation: 3.14 / 10,
		orientation: vector.NewQuaternion(1, 0, 0, 0),
		rotation:    vector.NewVector3(0, 0, 0),
	}
}

var alignSteering *SteeringOutput

func TestAlign_Get1(t *testing.T) {
	body := newBody()

	target := vector.QuaternionFromVectors(vector.NewVector3(0, 1, 0), vector.NewVector3(0, 0, 1))
	output := NewAlign(body, target, 0.1, 0.5)

	actual := output.Get()
	expected := vector.NewVector3(0, 0, 0)
	if !actual.linear.Equals(expected) {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

	expectedAngular := vector.NewVector3(3.14, 0, 0)
	if !actual.angular.Equals(expectedAngular) {
		t.Errorf("Expected %s, got %s", expectedAngular, actual)
	}
}

func TestAlign_Get2(t *testing.T) {
	body := newBody()

	o := vector.QuaternionFromVectors(vector.NewVector3(0, 1, 0), vector.NewVector3(0, 0, 1))
	body.Orientation().Set(o.R, o.I, o.J, o.K)

	target := vector.QuaternionFromVectors(vector.NewVector3(0, 1, 0), vector.NewVector3(0, 0, 1))
	output := NewAlign(body, target, 0.1, 0.5)

	actual := output.Get()
	expected := vector.NewVector3(0, 0, 0)
	if !actual.linear.Equals(expected) {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

	expectedAngular := vector.NewVector3(0, 0, 0)
	if !actual.angular.Equals(expectedAngular) {
		t.Errorf("Expected %s, got %s", expectedAngular, actual)
	}
}

func BenchmarkAlign_Get(b *testing.B) {
	body := newBody()
	target := vector.QuaternionFromVectors(vector.NewVector3(0, 1, 0), vector.NewVector3(0, 0, 1))
	output := NewAlign(body, target, 0.1, 0.5)

	var tSteering *SteeringOutput
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tSteering = output.Get()
	}
	alignSteering = tSteering
}
