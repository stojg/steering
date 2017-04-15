package steering

import (
	"github.com/stojg/vector"
	"testing"
)

var seekSteering *SteeringOutput

func TestSeek_Get(t *testing.T) {
	body := newBody()

	output := NewSeek(body, vector.NewVector3(100, 0, 100))
	actual := output.Get()
	expected := vector.NewVector3(0.70710, 0, 0.70710)
	if !actual.linear.Equals(expected) {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

}

func BenchmarkSeek_Get(b *testing.B) {
	body := newBody()
	output := NewSeek(body, vector.NewVector3(100, 0, 100))
	var tSteering *SteeringOutput

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tSteering = output.Get()
	}
	seekSteering = tSteering
}
