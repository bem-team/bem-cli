// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/bem-team/bem-cli/internal/mocktest"
)

func TestEntityTypesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"entity-types", "create",
			"--name", "Drug",
			"--attribute-schema", "{}",
			"--description", "A pharmaceutical compound",
			"--parent-type-id", "parentTypeID",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: Drug\n" +
			"attributeSchema: {}\n" +
			"description: A pharmaceutical compound\n" +
			"parentTypeID: parentTypeID\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"entity-types", "create",
		)
	})
}

func TestEntityTypesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"entity-types", "retrieve",
			"--type-id", "typeID",
		)
	})
}

func TestEntityTypesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"entity-types", "update",
			"--type-id", "typeID",
			"--attribute-schema", "{}",
			"--description", "description",
			"--parent-type-id", "parentTypeID",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"attributeSchema: {}\n" +
			"description: description\n" +
			"parentTypeID: parentTypeID\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"entity-types", "update",
			"--type-id", "typeID",
		)
	})
}

func TestEntityTypesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"entity-types", "delete",
			"--type-id", "typeID",
		)
	})
}
