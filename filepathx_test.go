package filepathx

import (
	"os"
	"testing"
)

func TestGlob_ZeroDoubleStars_oneMatch(t *testing.T) {
	// test passthru to vanilla path/filepath
	path := "./a/b/c.d/e.f"
	err := os.MkdirAll(path, 0755)
	if err != nil {
		t.Fatalf("ZeroDoubleStars: os.MkdirAll: %s", err)
	}
	matches, err := Glob("./*/*/*.d")
	if err != nil {
		t.Fatalf("ZeroDoubleStars: Glob: %s", err)
	}
	if len(matches) != 1 {
		t.Fatalf("ZeroDoubleStars: got %d matches, expected 1", len(matches))
	}
	expected := "a/b/c.d"
	if matches[0] != expected {
		t.Fatalf("ZeroDoubleStars: matched [%s], expected [%s]", matches[0], expected)
	}
}

func TestGlob_OneDoubleStar_oneMatch(t *testing.T) {
	// test a single double-star
	path := "./a/b/c.d/e.f"
	err := os.MkdirAll(path, 0755)
	if err != nil {
		t.Fatalf("OneDoubleStar: os.MkdirAll: %s", err)
	}
	matches, err := Glob("./**/*.f")
	if err != nil {
		t.Fatalf("OneDoubleStar: Glob: %s", err)
	}
	if len(matches) != 1 {
		t.Fatalf("OneDoubleStar: got %d matches, expected 1", len(matches))
	}
	expected := "a/b/c.d/e.f"
	if matches[0] != expected {
		t.Fatalf("OneDoubleStar: matched [%s], expected [%s]", matches[0], expected)
	}
}

func TestGlob_OneDoubleStar_twoMatches(t *testing.T) {
	// test a single double-star
	path := "./a/b/c.d/e.f"
	err := os.MkdirAll(path, 0755)
	if err != nil {
		t.Fatalf("OneDoubleStar: os.MkdirAll: %s", err)
	}
	matches, err := Glob("./a/**/*.*")
	if err != nil {
		t.Fatalf("OneDoubleStar: Glob: %s", err)
	}
	if len(matches) != 2 {
		t.Fatalf("OneDoubleStar: got %d matches, expected 2", len(matches))
	}
	expected := []string{"a/b/c.d", "a/b/c.d/e.f"}
	for i, match := range matches {
		if match != expected[i] {
			t.Fatalf("OneDoubleStar: matched [%s], expected [%s]", match, expected[i])
		}
	}
}

func TestGlob_TwoDoubleStars_oneMatch(t *testing.T) {
	// test two double-stars
	path := "./a/b/c.d/e.f"
	err := os.MkdirAll(path, 0755)
	if err != nil {
		t.Fatalf("TwoDoubleStars: os.MkdirAll: %s", err)
	}
	matches, err := Glob("./**/b/**/*.f")
	if err != nil {
		t.Fatalf("TwoDoubleStars: Glob: %s", err)
	}
	if len(matches) != 1 {
		t.Fatalf("TwoDoubleStars: got %d matches, expected 1", len(matches))
	}
	expected := "a/b/c.d/e.f"
	if matches[0] != expected {
		t.Fatalf("TwoDoubleStars: matched [%s], expected [%s]", matches[0], expected)
	}
}
