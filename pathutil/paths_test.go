package pathutil

import (
	"testing"
)

func TestName(t *testing.T) {
	if Name() != "github.com/Masterminds/gopt" {
		t.Fatalf("Got path %s", Name())
	}
}
