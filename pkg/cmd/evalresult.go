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

var evalResultsFetchResults = cli.Command{
	Name:    "fetch-results",
	Usage:   "**Fetch evaluation results for a batch of transformations (POST).**",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "transformation-id",
			Usage:    "Transformation IDs to fetch results for. Up to 100 per request.",
			Required: true,
			BodyPath: "transformationIDs",
		},
		&requestflag.Flag[string]{
			Name:     "evaluation-version",
			Usage:    "Optional evaluation version filter.",
			BodyPath: "evaluationVersion",
		},
	},
	Action:          handleEvalResultsFetchResults,
	HideHelpCommand: true,
}

var evalResultsRetrieveResults = cli.Command{
	Name:    "retrieve-results",
	Usage:   "**Fetch evaluation results for a batch of transformations.**",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "transformation-ids",
			Usage:     "Comma-separated list of transformation IDs to fetch results for.\nBetween 1 and 100 IDs per request.",
			Required:  true,
			QueryPath: "transformationIDs",
		},
		&requestflag.Flag[string]{
			Name:      "evaluation-version",
			Usage:     "Optional evaluation version filter.",
			QueryPath: "evaluationVersion",
		},
	},
	Action:          handleEvalResultsRetrieveResults,
	HideHelpCommand: true,
}

func handleEvalResultsFetchResults(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := bem.EvalResultFetchResultsParams{}

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
	_, err = client.Eval.Results.FetchResults(ctx, params, options...)
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
		Title:          "eval:results fetch-results",
		Transform:      transform,
	})
}

func handleEvalResultsRetrieveResults(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := bem.EvalResultGetResultsParams{}

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
	_, err = client.Eval.Results.GetResults(ctx, params, options...)
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
		Title:          "eval:results retrieve-results",
		Transform:      transform,
	})
}
