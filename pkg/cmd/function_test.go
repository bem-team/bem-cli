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
			"--type", "transform",
			"--display-name", "displayName",
			"--output-schema", "{}",
			"--output-schema-name", "outputSchemaName",
			"--tabular-chunking-enabled=true",
			"--tag", "string",
			"--enable-bounding-boxes=true",
			"--pre-count=true",
			"--description", "description",
			"--route", "{name: name, description: description, functionID: functionID, functionName: functionName, isErrorFallback: true, origin: {email: {patterns: [string]}}, regex: {patterns: [string]}}",
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
			"--type", "transform",
			"--display-name", "displayName",
			"--output-schema", "{}",
			"--output-schema-name", "outputSchemaName",
			"--tabular-chunking-enabled=true",
			"--tag", "string",
			"--enable-bounding-boxes=true",
			"--pre-count=true",
			"--description", "description",
			"--route.name", "name",
			"--route.description", "description",
			"--route.function-id", "functionID",
			"--route.function-name", "functionName",
			"--route.is-error-fallback=true",
			"--route.origin", "{email: {patterns: [string]}}",
			"--route.regex", "{patterns: [string]}",
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
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"functionName: functionName\n" +
			"type: transform\n" +
			"displayName: displayName\n" +
			"outputSchema: {}\n" +
			"outputSchemaName: outputSchemaName\n" +
			"tabularChunkingEnabled: true\n" +
			"tags:\n" +
			"  - string\n" +
			"enableBoundingBoxes: true\n" +
			"preCount: true\n" +
			"description: description\n" +
			"routes:\n" +
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
			"      topK: 1\n")
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
			"--type", "transform",
			"--display-name", "displayName",
			"--function-name", "functionName",
			"--output-schema", "{}",
			"--output-schema-name", "outputSchemaName",
			"--tabular-chunking-enabled=true",
			"--tag", "string",
			"--enable-bounding-boxes=true",
			"--pre-count=true",
			"--description", "description",
			"--route", "{name: name, description: description, functionID: functionID, functionName: functionName, isErrorFallback: true, origin: {email: {patterns: [string]}}, regex: {patterns: [string]}}",
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
			"--type", "transform",
			"--display-name", "displayName",
			"--function-name", "functionName",
			"--output-schema", "{}",
			"--output-schema-name", "outputSchemaName",
			"--tabular-chunking-enabled=true",
			"--tag", "string",
			"--enable-bounding-boxes=true",
			"--pre-count=true",
			"--description", "description",
			"--route.name", "name",
			"--route.description", "description",
			"--route.function-id", "functionID",
			"--route.function-name", "functionName",
			"--route.is-error-fallback=true",
			"--route.origin", "{email: {patterns: [string]}}",
			"--route.regex", "{patterns: [string]}",
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
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"type: transform\n" +
			"displayName: displayName\n" +
			"functionName: functionName\n" +
			"outputSchema: {}\n" +
			"outputSchemaName: outputSchemaName\n" +
			"tabularChunkingEnabled: true\n" +
			"tags:\n" +
			"  - string\n" +
			"enableBoundingBoxes: true\n" +
			"preCount: true\n" +
			"description: description\n" +
			"routes:\n" +
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
			"      topK: 1\n")
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
