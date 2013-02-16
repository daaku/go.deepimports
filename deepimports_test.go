package deepimports_test

import (
	"github.com/daaku/go.deepimports"
	"testing"
)

func TestFind(t *testing.T) {
	paths := []string{"github.com/daaku/go.deepimports/_test/a"}
	r, err := deepimports.Find(paths, "")
	if err != nil {
		t.Fatal(err)
	}
	if len(r) != 2 {
		t.Fatalf("expected 2 imports got %d", len(r))
	}
}
