// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/bem-team/bem-cli/internal/mocktest"
)

func TestFunctionsRegressionApplyCorrections(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"functions:regression", "apply-corrections",
			"--baseline-version-num", "3",
			"--comparison-version-num", "4",
			"--function-name", "invoice-extractor",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"baselineVersionNum: 3\n" +
			"comparisonVersionNum: 4\n" +
			"functionName: invoice-extractor\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"functions:regression", "apply-corrections",
		)
	})
}

func TestFunctionsRegressionRun(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"functions:regression", "run",
			"--function-name", "invoice-extractor",
			"--baseline-version-num", "3",
			"--comparison-version-num", "5",
			"--only-corrected-data=true",
			"--sample-size", "100",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"functionName: invoice-extractor\n" +
			"baselineVersionNum: 3\n" +
			"comparisonVersionNum: 5\n" +
			"onlyCorrectedData: true\n" +
			"sampleSize: 100\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"functions:regression", "run",
		)
	})
}
