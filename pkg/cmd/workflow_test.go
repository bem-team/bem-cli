// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/bem-team/bem-cli/internal/mocktest"
	"github.com/bem-team/bem-cli/internal/requestflag"
)

func TestWorkflowsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workflows", "create",
			"--main-node-name", "mainNodeName",
			"--name", "name",
			"--node", "{function: {id: id, name: name, versionNum: 0}, metadata: {}, name: name}",
			"--connector", "{name: name, type: paragon, connectorID: connectorID, paragon: {configuration: {}, integration: integration}}",
			"--display-name", "displayName",
			"--edge", "{destinationNodeName: destinationNodeName, sourceNodeName: sourceNodeName, destinationName: destinationName, metadata: {}}",
			"--tag", "string",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(workflowsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workflows", "create",
			"--main-node-name", "mainNodeName",
			"--name", "name",
			"--node.function", "{id: id, name: name, versionNum: 0}",
			"--node.metadata", "{}",
			"--node.name", "name",
			"--connector.name", "name",
			"--connector.type", "paragon",
			"--connector.connector-id", "connectorID",
			"--connector.paragon", "{configuration: {}, integration: integration}",
			"--display-name", "displayName",
			"--edge.destination-node-name", "destinationNodeName",
			"--edge.source-node-name", "sourceNodeName",
			"--edge.destination-name", "destinationName",
			"--edge.metadata", "{}",
			"--tag", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"mainNodeName: mainNodeName\n" +
			"name: name\n" +
			"nodes:\n" +
			"  - function:\n" +
			"      id: id\n" +
			"      name: name\n" +
			"      versionNum: 0\n" +
			"    metadata: {}\n" +
			"    name: name\n" +
			"connectors:\n" +
			"  - name: name\n" +
			"    type: paragon\n" +
			"    connectorID: connectorID\n" +
			"    paragon:\n" +
			"      configuration: {}\n" +
			"      integration: integration\n" +
			"displayName: displayName\n" +
			"edges:\n" +
			"  - destinationNodeName: destinationNodeName\n" +
			"    sourceNodeName: sourceNodeName\n" +
			"    destinationName: destinationName\n" +
			"    metadata: {}\n" +
			"tags:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"workflows", "create",
		)
	})
}

func TestWorkflowsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workflows", "retrieve",
			"--workflow-name", "workflowName",
		)
	})
}

func TestWorkflowsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workflows", "update",
			"--workflow-name", "workflowName",
			"--connector", "{name: name, type: paragon, connectorID: connectorID, paragon: {configuration: {}, integration: integration}}",
			"--display-name", "displayName",
			"--edge", "{destinationNodeName: destinationNodeName, sourceNodeName: sourceNodeName, destinationName: destinationName, metadata: {}}",
			"--main-node-name", "mainNodeName",
			"--name", "name",
			"--node", "{function: {id: id, name: name, versionNum: 0}, metadata: {}, name: name}",
			"--tag", "string",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(workflowsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workflows", "update",
			"--workflow-name", "workflowName",
			"--connector.name", "name",
			"--connector.type", "paragon",
			"--connector.connector-id", "connectorID",
			"--connector.paragon", "{configuration: {}, integration: integration}",
			"--display-name", "displayName",
			"--edge.destination-node-name", "destinationNodeName",
			"--edge.source-node-name", "sourceNodeName",
			"--edge.destination-name", "destinationName",
			"--edge.metadata", "{}",
			"--main-node-name", "mainNodeName",
			"--name", "name",
			"--node.function", "{id: id, name: name, versionNum: 0}",
			"--node.metadata", "{}",
			"--node.name", "name",
			"--tag", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"connectors:\n" +
			"  - name: name\n" +
			"    type: paragon\n" +
			"    connectorID: connectorID\n" +
			"    paragon:\n" +
			"      configuration: {}\n" +
			"      integration: integration\n" +
			"displayName: displayName\n" +
			"edges:\n" +
			"  - destinationNodeName: destinationNodeName\n" +
			"    sourceNodeName: sourceNodeName\n" +
			"    destinationName: destinationName\n" +
			"    metadata: {}\n" +
			"mainNodeName: mainNodeName\n" +
			"name: name\n" +
			"nodes:\n" +
			"  - function:\n" +
			"      id: id\n" +
			"      name: name\n" +
			"      versionNum: 0\n" +
			"    metadata: {}\n" +
			"    name: name\n" +
			"tags:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"workflows", "update",
			"--workflow-name", "workflowName",
		)
	})
}

func TestWorkflowsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workflows", "list",
			"--max-items", "10",
			"--display-name", "displayName",
			"--ending-before", "endingBefore",
			"--function-id", "string",
			"--function-name", "string",
			"--limit", "1",
			"--sort-order", "asc",
			"--starting-after", "startingAfter",
			"--tag", "string",
			"--workflow-id", "string",
			"--workflow-name", "string",
		)
	})
}

func TestWorkflowsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workflows", "delete",
			"--workflow-name", "workflowName",
		)
	})
}

func TestWorkflowsCall(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workflows", "call",
			"--workflow-name", "workflowName",
			"--input", "{batchFiles: {inputs: [{inputContent: inputContent, inputType: csv, itemReferenceID: itemReferenceID}]}, singleFile: {inputContent: inputContent, inputType: csv}}",
			"--wait=true",
			"--call-reference-id", "callReferenceID",
			"--metadata", "{}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(workflowsCall)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workflows", "call",
			"--workflow-name", "workflowName",
			"--input.batch-files", "{inputs: [{inputContent: inputContent, inputType: csv, itemReferenceID: itemReferenceID}]}",
			"--input.single-file", "{inputContent: inputContent, inputType: csv}",
			"--wait=true",
			"--call-reference-id", "callReferenceID",
			"--metadata", "{}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"input:\n" +
			"  batchFiles:\n" +
			"    inputs:\n" +
			"      - inputContent: inputContent\n" +
			"        inputType: csv\n" +
			"        itemReferenceID: itemReferenceID\n" +
			"  singleFile:\n" +
			"    inputContent: inputContent\n" +
			"    inputType: csv\n" +
			"callReferenceID: callReferenceID\n" +
			"metadata: {}\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"workflows", "call",
			"--workflow-name", "workflowName",
			"--wait=true",
		)
	})
}

func TestWorkflowsCopy(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workflows", "copy",
			"--source-workflow-name", "sourceWorkflowName",
			"--target-workflow-name", "targetWorkflowName",
			"--source-workflow-version-num", "1",
			"--tag", "string",
			"--target-display-name", "targetDisplayName",
			"--target-environment", "targetEnvironment",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"sourceWorkflowName: sourceWorkflowName\n" +
			"targetWorkflowName: targetWorkflowName\n" +
			"sourceWorkflowVersionNum: 1\n" +
			"tags:\n" +
			"  - string\n" +
			"targetDisplayName: targetDisplayName\n" +
			"targetEnvironment: targetEnvironment\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"workflows", "copy",
		)
	})
}
