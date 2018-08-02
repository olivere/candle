package candle

import "testing"

func TestVersion(t *testing.T) {
	if want, have := "v0.1.1", Version(); want != have {
		t.Fatalf("Version: want %q, have %q", want, have)
	}
}
