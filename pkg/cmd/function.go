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

var functionsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a Function",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "function-name",
			Usage:    "Name of function. Must be UNIQUE on a per-environment basis.",
			Required: true,
			BodyPath: "functionName",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    `Allowed values: "transform".`,
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[string]{
			Name:     "display-name",
			Usage:    "Display name of function. Human-readable name to help you identify the function.",
			BodyPath: "displayName",
		},
		&requestflag.Flag[any]{
			Name:     "output-schema",
			Usage:    "Desired output structure defined in standard JSON Schema convention.",
			BodyPath: "outputSchema",
		},
		&requestflag.Flag[string]{
			Name:     "output-schema-name",
			Usage:    "Name of output schema object.",
			BodyPath: "outputSchemaName",
		},
		&requestflag.Flag[bool]{
			Name:     "tabular-chunking-enabled",
			Usage:    "Whether tabular chunking is enabled on the pipeline. This processes tables in CSV/Excel in row batches, rather than all rows at once.",
			BodyPath: "tabularChunkingEnabled",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Array of tags to categorize and organize functions.",
			BodyPath: "tags",
		},
		&requestflag.Flag[bool]{
			Name:     "enable-bounding-boxes",
			Usage:    "Whether bounding box extraction is enabled. Only applicable to analyze and extract functions.\nWhen true, the function returns the document regions (page, coordinates) from which each\nfield was extracted. Enabling this automatically configures the function to use the bounding\nbox model. Disabling resets to the default.",
			BodyPath: "enableBoundingBoxes",
		},
		&requestflag.Flag[bool]{
			Name:     "pre-count",
			Usage:    "Reducing the risk of the model stopping early on long documents.\nTrade-off: Increases total latency. Compatible with\n`enableBoundingBoxes`.",
			BodyPath: "preCount",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "Description of router. Can be used to provide additional context on router's purpose and expected inputs.",
			BodyPath: "description",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "route",
			Usage:    "List of routes.",
			BodyPath: "routes",
		},
		&requestflag.Flag[string]{
			Name:     "destination-type",
			Usage:    "Destination type for a Send function.",
			BodyPath: "destinationType",
		},
		&requestflag.Flag[string]{
			Name:     "google-drive-folder-id",
			Usage:    "Google Drive folder ID. Required when destinationType is google_drive. Managed via Paragon OAuth.",
			BodyPath: "googleDriveFolderId",
		},
		&requestflag.Flag[string]{
			Name:     "s3-bucket",
			Usage:    "S3 bucket to upload the payload to. Required when destinationType is s3.",
			BodyPath: "s3Bucket",
		},
		&requestflag.Flag[string]{
			Name:     "s3-prefix",
			Usage:    "Optional S3 key prefix (folder path).",
			BodyPath: "s3Prefix",
		},
		&requestflag.Flag[bool]{
			Name:     "webhook-signing-enabled",
			Usage:    "Whether to sign webhook deliveries with an HMAC-SHA256 `bem-signature` header.\nDefaults to `true` when omitted — signing is on by default for new send functions.\nSet explicitly to `false` to disable.",
			BodyPath: "webhookSigningEnabled",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-url",
			Usage:    "Webhook URL to POST the payload to. Required when destinationType is webhook.",
			BodyPath: "webhookUrl",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "print-page-split-config",
			BodyPath: "printPageSplitConfig",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "semantic-page-split-config",
			BodyPath: "semanticPageSplitConfig",
		},
		&requestflag.Flag[string]{
			Name:     "split-type",
			Usage:    `Allowed values: "print_page", "semantic_page".`,
			BodyPath: "splitType",
		},
		&requestflag.Flag[string]{
			Name:     "join-type",
			Usage:    "The type of join to perform.",
			BodyPath: "joinType",
		},
		&requestflag.Flag[string]{
			Name:     "shaping-schema",
			Usage:    "JMESPath expression that defines how to transform and customize the input payload structure.\nPayload shaping allows you to extract, reshape, and reorganize data from complex input payloads\ninto a simplified, standardized output format. Use JMESPath syntax to select specific fields,\nperform calculations, and create new data structures tailored to your needs.",
			BodyPath: "shapingSchema",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "config",
			Usage:    "Configuration for enrich function with semantic search steps.\n\n**How Enrich Functions Work:**\n\nEnrich functions use semantic search to augment JSON data with relevant information from collections.\nThey take JSON input (typically from a transform function), extract specified fields, perform vector-based\nsemantic search against collections, and inject the results back into the data.\n\n**Input Requirements:**\n- Must receive JSON input (typically uploaded to S3 from a previous function)\n- Can be chained after transform or other functions that produce JSON output\n\n**Example Use Cases:**\n- Match product descriptions to SKU codes from a product catalog\n- Enrich customer data with account information\n- Link order line items to inventory records\n\n**Configuration:**\n- Define one or more enrichment steps\n- Each step extracts values, searches a collection, and injects results\n- Steps are executed sequentially",
			BodyPath: "config",
		},
	},
	Action:          handleFunctionsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"route": {
		&requestflag.InnerFlag[string]{
			Name:       "route.name",
			InnerField: "name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "route.description",
			InnerField: "description",
		},
		&requestflag.InnerFlag[string]{
			Name:       "route.function-id",
			InnerField: "functionID",
		},
		&requestflag.InnerFlag[string]{
			Name:       "route.function-name",
			InnerField: "functionName",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "route.is-error-fallback",
			InnerField: "isErrorFallback",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "route.origin",
			InnerField: "origin",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "route.regex",
			InnerField: "regex",
		},
	},
	"print-page-split-config": {
		&requestflag.InnerFlag[string]{
			Name:       "print-page-split-config.next-function-id",
			InnerField: "nextFunctionID",
		},
		&requestflag.InnerFlag[string]{
			Name:       "print-page-split-config.next-function-name",
			InnerField: "nextFunctionName",
		},
	},
	"semantic-page-split-config": {
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "semantic-page-split-config.item-classes",
			InnerField: "itemClasses",
		},
	},
	"config": {
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "config.steps",
			Usage:      "Array of enrichment steps to execute sequentially",
			InnerField: "steps",
		},
	},
})

var functionsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a Function",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "function-name",
			Required: true,
		},
	},
	Action:          handleFunctionsRetrieve,
	HideHelpCommand: true,
}

var functionsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update a Function",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "path-function-name",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    `Allowed values: "transform".`,
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[string]{
			Name:     "display-name",
			Usage:    "Display name of function. Human-readable name to help you identify the function.",
			BodyPath: "displayName",
		},
		&requestflag.Flag[string]{
			Name:     "function-name",
			Usage:    "Name of function. Must be UNIQUE on a per-environment basis.",
			BodyPath: "functionName",
		},
		&requestflag.Flag[any]{
			Name:     "output-schema",
			Usage:    "Desired output structure defined in standard JSON Schema convention.",
			BodyPath: "outputSchema",
		},
		&requestflag.Flag[string]{
			Name:     "output-schema-name",
			Usage:    "Name of output schema object.",
			BodyPath: "outputSchemaName",
		},
		&requestflag.Flag[bool]{
			Name:     "tabular-chunking-enabled",
			Usage:    "Whether tabular chunking is enabled on the pipeline. This processes tables in CSV/Excel in row batches, rather than all rows at once.",
			BodyPath: "tabularChunkingEnabled",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Array of tags to categorize and organize functions.",
			BodyPath: "tags",
		},
		&requestflag.Flag[bool]{
			Name:     "enable-bounding-boxes",
			Usage:    "Whether bounding box extraction is enabled. Only applicable to analyze and extract functions.\nWhen true, the function returns the document regions (page, coordinates) from which each\nfield was extracted. Enabling this automatically configures the function to use the bounding\nbox model. Disabling resets to the default.",
			BodyPath: "enableBoundingBoxes",
		},
		&requestflag.Flag[bool]{
			Name:     "pre-count",
			Usage:    "Reducing the risk of the model stopping early on long documents.\nTrade-off: Increases total latency. Compatible with\n`enableBoundingBoxes`.",
			BodyPath: "preCount",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "Description of router. Can be used to provide additional context on router's purpose and expected inputs.",
			BodyPath: "description",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "route",
			Usage:    "List of routes.",
			BodyPath: "routes",
		},
		&requestflag.Flag[string]{
			Name:     "destination-type",
			Usage:    "Destination type for a Send function.",
			BodyPath: "destinationType",
		},
		&requestflag.Flag[string]{
			Name:     "google-drive-folder-id",
			Usage:    "Google Drive folder ID. Required when destinationType is google_drive. Managed via Paragon OAuth.",
			BodyPath: "googleDriveFolderId",
		},
		&requestflag.Flag[string]{
			Name:     "s3-bucket",
			Usage:    "S3 bucket to upload the payload to. Required when destinationType is s3.",
			BodyPath: "s3Bucket",
		},
		&requestflag.Flag[string]{
			Name:     "s3-prefix",
			Usage:    "Optional S3 key prefix (folder path).",
			BodyPath: "s3Prefix",
		},
		&requestflag.Flag[bool]{
			Name:     "webhook-signing-enabled",
			Usage:    "Whether to sign webhook deliveries with an HMAC-SHA256 `bem-signature` header.\nDefaults to `true` when omitted — signing is on by default for new send functions.\nSet explicitly to `false` to disable.",
			BodyPath: "webhookSigningEnabled",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-url",
			Usage:    "Webhook URL to POST the payload to. Required when destinationType is webhook.",
			BodyPath: "webhookUrl",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "print-page-split-config",
			BodyPath: "printPageSplitConfig",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "semantic-page-split-config",
			BodyPath: "semanticPageSplitConfig",
		},
		&requestflag.Flag[string]{
			Name:     "split-type",
			Usage:    `Allowed values: "print_page", "semantic_page".`,
			BodyPath: "splitType",
		},
		&requestflag.Flag[string]{
			Name:     "join-type",
			Usage:    "The type of join to perform.",
			BodyPath: "joinType",
		},
		&requestflag.Flag[string]{
			Name:     "shaping-schema",
			Usage:    "JMESPath expression that defines how to transform and customize the input payload structure.\nPayload shaping allows you to extract, reshape, and reorganize data from complex input payloads\ninto a simplified, standardized output format. Use JMESPath syntax to select specific fields,\nperform calculations, and create new data structures tailored to your needs.",
			BodyPath: "shapingSchema",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "config",
			Usage:    "Configuration for enrich function with semantic search steps.\n\n**How Enrich Functions Work:**\n\nEnrich functions use semantic search to augment JSON data with relevant information from collections.\nThey take JSON input (typically from a transform function), extract specified fields, perform vector-based\nsemantic search against collections, and inject the results back into the data.\n\n**Input Requirements:**\n- Must receive JSON input (typically uploaded to S3 from a previous function)\n- Can be chained after transform or other functions that produce JSON output\n\n**Example Use Cases:**\n- Match product descriptions to SKU codes from a product catalog\n- Enrich customer data with account information\n- Link order line items to inventory records\n\n**Configuration:**\n- Define one or more enrichment steps\n- Each step extracts values, searches a collection, and injects results\n- Steps are executed sequentially",
			BodyPath: "config",
		},
	},
	Action:          handleFunctionsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"route": {
		&requestflag.InnerFlag[string]{
			Name:       "route.name",
			InnerField: "name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "route.description",
			InnerField: "description",
		},
		&requestflag.InnerFlag[string]{
			Name:       "route.function-id",
			InnerField: "functionID",
		},
		&requestflag.InnerFlag[string]{
			Name:       "route.function-name",
			InnerField: "functionName",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "route.is-error-fallback",
			InnerField: "isErrorFallback",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "route.origin",
			InnerField: "origin",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "route.regex",
			InnerField: "regex",
		},
	},
	"print-page-split-config": {
		&requestflag.InnerFlag[string]{
			Name:       "print-page-split-config.next-function-id",
			InnerField: "nextFunctionID",
		},
		&requestflag.InnerFlag[string]{
			Name:       "print-page-split-config.next-function-name",
			InnerField: "nextFunctionName",
		},
	},
	"semantic-page-split-config": {
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "semantic-page-split-config.item-classes",
			InnerField: "itemClasses",
		},
	},
	"config": {
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "config.steps",
			Usage:      "Array of enrichment steps to execute sequentially",
			InnerField: "steps",
		},
	},
})

var functionsList = cli.Command{
	Name:    "list",
	Usage:   "List Functions",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "display-name",
			QueryPath: "displayName",
		},
		&requestflag.Flag[string]{
			Name:      "ending-before",
			QueryPath: "endingBefore",
		},
		&requestflag.Flag[[]string]{
			Name:      "function-id",
			QueryPath: "functionIDs",
		},
		&requestflag.Flag[[]string]{
			Name:      "function-name",
			QueryPath: "functionNames",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Default:   50,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "sort-order",
			Usage:     `Allowed values: "asc", "desc".`,
			Default:   "asc",
			QueryPath: "sortOrder",
		},
		&requestflag.Flag[string]{
			Name:      "starting-after",
			QueryPath: "startingAfter",
		},
		&requestflag.Flag[[]string]{
			Name:      "tag",
			QueryPath: "tags",
		},
		&requestflag.Flag[[]string]{
			Name:      "type",
			QueryPath: "types",
		},
		&requestflag.Flag[[]string]{
			Name:      "workflow-id",
			QueryPath: "workflowIDs",
		},
		&requestflag.Flag[[]string]{
			Name:      "workflow-name",
			QueryPath: "workflowNames",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleFunctionsList,
	HideHelpCommand: true,
}

var functionsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a Function",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "function-name",
			Required: true,
		},
	},
	Action:          handleFunctionsDelete,
	HideHelpCommand: true,
}

func handleFunctionsCreate(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := bem.FunctionNewParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Functions.New(ctx, params, options...)
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
		Title:          "functions create",
		Transform:      transform,
	})
}

func handleFunctionsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("function-name") && len(unusedArgs) > 0 {
		cmd.Set("function-name", unusedArgs[0])
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
	_, err = client.Functions.Get(ctx, cmd.Value("function-name").(string), options...)
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
		Title:          "functions retrieve",
		Transform:      transform,
	})
}

func handleFunctionsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("path-function-name") && len(unusedArgs) > 0 {
		cmd.Set("path-function-name", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := bem.FunctionUpdateParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Functions.Update(
		ctx,
		cmd.Value("path-function-name").(string),
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
		Title:          "functions update",
		Transform:      transform,
	})
}

func handleFunctionsList(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := bem.FunctionListParams{}

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

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Functions.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "functions list",
			Transform:      transform,
		})
	} else {
		iter := client.Functions.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "functions list",
			Transform:      transform,
		})
	}
}

func handleFunctionsDelete(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("function-name") && len(unusedArgs) > 0 {
		cmd.Set("function-name", unusedArgs[0])
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

	return client.Functions.Delete(ctx, cmd.Value("function-name").(string), options...)
}
