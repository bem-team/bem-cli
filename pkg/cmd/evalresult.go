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

var evalResultsRetrieveResults = cli.Command{
	Name:    "retrieve-results",
	Usage:   "**Fetch evaluation results for a batch of events.**",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "evaluation-version",
			Usage:     "Optional evaluation version filter.",
			QueryPath: "evaluationVersion",
		},
		&requestflag.Flag[string]{
			Name:      "event-ids",
			Usage:     "Comma-separated list of event KSUIDs to fetch results for. Between\n1 and 100 IDs per request. Mutually exclusive with\n`transformationIDs`.",
			QueryPath: "eventIDs",
		},
		&requestflag.Flag[string]{
			Name:      "transformation-ids",
			Usage:     "Comma-separated list of transformation IDs to fetch results for.\nBetween 1 and 100 IDs per request. Mutually exclusive with\n`eventIDs`. Prefer `eventIDs` for new integrations.",
			QueryPath: "transformationIDs",
		},
	},
	Action:          handleEvalResultsRetrieveResults,
	HideHelpCommand: true,
}

func handleEvalResultsRetrieveResults(ctx context.Context, cmd *cli.Command) error {
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

	params := bem.EvalResultGetResultsParams{}

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
