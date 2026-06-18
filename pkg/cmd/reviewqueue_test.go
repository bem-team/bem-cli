// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/bem-team/bem-cli/internal/mocktest"
)

func TestReviewQueueList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"review-queue", "list",
			"--assigned-to", "assignedTo",
			"--bucket", "bucket",
			"--cursor", "cursor",
			"--limit", "1",
			"--since", "since",
			"--status", "string",
			"--type", "string",
		)
	})
}
