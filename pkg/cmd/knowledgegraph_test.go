// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/bem-team/bem-cli/internal/mocktest"
)

func TestKnowledgeGraphRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"knowledge-graph", "retrieve",
			"--bucket", "bucket",
			"--cursor", "cursor",
			"--limit", "0",
			"--max-depth", "0",
			"--node-id", "nodeID",
			"--search", "search",
			"--since", "'2019-12-27T18:11:19.117Z'",
			"--type", "string",
		)
	})
}
