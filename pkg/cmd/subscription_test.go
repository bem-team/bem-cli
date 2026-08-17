// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/bem-team/bem-cli/internal/mocktest"
)

func TestSubscriptionsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"subscriptions", "create",
			"--name", "name",
			"--type", "transform",
			"--collection-id", "collectionID",
			"--collection-name", "collectionName",
			"--disabled=true",
			"--function-id", "functionID",
			"--function-name", "functionName",
			"--google-drive-folder-id", "googleDriveFolderID",
			"--s3-bucket", "s3Bucket",
			"--s3-file-path", "s3FilePath",
			"--webhook-url", "webhookURL",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: name\n" +
			"type: transform\n" +
			"collectionID: collectionID\n" +
			"collectionName: collectionName\n" +
			"disabled: true\n" +
			"functionID: functionID\n" +
			"functionName: functionName\n" +
			"googleDriveFolderID: googleDriveFolderID\n" +
			"s3Bucket: s3Bucket\n" +
			"s3FilePath: s3FilePath\n" +
			"webhookURL: webhookURL\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"subscriptions", "create",
		)
	})
}

func TestSubscriptionsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"subscriptions", "retrieve",
			"--subscription-id", "subscriptionID",
		)
	})
}

func TestSubscriptionsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"subscriptions", "update",
			"--subscription-id", "subscriptionID",
			"--disabled=true",
			"--function-name", "functionName",
			"--google-drive-folder-id", "googleDriveFolderID",
			"--name", "name",
			"--s3-bucket", "s3Bucket",
			"--s3-file-path", "s3FilePath",
			"--type", "transform",
			"--webhook-url", "webhookURL",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"disabled: true\n" +
			"functionName: functionName\n" +
			"googleDriveFolderID: googleDriveFolderID\n" +
			"name: name\n" +
			"s3Bucket: s3Bucket\n" +
			"s3FilePath: s3FilePath\n" +
			"type: transform\n" +
			"webhookURL: webhookURL\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"subscriptions", "update",
			"--subscription-id", "subscriptionID",
		)
	})
}

func TestSubscriptionsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"subscriptions", "list",
			"--ending-before", "endingBefore",
			"--function-name", "string",
			"--limit", "1",
			"--starting-after", "startingAfter",
		)
	})
}

func TestSubscriptionsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"subscriptions", "delete",
			"--subscription-id", "subscriptionID",
		)
	})
}
