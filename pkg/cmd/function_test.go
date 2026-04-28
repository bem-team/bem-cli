// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/bem-team/bem-cli/internal/mocktest"
	"github.com/bem-team/bem-cli/internal/requestflag"
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
			"--classification", "{name: name, description: description, functionID: functionID, functionName: functionName, isErrorFallback: true, origin: {email: {patterns: [string]}}, regex: {patterns: [string]}}",
			"--description", "description",
			"--destination-type", "webhook",
			"--google-drive-folder-id", "googleDriveFolderId",
			"--s3-bucket", "s3Bucket",
			"--s3-prefix", "s3Prefix",
			"--webhook-signing-enabled=true",
			"--webhook-url", "webhookUrl",
			"--print-page-split-config", "{nextFunctionID: nextFunctionID, nextFunctionName: nextFunctionName}",
			"--semantic-page-split-config", "{itemClasses: [{name: name, description: description, nextFunctionID: nextFunctionID, nextFunctionName: nextFunctionName}]}",
			"--split-type", "print_page",
			"--join-type", "standard",
			"--shaping-schema", "shapingSchema",
			"--config", "{steps: [{collectionName: collectionName, sourceField: sourceField, targetField: targetField, includeCosineDistance: true, includeSubcollections: true, scoreThreshold: 0, searchMode: semantic, topK: 1}]}",
			"--parse-config", "{extractEntities: true, linkAcrossDocuments: true, schema: {}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(functionsCreate)

		// Alternative argument passing style using inner flags
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
			"--classification.name", "name",
			"--classification.description", "description",
			"--classification.function-id", "functionID",
			"--classification.function-name", "functionName",
			"--classification.is-error-fallback=true",
			"--classification.origin", "{email: {patterns: [string]}}",
			"--classification.regex", "{patterns: [string]}",
			"--description", "description",
			"--destination-type", "webhook",
			"--google-drive-folder-id", "googleDriveFolderId",
			"--s3-bucket", "s3Bucket",
			"--s3-prefix", "s3Prefix",
			"--webhook-signing-enabled=true",
			"--webhook-url", "webhookUrl",
			"--print-page-split-config.next-function-id", "nextFunctionID",
			"--print-page-split-config.next-function-name", "nextFunctionName",
			"--semantic-page-split-config.item-classes", "[{name: name, description: description, nextFunctionID: nextFunctionID, nextFunctionName: nextFunctionName}]",
			"--split-type", "print_page",
			"--join-type", "standard",
			"--shaping-schema", "shapingSchema",
			"--config.steps", "[{collectionName: collectionName, sourceField: sourceField, targetField: targetField, includeCosineDistance: true, includeSubcollections: true, scoreThreshold: 0, searchMode: semantic, topK: 1}]",
			"--parse-config.extract-entities=true",
			"--parse-config.link-across-documents=true",
			"--parse-config.schema", "{}",
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
			"  - string\n" +
			"classifications:\n" +
			"  - name: name\n" +
			"    description: description\n" +
			"    functionID: functionID\n" +
			"    functionName: functionName\n" +
			"    isErrorFallback: true\n" +
			"    origin:\n" +
			"      email:\n" +
			"        patterns:\n" +
			"          - string\n" +
			"    regex:\n" +
			"      patterns:\n" +
			"        - string\n" +
			"description: description\n" +
			"destinationType: webhook\n" +
			"googleDriveFolderId: googleDriveFolderId\n" +
			"s3Bucket: s3Bucket\n" +
			"s3Prefix: s3Prefix\n" +
			"webhookSigningEnabled: true\n" +
			"webhookUrl: webhookUrl\n" +
			"printPageSplitConfig:\n" +
			"  nextFunctionID: nextFunctionID\n" +
			"  nextFunctionName: nextFunctionName\n" +
			"semanticPageSplitConfig:\n" +
			"  itemClasses:\n" +
			"    - name: name\n" +
			"      description: description\n" +
			"      nextFunctionID: nextFunctionID\n" +
			"      nextFunctionName: nextFunctionName\n" +
			"splitType: print_page\n" +
			"joinType: standard\n" +
			"shapingSchema: shapingSchema\n" +
			"config:\n" +
			"  steps:\n" +
			"    - collectionName: collectionName\n" +
			"      sourceField: sourceField\n" +
			"      targetField: targetField\n" +
			"      includeCosineDistance: true\n" +
			"      includeSubcollections: true\n" +
			"      scoreThreshold: 0\n" +
			"      searchMode: semantic\n" +
			"      topK: 1\n" +
			"parseConfig:\n" +
			"  extractEntities: true\n" +
			"  linkAcrossDocuments: true\n" +
			"  schema: {}\n")
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
			"--classification", "{name: name, description: description, functionID: functionID, functionName: functionName, isErrorFallback: true, origin: {email: {patterns: [string]}}, regex: {patterns: [string]}}",
			"--description", "description",
			"--destination-type", "webhook",
			"--google-drive-folder-id", "googleDriveFolderId",
			"--s3-bucket", "s3Bucket",
			"--s3-prefix", "s3Prefix",
			"--webhook-signing-enabled=true",
			"--webhook-url", "webhookUrl",
			"--print-page-split-config", "{nextFunctionID: nextFunctionID, nextFunctionName: nextFunctionName}",
			"--semantic-page-split-config", "{itemClasses: [{name: name, description: description, nextFunctionID: nextFunctionID, nextFunctionName: nextFunctionName}]}",
			"--split-type", "print_page",
			"--join-type", "standard",
			"--shaping-schema", "shapingSchema",
			"--config", "{steps: [{collectionName: collectionName, sourceField: sourceField, targetField: targetField, includeCosineDistance: true, includeSubcollections: true, scoreThreshold: 0, searchMode: semantic, topK: 1}]}",
			"--parse-config", "{extractEntities: true, linkAcrossDocuments: true, schema: {}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(functionsUpdate)

		// Alternative argument passing style using inner flags
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
			"--classification.name", "name",
			"--classification.description", "description",
			"--classification.function-id", "functionID",
			"--classification.function-name", "functionName",
			"--classification.is-error-fallback=true",
			"--classification.origin", "{email: {patterns: [string]}}",
			"--classification.regex", "{patterns: [string]}",
			"--description", "description",
			"--destination-type", "webhook",
			"--google-drive-folder-id", "googleDriveFolderId",
			"--s3-bucket", "s3Bucket",
			"--s3-prefix", "s3Prefix",
			"--webhook-signing-enabled=true",
			"--webhook-url", "webhookUrl",
			"--print-page-split-config.next-function-id", "nextFunctionID",
			"--print-page-split-config.next-function-name", "nextFunctionName",
			"--semantic-page-split-config.item-classes", "[{name: name, description: description, nextFunctionID: nextFunctionID, nextFunctionName: nextFunctionName}]",
			"--split-type", "print_page",
			"--join-type", "standard",
			"--shaping-schema", "shapingSchema",
			"--config.steps", "[{collectionName: collectionName, sourceField: sourceField, targetField: targetField, includeCosineDistance: true, includeSubcollections: true, scoreThreshold: 0, searchMode: semantic, topK: 1}]",
			"--parse-config.extract-entities=true",
			"--parse-config.link-across-documents=true",
			"--parse-config.schema", "{}",
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
			"  - string\n" +
			"classifications:\n" +
			"  - name: name\n" +
			"    description: description\n" +
			"    functionID: functionID\n" +
			"    functionName: functionName\n" +
			"    isErrorFallback: true\n" +
			"    origin:\n" +
			"      email:\n" +
			"        patterns:\n" +
			"          - string\n" +
			"    regex:\n" +
			"      patterns:\n" +
			"        - string\n" +
			"description: description\n" +
			"destinationType: webhook\n" +
			"googleDriveFolderId: googleDriveFolderId\n" +
			"s3Bucket: s3Bucket\n" +
			"s3Prefix: s3Prefix\n" +
			"webhookSigningEnabled: true\n" +
			"webhookUrl: webhookUrl\n" +
			"printPageSplitConfig:\n" +
			"  nextFunctionID: nextFunctionID\n" +
			"  nextFunctionName: nextFunctionName\n" +
			"semanticPageSplitConfig:\n" +
			"  itemClasses:\n" +
			"    - name: name\n" +
			"      description: description\n" +
			"      nextFunctionID: nextFunctionID\n" +
			"      nextFunctionName: nextFunctionName\n" +
			"splitType: print_page\n" +
			"joinType: standard\n" +
			"shapingSchema: shapingSchema\n" +
			"config:\n" +
			"  steps:\n" +
			"    - collectionName: collectionName\n" +
			"      sourceField: sourceField\n" +
			"      targetField: targetField\n" +
			"      includeCosineDistance: true\n" +
			"      includeSubcollections: true\n" +
			"      scoreThreshold: 0\n" +
			"      searchMode: semantic\n" +
			"      topK: 1\n" +
			"parseConfig:\n" +
			"  extractEntities: true\n" +
			"  linkAcrossDocuments: true\n" +
			"  schema: {}\n")
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
