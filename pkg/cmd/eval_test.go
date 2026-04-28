// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/bem-team/bem-cli/internal/mocktest"
)

func TestEvalTriggerEvaluation(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"eval", "trigger-evaluation",
			"--transformation-id", "tr_01HXAB...",
			"--transformation-id", "tr_01HXCD...",
			"--evaluation-version", "0.1.0-gemini",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"transformationIDs:\n" +
			"  - tr_01HXAB...\n" +
			"  - tr_01HXCD...\n" +
			"evaluationVersion: 0.1.0-gemini\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"eval", "trigger-evaluation",
		)
	})
}
