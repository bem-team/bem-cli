// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/bem-team/bem-cli/internal/apiquery"
	"github.com/bem-team/bem-cli/internal/requestflag"
	"github.com/bem-team/bem-go-sdk"
	"github.com/bem-team/bem-go-sdk/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var viewsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "**Create a view.**",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "aggregation",
			Usage:    "List of aggregations defined for the view",
			Required: true,
			BodyPath: "aggregations",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "column",
			Usage:    "List of columns in the view",
			Required: true,
			BodyPath: "columns",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "filter",
			Usage:    "List of filters applied to the view",
			Required: true,
			BodyPath: "filters",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "function",
			Usage:    "List of functions that this view queries transformations from",
			Required: true,
			BodyPath: "functions",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the view",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "Description of the view",
			BodyPath: "description",
		},
	},
	Action:          handleViewsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"aggregation": {
		&requestflag.InnerFlag[string]{
			Name:       "aggregation.function",
			Usage:      "Aggregation function to apply to a view column",
			InnerField: "function",
		},
		&requestflag.InnerFlag[string]{
			Name:       "aggregation.name",
			Usage:      "Name of the aggregation",
			InnerField: "name",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "aggregation.aggregate-column-name",
			Usage:      "Name of the column to aggregate (required for count_distinct, sum, average, min, max functions)",
			InnerField: "aggregateColumnName",
		},
		&requestflag.InnerFlag[string]{
			Name:       "aggregation.display-type",
			Usage:      "How to display the aggregation results",
			InnerField: "displayType",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "aggregation.group-by-column-name",
			Usage:      "Name of the column to group by (optional, for grouped aggregations)",
			InnerField: "groupByColumnName",
		},
	},
	"column": {
		&requestflag.InnerFlag[int64]{
			Name:       "column.display-order-index",
			Usage:      "Order in which this column should be displayed (0-based index)",
			InnerField: "displayOrderIndex",
		},
		&requestflag.InnerFlag[string]{
			Name:       "column.name",
			Usage:      "Name of the column",
			InnerField: "name",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "column.value-schema-path",
			Usage:      `JSON path to the value in the transformation output schema (e.g., ["invoiceDetails", "invoiceNumber"])`,
			InnerField: "valueSchemaPath",
		},
	},
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.column-name",
			Usage:      "Name of the column to filter on",
			InnerField: "columnName",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.filter-type",
			Usage:      "Type of filter to apply to a view column",
			InnerField: "filterType",
		},
		&requestflag.InnerFlag[*float64]{
			Name:       "filter.number",
			Usage:      "Numeric value for the filter (required for number filter types)",
			InnerField: "number",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "filter.string",
			Usage:      "String value for the filter (required for string filter types)",
			InnerField: "string",
		},
	},
	"function": {
		&requestflag.InnerFlag[string]{
			Name:       "function.id",
			Usage:      "Unique identifier of function. Provide either id or name, not both.",
			InnerField: "id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "function.name",
			Usage:      "Name of function. Must be UNIQUE on a per-environment basis. Provide either id or name, not both.",
			InnerField: "name",
		},
	},
})

var viewsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "**Retrieve a view by ID.**",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "view-id",
			Required:  true,
			PathParam: "view_id",
		},
	},
	Action:          handleViewsRetrieve,
	HideHelpCommand: true,
}

var viewsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "**Update a view. Updates create a new version.**",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "view-id",
			Required:  true,
			PathParam: "view_id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "aggregation",
			Usage:    "List of aggregations defined for the view",
			Required: true,
			BodyPath: "aggregations",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "column",
			Usage:    "List of columns in the view",
			Required: true,
			BodyPath: "columns",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "filter",
			Usage:    "List of filters applied to the view",
			Required: true,
			BodyPath: "filters",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "function",
			Usage:    "List of functions that this view queries transformations from",
			Required: true,
			BodyPath: "functions",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the view",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "Description of the view",
			BodyPath: "description",
		},
	},
	Action:          handleViewsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"aggregation": {
		&requestflag.InnerFlag[string]{
			Name:       "aggregation.function",
			Usage:      "Aggregation function to apply to a view column",
			InnerField: "function",
		},
		&requestflag.InnerFlag[string]{
			Name:       "aggregation.name",
			Usage:      "Name of the aggregation",
			InnerField: "name",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "aggregation.aggregate-column-name",
			Usage:      "Name of the column to aggregate (required for count_distinct, sum, average, min, max functions)",
			InnerField: "aggregateColumnName",
		},
		&requestflag.InnerFlag[string]{
			Name:       "aggregation.display-type",
			Usage:      "How to display the aggregation results",
			InnerField: "displayType",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "aggregation.group-by-column-name",
			Usage:      "Name of the column to group by (optional, for grouped aggregations)",
			InnerField: "groupByColumnName",
		},
	},
	"column": {
		&requestflag.InnerFlag[int64]{
			Name:       "column.display-order-index",
			Usage:      "Order in which this column should be displayed (0-based index)",
			InnerField: "displayOrderIndex",
		},
		&requestflag.InnerFlag[string]{
			Name:       "column.name",
			Usage:      "Name of the column",
			InnerField: "name",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "column.value-schema-path",
			Usage:      `JSON path to the value in the transformation output schema (e.g., ["invoiceDetails", "invoiceNumber"])`,
			InnerField: "valueSchemaPath",
		},
	},
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.column-name",
			Usage:      "Name of the column to filter on",
			InnerField: "columnName",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.filter-type",
			Usage:      "Type of filter to apply to a view column",
			InnerField: "filterType",
		},
		&requestflag.InnerFlag[*float64]{
			Name:       "filter.number",
			Usage:      "Numeric value for the filter (required for number filter types)",
			InnerField: "number",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "filter.string",
			Usage:      "String value for the filter (required for string filter types)",
			InnerField: "string",
		},
	},
	"function": {
		&requestflag.InnerFlag[string]{
			Name:       "function.id",
			Usage:      "Unique identifier of function. Provide either id or name, not both.",
			InnerField: "id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "function.name",
			Usage:      "Name of function. Must be UNIQUE on a per-environment basis. Provide either id or name, not both.",
			InnerField: "name",
		},
	},
})

var viewsList = cli.Command{
	Name:    "list",
	Usage:   "**List views in the current environment, optionally filtered by the functions\nthey read from.**",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "ending-before",
			Usage:     "Cursor — a `viewID` defining your place in the list.",
			QueryPath: "endingBefore",
		},
		&requestflag.Flag[[]string]{
			Name:      "function-id",
			Usage:     "Return only views that read from at least one of the named functions.",
			QueryPath: "functionIDs",
		},
		&requestflag.Flag[[]string]{
			Name:      "function-name",
			Usage:     "Return only views that read from at least one of the named functions.",
			QueryPath: "functionNames",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Default:   50,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "sort-order",
			Usage:     "Sort order over view IDs (default `asc`).",
			Default:   "asc",
			QueryPath: "sortOrder",
		},
		&requestflag.Flag[string]{
			Name:      "starting-after",
			Usage:     "Cursor — a `viewID` defining your place in the list.",
			QueryPath: "startingAfter",
		},
		&requestflag.Flag[[]string]{
			Name:      "view-id",
			Usage:     "Return only the specified view IDs.",
			QueryPath: "viewIDs",
		},
		&requestflag.Flag[string]{
			Name:      "view-name-substring",
			Usage:     "Case-insensitive substring search over view names.",
			QueryPath: "viewNameSubstring",
		},
	},
	Action:          handleViewsList,
	HideHelpCommand: true,
}

var viewsDelete = cli.Command{
	Name:    "delete",
	Usage:   "**Delete a view and every one of its versions.**",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "view-id",
			Required:  true,
			PathParam: "view_id",
		},
	},
	Action:          handleViewsDelete,
	HideHelpCommand: true,
}

var viewsGenerateAggregationData = requestflag.WithInnerFlags(cli.Command{
	Name:    "generate-aggregation-data",
	Usage:   "**Generate aggregation results for a view.**",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "aggregation",
			Usage:    "List of aggregations defined for the view",
			Required: true,
			BodyPath: "aggregations",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "column",
			Usage:    "List of columns in the view",
			Required: true,
			BodyPath: "columns",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "filter",
			Usage:    "List of filters applied to the view",
			Required: true,
			BodyPath: "filters",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "function",
			Usage:    "List of functions that this view queries transformations from",
			Required: true,
			BodyPath: "functions",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the view",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "time-window",
			Usage:    "Time window for filtering transformations in a view",
			Required: true,
			BodyPath: "timeWindow",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "Description of the view",
			BodyPath: "description",
		},
	},
	Action:          handleViewsGenerateAggregationData,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"aggregation": {
		&requestflag.InnerFlag[string]{
			Name:       "aggregation.function",
			Usage:      "Aggregation function to apply to a view column",
			InnerField: "function",
		},
		&requestflag.InnerFlag[string]{
			Name:       "aggregation.name",
			Usage:      "Name of the aggregation",
			InnerField: "name",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "aggregation.aggregate-column-name",
			Usage:      "Name of the column to aggregate (required for count_distinct, sum, average, min, max functions)",
			InnerField: "aggregateColumnName",
		},
		&requestflag.InnerFlag[string]{
			Name:       "aggregation.display-type",
			Usage:      "How to display the aggregation results",
			InnerField: "displayType",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "aggregation.group-by-column-name",
			Usage:      "Name of the column to group by (optional, for grouped aggregations)",
			InnerField: "groupByColumnName",
		},
	},
	"column": {
		&requestflag.InnerFlag[int64]{
			Name:       "column.display-order-index",
			Usage:      "Order in which this column should be displayed (0-based index)",
			InnerField: "displayOrderIndex",
		},
		&requestflag.InnerFlag[string]{
			Name:       "column.name",
			Usage:      "Name of the column",
			InnerField: "name",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "column.value-schema-path",
			Usage:      `JSON path to the value in the transformation output schema (e.g., ["invoiceDetails", "invoiceNumber"])`,
			InnerField: "valueSchemaPath",
		},
	},
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.column-name",
			Usage:      "Name of the column to filter on",
			InnerField: "columnName",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.filter-type",
			Usage:      "Type of filter to apply to a view column",
			InnerField: "filterType",
		},
		&requestflag.InnerFlag[*float64]{
			Name:       "filter.number",
			Usage:      "Numeric value for the filter (required for number filter types)",
			InnerField: "number",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "filter.string",
			Usage:      "String value for the filter (required for string filter types)",
			InnerField: "string",
		},
	},
	"function": {
		&requestflag.InnerFlag[string]{
			Name:       "function.id",
			Usage:      "Unique identifier of function. Provide either id or name, not both.",
			InnerField: "id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "function.name",
			Usage:      "Name of function. Must be UNIQUE on a per-environment basis. Provide either id or name, not both.",
			InnerField: "name",
		},
	},
	"time-window": {
		&requestflag.InnerFlag[any]{
			Name:       "time-window.end",
			Usage:      "End of the time window in ISO 8601 (RFC 3339) format in UTC",
			InnerField: "end",
		},
		&requestflag.InnerFlag[any]{
			Name:       "time-window.start",
			Usage:      "Start of the time window in ISO 8601 (RFC 3339) format in UTC",
			InnerField: "start",
		},
	},
})

var viewsGenerateTableData = requestflag.WithInnerFlags(cli.Command{
	Name:    "generate-table-data",
	Usage:   "**Generate paginated table data for a view.**",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "aggregation",
			Usage:    "List of aggregations defined for the view",
			Required: true,
			BodyPath: "aggregations",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "column",
			Usage:    "List of columns in the view",
			Required: true,
			BodyPath: "columns",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "filter",
			Usage:    "List of filters applied to the view",
			Required: true,
			BodyPath: "filters",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "function",
			Usage:    "List of functions that this view queries transformations from",
			Required: true,
			BodyPath: "functions",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the view",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "time-window",
			Usage:    "Time window for filtering transformations in a view",
			Required: true,
			BodyPath: "timeWindow",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "Description of the view",
			BodyPath: "description",
		},
		&requestflag.Flag[*int64]{
			Name:     "limit",
			Usage:    "Maximum number of rows to return (default: 50, max: 200)",
			BodyPath: "limit",
		},
		&requestflag.Flag[*int64]{
			Name:     "offset",
			Usage:    "Number of rows to skip for pagination",
			BodyPath: "offset",
		},
	},
	Action:          handleViewsGenerateTableData,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"aggregation": {
		&requestflag.InnerFlag[string]{
			Name:       "aggregation.function",
			Usage:      "Aggregation function to apply to a view column",
			InnerField: "function",
		},
		&requestflag.InnerFlag[string]{
			Name:       "aggregation.name",
			Usage:      "Name of the aggregation",
			InnerField: "name",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "aggregation.aggregate-column-name",
			Usage:      "Name of the column to aggregate (required for count_distinct, sum, average, min, max functions)",
			InnerField: "aggregateColumnName",
		},
		&requestflag.InnerFlag[string]{
			Name:       "aggregation.display-type",
			Usage:      "How to display the aggregation results",
			InnerField: "displayType",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "aggregation.group-by-column-name",
			Usage:      "Name of the column to group by (optional, for grouped aggregations)",
			InnerField: "groupByColumnName",
		},
	},
	"column": {
		&requestflag.InnerFlag[int64]{
			Name:       "column.display-order-index",
			Usage:      "Order in which this column should be displayed (0-based index)",
			InnerField: "displayOrderIndex",
		},
		&requestflag.InnerFlag[string]{
			Name:       "column.name",
			Usage:      "Name of the column",
			InnerField: "name",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "column.value-schema-path",
			Usage:      `JSON path to the value in the transformation output schema (e.g., ["invoiceDetails", "invoiceNumber"])`,
			InnerField: "valueSchemaPath",
		},
	},
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.column-name",
			Usage:      "Name of the column to filter on",
			InnerField: "columnName",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.filter-type",
			Usage:      "Type of filter to apply to a view column",
			InnerField: "filterType",
		},
		&requestflag.InnerFlag[*float64]{
			Name:       "filter.number",
			Usage:      "Numeric value for the filter (required for number filter types)",
			InnerField: "number",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "filter.string",
			Usage:      "String value for the filter (required for string filter types)",
			InnerField: "string",
		},
	},
	"function": {
		&requestflag.InnerFlag[string]{
			Name:       "function.id",
			Usage:      "Unique identifier of function. Provide either id or name, not both.",
			InnerField: "id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "function.name",
			Usage:      "Name of function. Must be UNIQUE on a per-environment basis. Provide either id or name, not both.",
			InnerField: "name",
		},
	},
	"time-window": {
		&requestflag.InnerFlag[any]{
			Name:       "time-window.end",
			Usage:      "End of the time window in ISO 8601 (RFC 3339) format in UTC",
			InnerField: "end",
		},
		&requestflag.InnerFlag[any]{
			Name:       "time-window.start",
			Usage:      "Start of the time window in ISO 8601 (RFC 3339) format in UTC",
			InnerField: "start",
		},
	},
})

func handleViewsCreate(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := bem.ViewNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Views.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "views create",
		Transform:      transform,
	})
}

func handleViewsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("view-id") && len(unusedArgs) > 0 {
		cmd.Set("view-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Views.Get(ctx, cmd.Value("view-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "views retrieve",
		Transform:      transform,
	})
}

func handleViewsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("view-id") && len(unusedArgs) > 0 {
		cmd.Set("view-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := bem.ViewUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Views.Update(
		ctx,
		cmd.Value("view-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "views update",
		Transform:      transform,
	})
}

func handleViewsList(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := bem.ViewListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Views.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "views list",
		Transform:      transform,
	})
}

func handleViewsDelete(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("view-id") && len(unusedArgs) > 0 {
		cmd.Set("view-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	return client.Views.Delete(ctx, cmd.Value("view-id").(string), options...)
}

func handleViewsGenerateAggregationData(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := bem.ViewGenerateAggregationDataParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Views.GenerateAggregationData(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "views generate-aggregation-data",
		Transform:      transform,
	})
}

func handleViewsGenerateTableData(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := bem.ViewGenerateTableDataParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Views.GenerateTableData(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "views generate-table-data",
		Transform:      transform,
	})
}
