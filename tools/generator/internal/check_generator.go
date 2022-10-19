package internal

import (
	_ "embed"
	"regexp"
)

//go:embed patches.go
var patchesGo []byte

var rxMethodId = regexp.MustCompile(`"([a-z]+_[a-z_]+)"`)

// check the generator files for old ids.
func (gen *Generator) checkGeneratorOldIds(methodIds map[string]bool) {
	matches := rxMethodId.FindAllSubmatch(patchesGo, -1)
	for _, match := range matches {
		if _, ok := methodIds[string(match[1])]; !ok {
			gen.printf("    '%s' does not exist anymore, remove it from 'patches.go'.\n", string(match[1]))
		}
	}
}
