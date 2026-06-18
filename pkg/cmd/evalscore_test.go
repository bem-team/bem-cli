// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/bem-team/bem-cli/internal/mocktest"
	"github.com/bem-team/bem-cli/internal/requestflag"
)

func TestEvalScoreCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"eval:score", "create",
			"--function-name", "functionName",
			"--pair", "{expected: {}, input: {inputContent: inputContent, inputType: csv}}",
			"--function-version-num", "0",
			"--match-config", "{arrayMatch: by-index, fuzzyThreshold: 0, ignorePaths: [string], numericTolerance: 0, stringMatch: exact}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(evalScoreCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"eval:score", "create",
			"--function-name", "functionName",
			"--pair.expected", "{}",
			"--pair.input", "{inputContent: inputContent, inputType: csv}",
			"--function-version-num", "0",
			"--match-config.array-match", "by-index",
			"--match-config.fuzzy-threshold", "0",
			"--match-config.ignore-paths", "[string]",
			"--match-config.numeric-tolerance", "0",
			"--match-config.string-match", "exact",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"functionName: functionName\n" +
			"pairs:\n" +
			"  - expected: {}\n" +
			"    input:\n" +
			"      inputContent: inputContent\n" +
			"      inputType: csv\n" +
			"functionVersionNum: 0\n" +
			"matchConfig:\n" +
			"  arrayMatch: by-index\n" +
			"  fuzzyThreshold: 0\n" +
			"  ignorePaths:\n" +
			"    - string\n" +
			"  numericTolerance: 0\n" +
			"  stringMatch: exact\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"eval:score", "create",
		)
	})
}

func TestEvalScoreRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"eval:score", "retrieve",
			"--score-run-id", "scoreRunID",
		)
	})
}

func TestEvalScoreCancel(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"eval:score", "cancel",
			"--score-run-id", "scoreRunID",
		)
	})
}
