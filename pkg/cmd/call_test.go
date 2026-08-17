// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/bem-team/bem-cli/internal/mocktest"
)

func TestCallsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls", "retrieve",
			"--call-id", "callID",
		)
	})
}

func TestCallsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls", "list",
			"--max-items", "10",
			"--call-id", "string",
			"--call-type", "workflow",
			"--ending-before", "endingBefore",
			"--function-id", "string",
			"--function-name", "string",
			"--limit", "1",
			"--reference-id", "string",
			"--reference-id-substring", "referenceIDSubstring",
			"--sort-order", "asc",
			"--starting-after", "startingAfter",
			"--status", "pending",
			"--workflow-id", "string",
			"--workflow-name", "string",
		)
	})
}

func TestCallsRetrieveTrace(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"calls", "retrieve-trace",
			"--call-id", "callID",
		)
	})
}
