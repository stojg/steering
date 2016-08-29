package steering

import (
	"github.com/stojg/vector"
	"github.com/stojg/vivere/lib/components"
	"testing"
)

var alignSteering *SteeringOutput

func TestAlign_Get1(t *testing.T) {
	model := components.NewModel(10, 10, 10, 1)
	body := components.NewRidigBody(1)

	target := vector.QuaternionFromVectors(vector.NewVector3(0, 1, 0), vector.NewVector3(0, 0, 1))
	output := NewAlign(model, body, target, 0.1, 0.5)

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
	model := components.NewModel(10, 10, 10, 1)
	body := components.NewRidigBody(1)

	o := vector.QuaternionFromVectors(vector.NewVector3(0, 1, 0), vector.NewVector3(0, 0, 1))
	model.Orientation().Set(o.R, o.I, o.J, o.K)

	target := vector.QuaternionFromVectors(vector.NewVector3(0, 1, 0), vector.NewVector3(0, 0, 1))
	output := NewAlign(model, body, target, 0.1, 0.5)

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
	model := components.NewModel(10, 10, 10, 1)
	body := components.NewRidigBody(1)
	target := vector.QuaternionFromVectors(vector.NewVector3(0, 1, 0), vector.NewVector3(0, 0, 1))
	output := NewAlign(model, body, target, 0.1, 0.5)

	var tSteering *SteeringOutput
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tSteering = output.Get()
	}
	alignSteering = tSteering
}
