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

var evalTriggerEvaluation = cli.Command{
	Name:    "trigger-evaluation",
	Usage:   "**Queue evaluation jobs for a batch of transformations.**",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "transformation-id",
			Usage:    "Transformation IDs to evaluate. Up to 100 per request.",
			Required: true,
			BodyPath: "transformationIDs",
		},
		&requestflag.Flag[string]{
			Name:     "evaluation-version",
			Usage:    "Optional evaluation version (e.g. `0.1.0-gemini`). When omitted the\nserver's default evaluation version is used.",
			BodyPath: "evaluationVersion",
		},
	},
	Action:          handleEvalTriggerEvaluation,
	HideHelpCommand: true,
}

func handleEvalTriggerEvaluation(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := bem.EvalTriggerEvaluationParams{}

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
	_, err = client.Eval.TriggerEvaluation(ctx, params, options...)
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
		Title:          "eval trigger-evaluation",
		Transform:      transform,
	})
}
