package steering

import (
	"github.com/stojg/vector"
	"github.com/stojg/vivere/lib/components"
	"testing"
)

var arriveSteering *SteeringOutput

func TestArrive_Get(t *testing.T) {
	model := components.NewModel(10, 10, 10, 1)
	body := components.NewRidigBody(1)
	target := vector.NewVector3(100, 0, 100)
	output := NewArrive(model, body, target, 100, 0.1, 10)

	actual := output.Get()
	expected := vector.NewVector3(1.224744871391589, 0, 1.224744871391589)
	if !actual.linear.Equals(expected) {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

}

func BenchmarkArrive_Get(b *testing.B) {
	model := components.NewModel(10, 10, 10, 1)
	body := components.NewRidigBody(1)
	target := vector.NewVector3(100, 0, 100)
	output := NewArrive(model, body, target, 100, 0.1, 200)

	var tSteering *SteeringOutput
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tSteering = output.Get()
	}
	arriveSteering = tSteering
}
