// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/bem-cli/internal/mocktest"
)

func TestFunctionsCopyCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"functions:copy", "create",
			"--source-function-name", "sourceFunctionName",
			"--target-function-name", "targetFunctionName",
			"--tag", "string",
			"--target-display-name", "targetDisplayName",
			"--target-environment", "targetEnvironment",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"sourceFunctionName: sourceFunctionName\n" +
			"targetFunctionName: targetFunctionName\n" +
			"tags:\n" +
			"  - string\n" +
			"targetDisplayName: targetDisplayName\n" +
			"targetEnvironment: targetEnvironment\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"functions:copy", "create",
		)
	})
}
