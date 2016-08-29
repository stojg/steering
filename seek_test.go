package steering

import (
	"github.com/stojg/vector"
	"github.com/stojg/vivere/lib/components"
	"testing"
)

var seekSteering *SteeringOutput

func TestSeek_Get(t *testing.T) {
	model := components.NewModel(10, 10, 10, 1)
	body := components.NewRidigBody(1)

	output := NewSeek(model, body, vector.NewVector3(100, 0, 100))
	actual := output.Get()
	expected := vector.NewVector3(0.70710, 0, 0.70710)
	if !actual.linear.Equals(expected) {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

}

func BenchmarkSeek_Get(b *testing.B) {
	model := components.NewModel(10, 10, 10, 1)
	body := components.NewRidigBody(1)
	output := NewSeek(model, body, vector.NewVector3(100, 0, 100))
	var tSteering *SteeringOutput

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tSteering = output.Get()
	}
	seekSteering = tSteering
}
