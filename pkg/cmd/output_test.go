// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/bem-team/bem-cli/internal/mocktest"
)

func TestOutputsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"outputs", "retrieve",
			"--event-id", "eventID",
		)
	})
}

func TestOutputsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"outputs", "list",
			"--max-items", "10",
			"--call-id", "string",
			"--ending-before", "endingBefore",
			"--function-id", "string",
			"--function-name", "string",
			"--include-intermediate=true",
			"--limit", "1",
			"--reference-id", "string",
			"--reference-id-substring", "referenceIDSubstring",
			"--sort-order", "asc",
			"--starting-after", "startingAfter",
			"--workflow-id", "string",
			"--workflow-name", "string",
		)
	})
}
