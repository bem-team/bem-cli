// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/bem-team/bem-cli/internal/mocktest"
	"github.com/bem-team/bem-cli/internal/requestflag"
)

func TestCollectionsItemsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"collections:items", "retrieve",
			"--collection-name", "collectionName",
			"--include-subcollections=true",
			"--limit", "1",
			"--page", "1",
		)
	})
}

func TestCollectionsItemsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"collections:items", "update",
			"--collection-name", "product_catalog",
			"--item", "{collectionItemID: clitm_2N6gH8ZKCmvb6BnFcGqhKJ98VzP, data: 'SKU-12345: Updated Industrial Widget - Premium Edition'}",
			"--item", "{collectionItemID: clitm_3M7hI9ALDnwc7CoGdHriLK09WaQ, data: {sku: SKU-67890, name: Updated Premium Gear, category: Hardware, price: 399.99}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(collectionsItemsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"collections:items", "update",
			"--collection-name", "product_catalog",
			"--item.collection-item-id", "clitm_2N6gH8ZKCmvb6BnFcGqhKJ98VzP",
			"--item.data", "SKU-12345: Updated Industrial Widget - Premium Edition",
			"--item.collection-item-id", "clitm_3M7hI9ALDnwc7CoGdHriLK09WaQ",
			"--item.data", "{sku: SKU-67890, name: Updated Premium Gear, category: Hardware, price: 399.99}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"collectionName: product_catalog\n" +
			"items:\n" +
			"  - collectionItemID: clitm_2N6gH8ZKCmvb6BnFcGqhKJ98VzP\n" +
			"    data: 'SKU-12345: Updated Industrial Widget - Premium Edition'\n" +
			"  - collectionItemID: clitm_3M7hI9ALDnwc7CoGdHriLK09WaQ\n" +
			"    data:\n" +
			"      sku: SKU-67890\n" +
			"      name: Updated Premium Gear\n" +
			"      category: Hardware\n" +
			"      price: 399.99\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"collections:items", "update",
		)
	})
}

func TestCollectionsItemsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"collections:items", "delete",
			"--collection-item-id", "collectionItemID",
			"--collection-name", "collectionName",
		)
	})
}

func TestCollectionsItemsAdd(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"collections:items", "add",
			"--collection-name", "product_catalog",
			"--item", "{data: {sku: SKU-11111, name: Deluxe Component, category: Hardware, price: 299.99}}",
			"--item", "{data: {sku: SKU-22222, name: Standard Part, category: Tools, price: 49.99}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(collectionsItemsAdd)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"collections:items", "add",
			"--collection-name", "product_catalog",
			"--item.data", "{sku: SKU-11111, name: Deluxe Component, category: Hardware, price: 299.99}",
			"--item.data", "{sku: SKU-22222, name: Standard Part, category: Tools, price: 49.99}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"collectionName: product_catalog\n" +
			"items:\n" +
			"  - data:\n" +
			"      sku: SKU-11111\n" +
			"      name: Deluxe Component\n" +
			"      category: Hardware\n" +
			"      price: 299.99\n" +
			"  - data:\n" +
			"      sku: SKU-22222\n" +
			"      name: Standard Part\n" +
			"      category: Tools\n" +
			"      price: 49.99\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"collections:items", "add",
		)
	})
}
