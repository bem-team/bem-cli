// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/bem-team/bem-cli/internal/mocktest"
	"github.com/bem-team/bem-cli/internal/requestflag"
)

func TestViewsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"views", "create",
			"--aggregation", "{function: count, name: name, aggregateColumnName: aggregateColumnName, displayType: table, groupByColumnName: groupByColumnName}",
			"--column", "{displayOrderIndex: 0, name: name, valueSchemaPath: [string]}",
			"--filter", "{columnName: columnName, filterType: equals_string, number: 0, string: string}",
			"--function", "{id: id, name: name}",
			"--name", "name",
			"--description", "description",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(viewsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"views", "create",
			"--aggregation.function", "count",
			"--aggregation.name", "name",
			"--aggregation.aggregate-column-name", "aggregateColumnName",
			"--aggregation.display-type", "table",
			"--aggregation.group-by-column-name", "groupByColumnName",
			"--column.display-order-index", "0",
			"--column.name", "name",
			"--column.value-schema-path", "[string]",
			"--filter.column-name", "columnName",
			"--filter.filter-type", "equals_string",
			"--filter.number", "0",
			"--filter.string", "string",
			"--function.id", "id",
			"--function.name", "name",
			"--name", "name",
			"--description", "description",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"aggregations:\n" +
			"  - function: count\n" +
			"    name: name\n" +
			"    aggregateColumnName: aggregateColumnName\n" +
			"    displayType: table\n" +
			"    groupByColumnName: groupByColumnName\n" +
			"columns:\n" +
			"  - displayOrderIndex: 0\n" +
			"    name: name\n" +
			"    valueSchemaPath:\n" +
			"      - string\n" +
			"filters:\n" +
			"  - columnName: columnName\n" +
			"    filterType: equals_string\n" +
			"    number: 0\n" +
			"    string: string\n" +
			"functions:\n" +
			"  - id: id\n" +
			"    name: name\n" +
			"name: name\n" +
			"description: description\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"views", "create",
		)
	})
}

func TestViewsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"views", "retrieve",
			"--view-id", "view_id",
		)
	})
}

func TestViewsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"views", "update",
			"--view-id", "view_id",
			"--aggregation", "{function: count, name: name, aggregateColumnName: aggregateColumnName, displayType: table, groupByColumnName: groupByColumnName}",
			"--column", "{displayOrderIndex: 0, name: name, valueSchemaPath: [string]}",
			"--filter", "{columnName: columnName, filterType: equals_string, number: 0, string: string}",
			"--function", "{id: id, name: name}",
			"--name", "name",
			"--description", "description",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(viewsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"views", "update",
			"--view-id", "view_id",
			"--aggregation.function", "count",
			"--aggregation.name", "name",
			"--aggregation.aggregate-column-name", "aggregateColumnName",
			"--aggregation.display-type", "table",
			"--aggregation.group-by-column-name", "groupByColumnName",
			"--column.display-order-index", "0",
			"--column.name", "name",
			"--column.value-schema-path", "[string]",
			"--filter.column-name", "columnName",
			"--filter.filter-type", "equals_string",
			"--filter.number", "0",
			"--filter.string", "string",
			"--function.id", "id",
			"--function.name", "name",
			"--name", "name",
			"--description", "description",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"aggregations:\n" +
			"  - function: count\n" +
			"    name: name\n" +
			"    aggregateColumnName: aggregateColumnName\n" +
			"    displayType: table\n" +
			"    groupByColumnName: groupByColumnName\n" +
			"columns:\n" +
			"  - displayOrderIndex: 0\n" +
			"    name: name\n" +
			"    valueSchemaPath:\n" +
			"      - string\n" +
			"filters:\n" +
			"  - columnName: columnName\n" +
			"    filterType: equals_string\n" +
			"    number: 0\n" +
			"    string: string\n" +
			"functions:\n" +
			"  - id: id\n" +
			"    name: name\n" +
			"name: name\n" +
			"description: description\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"views", "update",
			"--view-id", "view_id",
		)
	})
}

func TestViewsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"views", "list",
			"--max-items", "10",
			"--ending-before", "endingBefore",
			"--function-id", "string",
			"--function-name", "string",
			"--limit", "1",
			"--sort-order", "asc",
			"--starting-after", "startingAfter",
			"--view-id", "string",
			"--view-name-substring", "viewNameSubstring",
		)
	})
}

func TestViewsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"views", "delete",
			"--view-id", "view_id",
		)
	})
}

func TestViewsGenerateAggregationData(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"views", "generate-aggregation-data",
			"--aggregation", "{function: count, name: name, aggregateColumnName: aggregateColumnName, displayType: table, groupByColumnName: groupByColumnName}",
			"--column", "{displayOrderIndex: 0, name: name, valueSchemaPath: [string]}",
			"--filter", "{columnName: columnName, filterType: equals_string, number: 0, string: string}",
			"--function", "{id: id, name: name}",
			"--name", "name",
			"--time-window", "{end: '2019-12-27T18:11:19.117Z', start: '2019-12-27T18:11:19.117Z'}",
			"--description", "description",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(viewsGenerateAggregationData)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"views", "generate-aggregation-data",
			"--aggregation.function", "count",
			"--aggregation.name", "name",
			"--aggregation.aggregate-column-name", "aggregateColumnName",
			"--aggregation.display-type", "table",
			"--aggregation.group-by-column-name", "groupByColumnName",
			"--column.display-order-index", "0",
			"--column.name", "name",
			"--column.value-schema-path", "[string]",
			"--filter.column-name", "columnName",
			"--filter.filter-type", "equals_string",
			"--filter.number", "0",
			"--filter.string", "string",
			"--function.id", "id",
			"--function.name", "name",
			"--name", "name",
			"--time-window.end", "2019-12-27T18:11:19.117Z",
			"--time-window.start", "2019-12-27T18:11:19.117Z",
			"--description", "description",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"aggregations:\n" +
			"  - function: count\n" +
			"    name: name\n" +
			"    aggregateColumnName: aggregateColumnName\n" +
			"    displayType: table\n" +
			"    groupByColumnName: groupByColumnName\n" +
			"columns:\n" +
			"  - displayOrderIndex: 0\n" +
			"    name: name\n" +
			"    valueSchemaPath:\n" +
			"      - string\n" +
			"filters:\n" +
			"  - columnName: columnName\n" +
			"    filterType: equals_string\n" +
			"    number: 0\n" +
			"    string: string\n" +
			"functions:\n" +
			"  - id: id\n" +
			"    name: name\n" +
			"name: name\n" +
			"timeWindow:\n" +
			"  end: '2019-12-27T18:11:19.117Z'\n" +
			"  start: '2019-12-27T18:11:19.117Z'\n" +
			"description: description\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"views", "generate-aggregation-data",
		)
	})
}

func TestViewsGenerateTableData(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"views", "generate-table-data",
			"--aggregation", "{function: count, name: name, aggregateColumnName: aggregateColumnName, displayType: table, groupByColumnName: groupByColumnName}",
			"--column", "{displayOrderIndex: 0, name: name, valueSchemaPath: [string]}",
			"--filter", "{columnName: columnName, filterType: equals_string, number: 0, string: string}",
			"--function", "{id: id, name: name}",
			"--name", "name",
			"--time-window", "{end: '2019-12-27T18:11:19.117Z', start: '2019-12-27T18:11:19.117Z'}",
			"--description", "description",
			"--limit", "1",
			"--offset", "0",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(viewsGenerateTableData)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"views", "generate-table-data",
			"--aggregation.function", "count",
			"--aggregation.name", "name",
			"--aggregation.aggregate-column-name", "aggregateColumnName",
			"--aggregation.display-type", "table",
			"--aggregation.group-by-column-name", "groupByColumnName",
			"--column.display-order-index", "0",
			"--column.name", "name",
			"--column.value-schema-path", "[string]",
			"--filter.column-name", "columnName",
			"--filter.filter-type", "equals_string",
			"--filter.number", "0",
			"--filter.string", "string",
			"--function.id", "id",
			"--function.name", "name",
			"--name", "name",
			"--time-window.end", "2019-12-27T18:11:19.117Z",
			"--time-window.start", "2019-12-27T18:11:19.117Z",
			"--description", "description",
			"--limit", "1",
			"--offset", "0",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"aggregations:\n" +
			"  - function: count\n" +
			"    name: name\n" +
			"    aggregateColumnName: aggregateColumnName\n" +
			"    displayType: table\n" +
			"    groupByColumnName: groupByColumnName\n" +
			"columns:\n" +
			"  - displayOrderIndex: 0\n" +
			"    name: name\n" +
			"    valueSchemaPath:\n" +
			"      - string\n" +
			"filters:\n" +
			"  - columnName: columnName\n" +
			"    filterType: equals_string\n" +
			"    number: 0\n" +
			"    string: string\n" +
			"functions:\n" +
			"  - id: id\n" +
			"    name: name\n" +
			"name: name\n" +
			"timeWindow:\n" +
			"  end: '2019-12-27T18:11:19.117Z'\n" +
			"  start: '2019-12-27T18:11:19.117Z'\n" +
			"description: description\n" +
			"limit: 1\n" +
			"offset: 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"views", "generate-table-data",
		)
	})
}
