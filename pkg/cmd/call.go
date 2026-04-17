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

var callsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "**Retrieve a workflow call by ID.**",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "call-id",
			Required: true,
		},
	},
	Action:          handleCallsRetrieve,
	HideHelpCommand: true,
}

var callsList = cli.Command{
	Name:    "list",
	Usage:   "**List workflow calls with filtering and pagination.**",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:      "call-id",
			QueryPath: "callIDs",
		},
		&requestflag.Flag[string]{
			Name:      "ending-before",
			QueryPath: "endingBefore",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Default:   50,
			QueryPath: "limit",
		},
		&requestflag.Flag[[]string]{
			Name:      "reference-id",
			QueryPath: "referenceIDs",
		},
		&requestflag.Flag[string]{
			Name:      "reference-id-substring",
			Usage:     "Case-insensitive substring match against `callReferenceID`.",
			QueryPath: "referenceIDSubstring",
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
			Name:      "status",
			Usage:     "Filter by one or more statuses.",
			QueryPath: "statuses",
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
	Action:          handleCallsList,
	HideHelpCommand: true,
}

func handleCallsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("call-id") && len(unusedArgs) > 0 {
		cmd.Set("call-id", unusedArgs[0])
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
	_, err = client.Calls.Get(ctx, cmd.Value("call-id").(string), options...)
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
		Title:          "calls retrieve",
		Transform:      transform,
	})
}

func handleCallsList(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := bem.CallListParams{}

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
		_, err = client.Calls.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "calls list",
			Transform:      transform,
		})
	} else {
		iter := client.Calls.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "calls list",
			Transform:      transform,
		})
	}
}
