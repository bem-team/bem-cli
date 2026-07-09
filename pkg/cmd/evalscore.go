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
		&requestflag.Flag[[]map[string]any]{
			Name:     "pair",
			Usage:    "Up to 1000 pairs per request.",
			Required: true,
			BodyPath: "pairs",
		},
		&requestflag.Flag[int64]{
			Name:     "function-version-num",
			Usage:    "Optional version number to score against. P0: only the function's\ncurrent version is accepted; passing a different version returns 422.",
			BodyPath: "functionVersionNum",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "match-config",
			Usage:    "Comparator configuration. All fields optional; conservative defaults.",
			BodyPath: "matchConfig",
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
	"match-config": {
		&requestflag.InnerFlag[string]{
			Name:       "match-config.array-match",
			Usage:      "P0 supports only `by-index`.",
			InnerField: "arrayMatch",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "match-config.fuzzy-threshold",
			Usage:      "Levenshtein-ratio threshold used when `stringMatch == \"fuzzy\"`.\nRange `[0, 1]`. Default `0.85`.",
			InnerField: "fuzzyThreshold",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "match-config.ignore-paths",
			Usage:      "JSON Pointer paths to skip during comparison. The asterisk character\nmatches arbitrary object keys / array indices.\n\nExample values: /metadata, /lineItems with asterisk segment, etc.",
			InnerField: "ignorePaths",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "match-config.numeric-tolerance",
			Usage:      "Relative tolerance for numeric fields. `0` (default) means exact\nequality; `0.01` means ±1%.",
			InnerField: "numericTolerance",
		},
		&requestflag.InnerFlag[string]{
			Name:       "match-config.string-match",
			Usage:      "`exact` (default) or `fuzzy`.",
			InnerField: "stringMatch",
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
