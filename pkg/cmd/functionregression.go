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

var functionsRegressionApplyCorrections = cli.Command{
	Name:    "apply-corrections",
	Usage:   "**Copy baseline corrections onto regression transformations.**",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "baseline-version-num",
			Usage:    "**Baseline version number (source of corrected data)**\n\nThe function version number that contains transformations with corrected JSON\nthat should be copied to regression transformations.",
			Required: true,
			BodyPath: "baselineVersionNum",
		},
		&requestflag.Flag[int64]{
			Name:     "comparison-version-num",
			Usage:    "**Comparison version number (target for applying corrections)**\n\nThe function version number of regression transformations that should\nreceive the corrected JSON from the baseline version.",
			Required: true,
			BodyPath: "comparisonVersionNum",
		},
		&requestflag.Flag[string]{
			Name:     "function-name",
			Usage:    "**Name of the function to apply corrections for**\n\nMust be an existing function with both baseline and regression transformation data.",
			Required: true,
			BodyPath: "functionName",
		},
	},
	Action:          handleFunctionsRegressionApplyCorrections,
	HideHelpCommand: true,
}

var functionsRegressionRun = cli.Command{
	Name:    "run",
	Usage:   "**Kick off a regression run between two versions of a function.**",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "function-name",
			Usage:    "**Name of the function to test for regressions**\n\nMust be an existing function with historical transformation data containing user corrections.\nThe function must be currently active and callable.",
			Required: true,
			BodyPath: "functionName",
		},
		&requestflag.Flag[int64]{
			Name:     "baseline-version-num",
			Usage:    "**Function version number to use as baseline for comparison**\n\n- Defaults to `currentVersionNum - 1` (previous version)\n- Must be a valid, existing version number for the function\n- Used to retrieve historical transformation data for comparison\n- Cannot be the same as `comparisonVersionNum`",
			BodyPath: "baselineVersionNum",
		},
		&requestflag.Flag[int64]{
			Name:     "comparison-version-num",
			Usage:    "**Function version number to test against the baseline**\n\n- Defaults to current version number (latest version)\n- Must be a valid, existing version number for the function\n- This version will be used to create new function calls for testing\n- Cannot be the same as `baselineVersionNum`",
			BodyPath: "comparisonVersionNum",
		},
		&requestflag.Flag[bool]{
			Name:     "only-corrected-data",
			Usage:    "**Whether to only test transformations with user corrections**\n\n- Defaults to `true` (recommended)\n- When `true`: Only uses transformations with `correctedJSON` as ground truth\n- When `false`: May include transformations without corrections (less reliable)\n- Corrected data provides the most accurate regression testing results",
			Default:  true,
			BodyPath: "onlyCorrectedData",
		},
		&requestflag.Flag[int64]{
			Name:     "sample-size",
			Usage:    "**Number of historical samples to test**\n\n- Defaults to 50 samples\n- Minimum: 1, Maximum: 1000\n- Only transformations with `correctedJSON` (user corrections) are eligible\n- Actual sample size may be smaller if insufficient corrected data exists\n- Larger samples provide more statistical confidence but take longer to process",
			Default:  50,
			BodyPath: "sampleSize",
		},
	},
	Action:          handleFunctionsRegressionRun,
	HideHelpCommand: true,
}

func handleFunctionsRegressionApplyCorrections(ctx context.Context, cmd *cli.Command) error {
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

	params := bem.FunctionRegressionApplyCorrectionsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Functions.Regression.ApplyCorrections(ctx, params, options...)
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
		Title:          "functions:regression apply-corrections",
		Transform:      transform,
	})
}

func handleFunctionsRegressionRun(ctx context.Context, cmd *cli.Command) error {
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

	params := bem.FunctionRegressionRunParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Functions.Regression.Run(ctx, params, options...)
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
		Title:          "functions:regression run",
		Transform:      transform,
	})
}
