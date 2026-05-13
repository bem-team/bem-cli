// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/bem-team/bem-cli/internal/mocktest"
)

func TestFunctionsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"functions", "create",
			"--function-name", "functionName",
			"--type", "extract",
			"--display-name", "displayName",
			"--enable-bounding-boxes=true",
			"--output-schema", "{}",
			"--output-schema-name", "outputSchemaName",
			"--pre-count=true",
			"--tabular-chunking-enabled=true",
			"--tag", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"functionName: functionName\n" +
			"type: extract\n" +
			"displayName: displayName\n" +
			"enableBoundingBoxes: true\n" +
			"outputSchema: {}\n" +
			"outputSchemaName: outputSchemaName\n" +
			"preCount: true\n" +
			"tabularChunkingEnabled: true\n" +
			"tags:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"functions", "create",
		)
	})
}

func TestFunctionsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"functions", "retrieve",
			"--function-name", "functionName",
		)
	})
}

func TestFunctionsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"functions", "update",
			"--path-function-name", "functionName",
			"--type", "extract",
			"--display-name", "displayName",
			"--enable-bounding-boxes=true",
			"--function-name", "functionName",
			"--output-schema", "{}",
			"--output-schema-name", "outputSchemaName",
			"--pre-count=true",
			"--tabular-chunking-enabled=true",
			"--tag", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"type: extract\n" +
			"displayName: displayName\n" +
			"enableBoundingBoxes: true\n" +
			"functionName: functionName\n" +
			"outputSchema: {}\n" +
			"outputSchemaName: outputSchemaName\n" +
			"preCount: true\n" +
			"tabularChunkingEnabled: true\n" +
			"tags:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"functions", "update",
			"--path-function-name", "functionName",
		)
	})
}

func TestFunctionsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"functions", "list",
			"--max-items", "10",
			"--display-name", "displayName",
			"--ending-before", "endingBefore",
			"--function-id", "string",
			"--function-name", "string",
			"--limit", "1",
			"--sort-order", "asc",
			"--starting-after", "startingAfter",
			"--tag", "string",
			"--type", "transform",
			"--workflow-id", "string",
			"--workflow-name", "string",
		)
	})
}

func TestFunctionsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"functions", "delete",
			"--function-name", "functionName",
		)
	})
}
