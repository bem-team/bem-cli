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

var evalScoreCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "**Score a function against a list of (input, expected) pairs.**",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "function-name",
			Usage:    "Name of the function to score. Must be of type extract, transform, or analyze.",
			Required: true,
			BodyPath: "functionName",
		},
		&requestflag.Flag[string]{
			Name:     "dataset-id",
			Usage:    "A saved Golden Data Set (`gds_…`) to score against. Mutually exclusive with\n`pairs`; provide exactly one. Its input / corrected / schema columns are resolved\nby column role. When it carries a `schema`-role column, scoring types each row\nagainst that ground-truth schema instead of the function's own schema — so results\nhold up as functions/schemas evolve.",
			BodyPath: "datasetID",
		},
		&requestflag.Flag[int64]{
			Name:     "function-version-num",
			Usage:    "Optional version number to score against. P0: only the function's\ncurrent version is accepted; passing a different version returns 422.",
			BodyPath: "functionVersionNum",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "pair",
			Usage:    "Inline `(input, expected)` pairs to score, up to 1000 per request.\nMutually exclusive with `datasetID`; provide exactly one.",
			BodyPath: "pairs",
		},
	},
	Action:          handleEvalScoreCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"pair": {
		&requestflag.InnerFlag[any]{
			Name:       "pair.expected",
			Usage:      "Expected output for this input, as a JSON value. The comparator walks\n`expected ∪ actual` and produces a per-leaf classification.",
			InnerField: "expected",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "pair.input",
			Usage:      "A single file input with base64-encoded content.\n\nWhen using the Bem CLI, use `@path/to/file` in the `inputContent` field to\nautomatically read and base64-encode the file:\n`--input.single-file '{\"inputContent\": \"@file.pdf\", \"inputType\": \"pdf\"}' --wait`",
			InnerField: "input",
		},
	},
})

var evalScoreRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "**Get the status and per-pair results of a score run.**",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "score-run-id",
			Required:  true,
			PathParam: "scoreRunID",
		},
	},
	Action:          handleEvalScoreRetrieve,
	HideHelpCommand: true,
}

var evalScoreCancel = cli.Command{
	Name:    "cancel",
	Usage:   "**Cancel an in-flight score run.**",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "score-run-id",
			Required:  true,
			PathParam: "scoreRunID",
		},
	},
	Action:          handleEvalScoreCancel,
	HideHelpCommand: true,
}

func handleEvalScoreCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := bem.EvalScoreNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Eval.Score.New(ctx, params, options...)
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
		Title:          "eval:score create",
		Transform:      transform,
	})
}

func handleEvalScoreRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("score-run-id") && len(unusedArgs) > 0 {
		cmd.Set("score-run-id", unusedArgs[0])
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
	_, err = client.Eval.Score.Get(ctx, cmd.Value("score-run-id").(string), options...)
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
		Title:          "eval:score retrieve",
		Transform:      transform,
	})
}

func handleEvalScoreCancel(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("score-run-id") && len(unusedArgs) > 0 {
		cmd.Set("score-run-id", unusedArgs[0])
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
	_, err = client.Eval.Score.Cancel(ctx, cmd.Value("score-run-id").(string), options...)
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
		Title:          "eval:score cancel",
		Transform:      transform,
	})
}
