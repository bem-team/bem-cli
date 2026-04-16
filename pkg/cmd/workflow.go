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

var workflowsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a Workflow",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "main-node-name",
			Usage:    "Name of the entry-point node. Must not be a destination of any edge.",
			Required: true,
			BodyPath: "mainNodeName",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Unique name for the workflow. Must match `^[a-zA-Z0-9_-]{1,128}$`.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "node",
			Usage:    "Call-site nodes in the DAG. At least one is required.",
			Required: true,
			BodyPath: "nodes",
		},
		&requestflag.Flag[string]{
			Name:     "display-name",
			Usage:    "Human-readable display name.",
			BodyPath: "displayName",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "edge",
			Usage:    "Directed edges between nodes. Omit or leave empty for single-node workflows.",
			BodyPath: "edges",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Tags to categorize and organize the workflow.",
			BodyPath: "tags",
		},
	},
	Action:          handleWorkflowsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"node": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "node.function",
			InnerField: "function",
		},
		&requestflag.InnerFlag[string]{
			Name:       "node.name",
			Usage:      "Name for this call site. Must be unique within the workflow version.\nDefaults to the function's own name when omitted.",
			InnerField: "name",
		},
	},
	"edge": {
		&requestflag.InnerFlag[string]{
			Name:       "edge.destination-node-name",
			Usage:      "Name of the destination node.",
			InnerField: "destinationNodeName",
		},
		&requestflag.InnerFlag[string]{
			Name:       "edge.source-node-name",
			Usage:      "Name of the source node.",
			InnerField: "sourceNodeName",
		},
		&requestflag.InnerFlag[string]{
			Name:       "edge.destination-name",
			Usage:      "Labelled outlet on the source node that activates this edge.\nOmit for the default (unlabelled) outlet.",
			InnerField: "destinationName",
		},
	},
})

var workflowsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a Workflow",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "workflow-name",
			Required: true,
		},
	},
	Action:          handleWorkflowsRetrieve,
	HideHelpCommand: true,
}

var workflowsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update a Workflow",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "workflow-name",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "display-name",
			Usage:    "Human-readable display name.",
			BodyPath: "displayName",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "edge",
			BodyPath: "edges",
		},
		&requestflag.Flag[string]{
			Name:     "main-node-name",
			Usage:    "`mainNodeName`, `nodes`, and `edges` must be provided together to update the DAG\ntopology. If none are provided the topology is copied unchanged from the current\nversion.",
			BodyPath: "mainNodeName",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "New name for the workflow (renames it). Must match `^[a-zA-Z0-9_-]{1,128}$`.",
			BodyPath: "name",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "node",
			BodyPath: "nodes",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Tags to categorize and organize the workflow.",
			BodyPath: "tags",
		},
	},
	Action:          handleWorkflowsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"edge": {
		&requestflag.InnerFlag[string]{
			Name:       "edge.destination-node-name",
			Usage:      "Name of the destination node.",
			InnerField: "destinationNodeName",
		},
		&requestflag.InnerFlag[string]{
			Name:       "edge.source-node-name",
			Usage:      "Name of the source node.",
			InnerField: "sourceNodeName",
		},
		&requestflag.InnerFlag[string]{
			Name:       "edge.destination-name",
			Usage:      "Labelled outlet on the source node that activates this edge.\nOmit for the default (unlabelled) outlet.",
			InnerField: "destinationName",
		},
	},
	"node": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "node.function",
			InnerField: "function",
		},
		&requestflag.InnerFlag[string]{
			Name:       "node.name",
			Usage:      "Name for this call site. Must be unique within the workflow version.\nDefaults to the function's own name when omitted.",
			InnerField: "name",
		},
	},
})

var workflowsList = cli.Command{
	Name:    "list",
	Usage:   "List Workflows",
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
	Action:          handleWorkflowsList,
	HideHelpCommand: true,
}

var workflowsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a Workflow",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "workflow-name",
			Required: true,
		},
	},
	Action:          handleWorkflowsDelete,
	HideHelpCommand: true,
}

var workflowsCall = requestflag.WithInnerFlags(cli.Command{
	Name:    "call",
	Usage:   "**Invoke a workflow.**",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "workflow-name",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "input",
			Usage:    "Input file(s) for a call. Provide exactly one of `singleFile` or `batchFiles`.\n\nIn the CLI, use the nested flags `--input.single-file` or `--input.batch-files`\nwith `@path/to/file` for automatic file embedding:\n`--input.single-file '{\"inputContent\": \"@invoice.pdf\", \"inputType\": \"pdf\"}' --wait`",
			Required: true,
			BodyPath: "input",
		},
		&requestflag.Flag[bool]{
			Name:      "wait",
			Usage:     "Block until the call completes (up to 30 seconds) and return the finished\ncall object. Default: `false`. This is a boolean flag — use `--wait` or\n`--wait=true`, not `--wait true`.",
			QueryPath: "wait",
		},
		&requestflag.Flag[string]{
			Name:     "call-reference-id",
			Usage:    "Your reference ID for tracking this call.",
			BodyPath: "callReferenceID",
		},
		&requestflag.Flag[any]{
			Name:     "metadata",
			Usage:    "Arbitrary JSON object attached to this call. Stored on the call record and injected\ninto `transformedContent` under the reserved `_metadata` key (alongside `referenceID`).\nMust be a JSON object. Maximum size: 4 KB.",
			BodyPath: "metadata",
		},
	},
	Action:          handleWorkflowsCall,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"input": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "input.batch-files",
			Usage:      "Multiple files to process in one call. Each item in the `inputs` array has its own `inputContent` and `inputType`.",
			InnerField: "batchFiles",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "input.single-file",
			Usage:      "A single file input with base64-encoded content.\n\nWhen using the Bem CLI, use `@path/to/file` in the `inputContent` field to\nautomatically read and base64-encode the file:\n`--input.single-file '{\"inputContent\": \"@file.pdf\", \"inputType\": \"pdf\"}' --wait`",
			InnerField: "singleFile",
		},
	},
})

var workflowsCopy = cli.Command{
	Name:    "copy",
	Usage:   "Copy a Workflow",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "source-workflow-name",
			Usage:    "Name of the source workflow to copy from.",
			Required: true,
			BodyPath: "sourceWorkflowName",
		},
		&requestflag.Flag[string]{
			Name:     "target-workflow-name",
			Usage:    "Name for the new copied workflow. Must be unique within the target environment.",
			Required: true,
			BodyPath: "targetWorkflowName",
		},
		&requestflag.Flag[int64]{
			Name:     "source-workflow-version-num",
			Usage:    "Optional version number of the source workflow to copy. If not provided, copies the current version.",
			BodyPath: "sourceWorkflowVersionNum",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Optional tags for the copied workflow. If not provided, uses the source workflow's tags.",
			BodyPath: "tags",
		},
		&requestflag.Flag[string]{
			Name:     "target-display-name",
			Usage:    `Optional display name for the copied workflow. If not provided, uses the source workflow's display name with " (Copy)" appended.`,
			BodyPath: "targetDisplayName",
		},
		&requestflag.Flag[string]{
			Name:     "target-environment",
			Usage:    "Optional target environment name. If provided, copies the workflow to a different environment. When copying to a different environment, all functions used in the workflow will also be copied.",
			BodyPath: "targetEnvironment",
		},
	},
	Action:          handleWorkflowsCopy,
	HideHelpCommand: true,
}

func handleWorkflowsCreate(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := bem.WorkflowNewParams{}

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
	_, err = client.Workflows.New(ctx, params, options...)
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
		Title:          "workflows create",
		Transform:      transform,
	})
}

func handleWorkflowsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("workflow-name") && len(unusedArgs) > 0 {
		cmd.Set("workflow-name", unusedArgs[0])
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
	_, err = client.Workflows.Get(ctx, cmd.Value("workflow-name").(string), options...)
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
		Title:          "workflows retrieve",
		Transform:      transform,
	})
}

func handleWorkflowsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("workflow-name") && len(unusedArgs) > 0 {
		cmd.Set("workflow-name", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := bem.WorkflowUpdateParams{}

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
	_, err = client.Workflows.Update(
		ctx,
		cmd.Value("workflow-name").(string),
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
		Title:          "workflows update",
		Transform:      transform,
	})
}

func handleWorkflowsList(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := bem.WorkflowListParams{}

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
		_, err = client.Workflows.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			Title:          "workflows list",
			Transform:      transform,
		})
	} else {
		iter := client.Workflows.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			Title:          "workflows list",
			Transform:      transform,
		})
	}
}

func handleWorkflowsDelete(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("workflow-name") && len(unusedArgs) > 0 {
		cmd.Set("workflow-name", unusedArgs[0])
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

	return client.Workflows.Delete(ctx, cmd.Value("workflow-name").(string), options...)
}

func handleWorkflowsCall(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("workflow-name") && len(unusedArgs) > 0 {
		cmd.Set("workflow-name", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := bem.WorkflowCallParams{}

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
	_, err = client.Workflows.Call(
		ctx,
		cmd.Value("workflow-name").(string),
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
		Title:          "workflows call",
		Transform:      transform,
	})
}

func handleWorkflowsCopy(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := bem.WorkflowCopyParams{}

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
	_, err = client.Workflows.Copy(ctx, params, options...)
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
		Title:          "workflows copy",
		Transform:      transform,
	})
}
