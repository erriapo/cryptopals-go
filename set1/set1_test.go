package set1

import (
	"testing"
)

func TestChallenge(t *testing.T) {
    got := Base64("12345")
    if got != "2" {
	t.Errorf("Base64(%q) == %q, want %q", "12345", got, "2")
    }
}
