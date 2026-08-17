// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/bem-team/bem-cli/internal/mocktest"
	"github.com/bem-team/bem-cli/internal/requestflag"
)

func TestEntitiesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"entities", "update",
			"--id", "id",
			"--bucket", "bucket",
			"--add-synonym", "string",
			"--assigned-type-id", "assignedTypeID",
			"--canonical", "canonical",
			"--locale", "locale",
			"--remove-synonym-id", "string",
			"--status", "approved",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"addSynonyms:\n" +
			"  - string\n" +
			"assignedTypeID: assignedTypeID\n" +
			"canonical: canonical\n" +
			"locale: locale\n" +
			"removeSynonymIDs:\n" +
			"  - string\n" +
			"status: approved\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"entities", "update",
			"--id", "id",
			"--bucket", "bucket",
		)
	})
}

func TestEntitiesBulkCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"entities", "bulk-create",
			"--entity", "{canonical: Acme Corporation, type: organization, attributes: {headquarters: Springfield}, description: Industrial conglomerate, synonyms: [ACME, Acme Corp]}",
			"--bucket", "bucket",
			"--on-conflict", "merge",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(entitiesBulkCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"entities", "bulk-create",
			"--entity.canonical", "Acme Corporation",
			"--entity.type", "organization",
			"--entity.attributes", "{headquarters: Springfield}",
			"--entity.description", "Industrial conglomerate",
			"--entity.synonyms", "[ACME, Acme Corp]",
			"--bucket", "bucket",
			"--on-conflict", "merge",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"entities:\n" +
			"  - canonical: Acme Corporation\n" +
			"    type: organization\n" +
			"    attributes:\n" +
			"      headquarters: Springfield\n" +
			"    description: Industrial conglomerate\n" +
			"    synonyms:\n" +
			"      - ACME\n" +
			"      - Acme Corp\n" +
			"bucket: bucket\n" +
			"onConflict: merge\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"entities", "bulk-create",
		)
	})
}

func TestEntitiesBulkValidate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"entities", "bulk-validate",
			"--entity-id", "ent_2abc",
			"--entity-id", "ent_2def",
			"--status", "approved",
			"--bucket", "bucket",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"entityIDs:\n" +
			"  - ent_2abc\n" +
			"  - ent_2def\n" +
			"status: approved\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"entities", "bulk-validate",
			"--bucket", "bucket",
		)
	})
}

func TestEntitiesRetrieveRelations(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"entities", "retrieve-relations",
			"--id", "id",
			"--bucket", "bucket",
			"--cursor", "cursor",
			"--direction", "inbound",
			"--limit", "0",
			"--relation-type", "relationType",
		)
	})
}

func TestEntitiesRetrieveSeedStatus(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"entities", "retrieve-seed-status",
			"--id", "id",
		)
	})
}
