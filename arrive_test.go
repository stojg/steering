package steering

import (
	"github.com/stojg/vector"
	"testing"
)

var arriveSteering *SteeringOutput

func TestArrive_Get(t *testing.T) {
	body := newBody()
	target := vector.NewVector3(100, 0, 100)
	output := NewArrive(body, target, 100)

	actual := output.Get()
	expected := vector.NewVector3(1.224744871391589, 0, 1.224744871391589)
	if !actual.linear.Equals(expected) {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

}

func BenchmarkArrive_Get(b *testing.B) {
	body := newBody()
	target := vector.NewVector3(100, 0, 100)
	output := NewArrive(body, target, 100)

	var tSteering *SteeringOutput
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tSteering = output.Get()
	}
	arriveSteering = tSteering
}
