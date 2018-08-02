package candle

import "testing"

func TestVersion(t *testing.T) {
	if want, have := "v2.0.0", Version(); want != have {
		t.Fatalf("Version: want %q, have %q", want, have)
	}
}
