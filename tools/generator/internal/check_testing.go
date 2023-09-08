package internal

import (
	"bytes"
	"os"
)

func (gen *Generator) checkClientTest(methodIds map[string]bool) error {
	test, err := os.ReadFile(gen.outputPath + "client_test.go")
	if err != nil {
		return err
	}
	for id, deprecated := range methodIds {
		testName := snakeToCamel(id)
		isTested := bytes.Contains(test, []byte(testName))
		if deprecated && isTested {
			gen.printf("    '%s' is deprecated, but tested. Remove the test.\n", testName)
			continue
		}
		if !isTested {
			gen.printf("    '%s' is not tested. Add a test case or skip test via 'skipTest'.\n", testName)
		}
	}
	return nil
}
