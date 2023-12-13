package cargo_test

import (
	"testing"

	"github.com/go-kit/examples/shipping/cargo"
)

func BenchmarkNextTrackingID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cargo.NextTrackingID()
	}
}
