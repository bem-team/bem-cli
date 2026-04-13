// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/bem-team/bem-cli/internal/mocktest"
)

func TestWorkflowsVersionsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workflows:versions", "retrieve",
			"--workflow-name", "workflowName",
			"--version-num", "0",
		)
	})
}

func TestWorkflowsVersionsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workflows:versions", "list",
			"--max-items", "10",
			"--workflow-name", "workflowName",
			"--ending-before", "0",
			"--limit", "1",
			"--sort-order", "asc",
			"--starting-after", "0",
		)
	})
}
