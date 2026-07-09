// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/bem-team/bem-cli/internal/mocktest"
)

func TestEntityTypesReviewersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"entity-types:reviewers", "list",
			"--type-id", "typeID",
		)
	})
}

func TestEntityTypesReviewersAssign(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"entity-types:reviewers", "assign",
			"--type-id", "typeID",
			"--user-id", "usr_2xyz...",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("userID: usr_2xyz...")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"entity-types:reviewers", "assign",
			"--type-id", "typeID",
		)
	})
}

func TestEntityTypesReviewersRemove(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"entity-types:reviewers", "remove",
			"--type-id", "typeID",
			"--user-id", "userID",
		)
	})
}
