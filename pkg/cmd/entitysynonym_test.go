// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/bem-team/bem-cli/internal/mocktest"
)

func TestEntitiesSynonymsAdd(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"entities:synonyms", "add",
			"--id", "id",
			"--text", "ACME Corporation",
			"--bucket", "bucket",
			"--locale", "en-US",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"text: ACME Corporation\n" +
			"locale: en-US\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"entities:synonyms", "add",
			"--id", "id",
			"--bucket", "bucket",
		)
	})
}

func TestEntitiesSynonymsRemove(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"entities:synonyms", "remove",
			"--id", "id",
			"--synonym-id", "synonymID",
			"--bucket", "bucket",
		)
	})
}
