package glass

import "testing"

func TestGlass(t *testing.T) {
	want := "I can eat glass and it doesn't hurt me."
	if got := glass(); got != want {
		t.Errorf("glass() = %q, want %q", got, want)
	}
}
