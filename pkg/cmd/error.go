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

var errorsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "**Retrieve a single error event by ID.**",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "event-id",
			Required: true,
		},
	},
	Action:          handleErrorsRetrieve,
	HideHelpCommand: true,
}

var errorsList = cli.Command{
	Name:    "list",
	Usage:   "**List terminal error events.**",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:      "call-id",
			Usage:     "Filter to errors from specific calls.",
			QueryPath: "callIDs",
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
		&requestflag.Flag[[]string]{
			Name:      "reference-id",
			QueryPath: "referenceIDs",
		},
		&requestflag.Flag[string]{
			Name:      "reference-id-substring",
			Usage:     "Case-insensitive substring match against `referenceID`.",
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
	Action:          handleErrorsList,
	HideHelpCommand: true,
}

func handleErrorsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("event-id") && len(unusedArgs) > 0 {
		cmd.Set("event-id", unusedArgs[0])
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
	_, err = client.Errors.Get(ctx, cmd.Value("event-id").(string), options...)
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
		Title:          "errors retrieve",
		Transform:      transform,
	})
}

func handleErrorsList(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := bem.ErrorListParams{}

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
		_, err = client.Errors.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "errors list",
			Transform:      transform,
		})
	} else {
		iter := client.Errors.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "errors list",
			Transform:      transform,
		})
	}
}
