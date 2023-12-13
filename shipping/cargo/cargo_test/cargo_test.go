package cargo_test

import (
	"testing"

	"github.com/go-kit/examples/shipping/cargo"
)

// $ go test -bench=. -count=5 | tee v1.txt

func BenchmarkNextTrackingID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cargo.NextTrackingID()
	}
}

func BenchmarkNextTrackingIDV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cargo.NextTrackingIDV2()
	}
}
