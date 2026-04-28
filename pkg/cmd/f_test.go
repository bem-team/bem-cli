// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/bem-team/bem-cli/internal/mocktest"
	"github.com/bem-team/bem-cli/internal/requestflag"
)

func TestFsNavigate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"fs", "navigate",
			"--op", "ls",
			"--count-only=true",
			"--cursor", "cursor",
			"--filter", "{functionName: functionName, search: search, since: '2019-12-27T18:11:19.117Z', type: type}",
			"--ignore-case=true",
			"--limit", "0",
			"--n", "0",
			"--path", "path",
			"--pattern", "pattern",
			"--range", "{page: 0, pageRange: [0, 0], sectionTypes: [string]}",
			"--regex=true",
			"--scope", "scope",
			"--select", "string",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(fsNavigate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"fs", "navigate",
			"--op", "ls",
			"--count-only=true",
			"--cursor", "cursor",
			"--filter.function-name", "functionName",
			"--filter.search", "search",
			"--filter.since", "2019-12-27T18:11:19.117Z",
			"--filter.type", "type",
			"--ignore-case=true",
			"--limit", "0",
			"--n", "0",
			"--path", "path",
			"--pattern", "pattern",
			"--range.page", "0",
			"--range.page-range", "[0, 0]",
			"--range.section-types", "[string]",
			"--regex=true",
			"--scope", "scope",
			"--select", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"op: ls\n" +
			"countOnly: true\n" +
			"cursor: cursor\n" +
			"filter:\n" +
			"  functionName: functionName\n" +
			"  search: search\n" +
			"  since: '2019-12-27T18:11:19.117Z'\n" +
			"  type: type\n" +
			"ignoreCase: true\n" +
			"limit: 0\n" +
			"'n': 0\n" +
			"path: path\n" +
			"pattern: pattern\n" +
			"range:\n" +
			"  page: 0\n" +
			"  pageRange:\n" +
			"    - 0\n" +
			"    - 0\n" +
			"  sectionTypes:\n" +
			"    - string\n" +
			"regex: true\n" +
			"scope: scope\n" +
			"select:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"fs", "navigate",
		)
	})
}
