package filepathx

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGlob_ZeroDoubleStars_oneMatch(t *testing.T) {
	// test passthru to vanilla path/filepath
	path := "./a/b/c.d/e.f"
	err := os.MkdirAll(path, 0755)
	if err != nil {
		t.Fatalf("os.MkdirAll: %s", err)
	}
	matches, err := Glob("./*/*/*.d")
	if err != nil {
		t.Fatalf("Glob: %s", err)
	}
	if len(matches) != 1 {
		t.Fatalf("got %d matches, expected 1", len(matches))
	}
	expected := filepath.Join("a", "b", "c.d")
	if matches[0] != expected {
		t.Fatalf("matched [%s], expected [%s]", matches[0], expected)
	}
}

func TestGlob_OneDoubleStar_oneMatch(t *testing.T) {
	// test a single double-star
	path := "./a/b/c.d/e.f"
	err := os.MkdirAll(path, 0755)
	if err != nil {
		t.Fatalf("os.MkdirAll: %s", err)
	}
	matches, err := Glob("./**/*.f")
	if err != nil {
		t.Fatalf("Glob: %s", err)
	}
	if len(matches) != 1 {
		t.Fatalf("got %d matches, expected 1", len(matches))
	}
	expected := filepath.Join("a", "b", "c.d", "e.f")
	if matches[0] != expected {
		t.Fatalf("matched [%s], expected [%s]", matches[0], expected)
	}
}

func TestGlob_OneDoubleStar_twoMatches(t *testing.T) {
	// test a single double-star
	path := "./a/b/c.d/e.f"
	err := os.MkdirAll(path, 0755)
	if err != nil {
		t.Fatalf("os.MkdirAll: %s", err)
	}
	matches, err := Glob("./a/**/*.*")
	if err != nil {
		t.Fatalf("Glob: %s", err)
	}
	if len(matches) != 2 {
		t.Fatalf("got %d matches, expected 2", len(matches))
	}
	expected := []string{filepath.Join("a", "b", "c.d"), filepath.Join("a", "b", "c.d", "e.f")}
	for i, match := range matches {
		if match != expected[i] {
			t.Fatalf("matched [%s], expected [%s]", match, expected[i])
		}
	}
}

func TestGlob_TwoDoubleStars_oneMatch(t *testing.T) {
	// test two double-stars
	path := "./a/b/c.d/e.f"
	err := os.MkdirAll(path, 0755)
	if err != nil {
		t.Fatalf("os.MkdirAll: %s", err)
	}
	matches, err := Glob("./**/b/**/*.f")
	if err != nil {
		t.Fatalf("Glob: %s", err)
	}
	if len(matches) != 1 {
		t.Fatalf("got %d matches, expected 1", len(matches))
	}
	expected := filepath.Join("a", "b", "c.d", "e.f")
	if matches[0] != expected {
		t.Fatalf("matched [%s], expected [%s]", matches[0], expected)
	}
}

func TestExpand_DirectCall_emptySlice(t *testing.T) {
	var empty []string
	matches, err := Globs(empty).Expand()
	if err != nil {
		t.Fatalf("Glob: %s", err)
	}
	if len(matches) != 0 {
		t.Fatalf("got %d matches, expected 0", len(matches))
	}
}

func TestGlob_Stars_oneMatch(t *testing.T) {
	// test two double-stars
	path := "./a/b/c.d/e.f"
	err := os.MkdirAll(path, 0755)
	if err != nil {
		t.Fatalf("os.MkdirAll: %s", err)
	}
	matches, err := Glob("./*/b/**/*.f")
	if err != nil {
		t.Fatalf("Glob: %s", err)
	}
	if len(matches) != 1 {
		t.Fatalf("got %d matches, expected 1", len(matches))
	}
	expected := filepath.Join("a", "b", "c.d", "e.f")
	if matches[0] != expected {
		t.Fatalf("matched [%s], expected [%s]", matches[0], expected)
	}
}

func TestGlob_SingleDoubleStarAbsFolder(t *testing.T) {
	// Test with absolute path.
	tmp, err := os.MkdirTemp("", "SingleDoubleStarAbsFolder-")
	if err != nil {
		t.Fatalf("os.CreateTemp: %s", err)
	}
	defer os.RemoveAll(tmp)

	t.Log(tmp)
	path := filepath.Join(tmp, "a/b/c.d/e.f/g")
	err = os.MkdirAll(path, 0755)
	if err != nil {
		t.Fatalf("os.MkdirAll: %s", err)
	}
	matches, err := Glob(filepath.Join(tmp, "a/b/*/e.f/**"))
	if err != nil {
		t.Fatalf("Glob: %s", err)
	}
	expected := []string{filepath.Join(tmp, "a", "b", "c.d", "e.f"), filepath.Join(tmp, "a", "b", "c.d", "e.f", "g")}
	if len(matches) != len(expected) {
		t.Fatalf("got %d matches, expected %d", len(matches), len(expected))
	}
	for i, match := range matches {
		if match != expected[i] {
			t.Fatalf("matched [%s], expected [%s]", match, expected[i])
		}
	}
}
