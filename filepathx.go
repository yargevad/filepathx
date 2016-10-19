// Package filepathx is a superset of the core filepath package.
// It adds double-star globbing, where "**" represents a recursive wildcard matching zero-or-more directory levels deep.
package filepathx

import (
	"os"
	"path/filepath"
	"strings"
)

type Globs []string

// Glob adds double-star support to the core path/filepath.Glob function.
// It is useful when your globs might have double-stars, but you're not sure.
func Glob(pat string) ([]string, error) {
	if !strings.Contains(pat, "**") {
		// passthru to core package if no double-star
		return filepath.Glob(pat)
	}
	return Globs(strings.Split(pat, "**")).Expand()
}

// Expand finds matches for the provided list of glob patterns.
// A double-star pattern is assumed between each pair of array elements.
func (globs Globs) Expand() (matches []string, err error) {
	var prefixes = []string{""} // accumulate here
	for _, glob := range globs {
		var hits []string
		var hitMap = map[string]bool{}
		for _, prefix := range prefixes {
			matches, err := filepath.Glob(prefix + glob)
			if err != nil {
				return nil, err
			}
			for _, match := range matches {
				err = filepath.Walk(match, func(path string, info os.FileInfo, err error) error {
					if err != nil {
						return err
					}
					// save deduped match from current iteration
					if _, ok := hitMap[path]; !ok {
						hits = append(hits, path)
						hitMap[path] = true
					}
					return nil
				})
				if err != nil {
					return nil, err
				}
			}
		}
		prefixes = hits
	}

	return prefixes, nil
}
