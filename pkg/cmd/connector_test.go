// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/bem-team/bem-cli/internal/mocktest"
)

func TestConnectorsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"connectors", "create",
			"--name", "Box → Invoice workflow",
			"--type", "paragon",
			"--box-client-id", "boxClientID",
			"--box-client-secret", "boxClientSecret",
			"--box-enterprise-id", "boxEnterpriseID",
			"--box-folder-id", "boxFolderID",
			"--paragon-configuration", "{folderId: YOUR_GOOGLE_DRIVE_FOLDER_ID}",
			"--paragon-integration", "googledrive",
			"--workflow-id", "wf_2N6gH8ZKCmvb6BnFcGqhKJ98VzP",
			"--workflow-name", "workflowName",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: Box → Invoice workflow\n" +
			"type: paragon\n" +
			"boxClientID: boxClientID\n" +
			"boxClientSecret: boxClientSecret\n" +
			"boxEnterpriseID: boxEnterpriseID\n" +
			"boxFolderID: boxFolderID\n" +
			"paragonConfiguration:\n" +
			"  folderId: YOUR_GOOGLE_DRIVE_FOLDER_ID\n" +
			"paragonIntegration: googledrive\n" +
			"workflowID: wf_2N6gH8ZKCmvb6BnFcGqhKJ98VzP\n" +
			"workflowName: workflowName\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"connectors", "create",
		)
	})
}

func TestConnectorsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"connectors", "list",
			"--workflow-id", "workflowID",
			"--workflow-name", "workflowName",
		)
	})
}
