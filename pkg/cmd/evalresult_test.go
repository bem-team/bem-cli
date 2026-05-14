// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/bem-team/bem-cli/internal/mocktest"
)

func TestEvalResultsRetrieveResults(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"eval:results", "retrieve-results",
			"--evaluation-version", "evaluationVersion",
			"--event-ids", "eventIDs",
			"--transformation-ids", "transformationIDs",
		)
	})
}
