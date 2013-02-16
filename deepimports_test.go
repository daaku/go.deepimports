package deepimports_test

import (
	"github.com/daaku/go.deepimports"
	"testing"
)

const aPath = "github.com/daaku/go.deepimports/_test/a"
const bPath = "github.com/daaku/go.deepimports/_test/b"

func TestFind(t *testing.T) {
	paths := []string{aPath}
	r, err := deepimports.Find(paths, "")
	if err != nil {
		t.Fatal(err)
	}
	if len(r) != 2 {
		t.Fatalf("expected 2 imports got %d", len(r))
	}
	if r[0].ImportPath != aPath {
		t.Fatalf("expected %s got %s", aPath, r[0].ImportPath)
	}
	if r[1].ImportPath != bPath {
		t.Fatalf("expected %s got %s", bPath, r[1].ImportPath)
	}
}

func TestFindImportsOnly(t *testing.T) {
	paths := []string{"github.com/daaku/go.deepimports/_test/a"}
	r, err := deepimports.FindImportsOnly(paths, "")
	if err != nil {
		t.Fatal(err)
	}
	if len(r) != 1 {
		t.Fatalf("expected 1 imports got %d", len(r))
	}
	if r[0].ImportPath != bPath {
		t.Fatalf("expected %s got %s", bPath, r[0].ImportPath)
	}
}
