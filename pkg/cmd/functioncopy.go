// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/bem-team/bem-go-sdk"
	"github.com/bem-team/bem-go-sdk/option"
	"github.com/stainless-sdks/bem-cli/internal/apiquery"
	"github.com/stainless-sdks/bem-cli/internal/requestflag"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var functionsCopyCreate = cli.Command{
	Name:    "create",
	Usage:   "Copy a Function",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "source-function-name",
			Usage:    "Name of the function to copy from. Must be a valid existing function name.",
			Required: true,
			BodyPath: "sourceFunctionName",
		},
		&requestflag.Flag[string]{
			Name:     "target-function-name",
			Usage:    "Name for the new copied function. Must be unique within the target environment.",
			Required: true,
			BodyPath: "targetFunctionName",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Optional array of tags for the copied function. If not provided, defaults to the source function's tags.",
			BodyPath: "tags",
		},
		&requestflag.Flag[string]{
			Name:     "target-display-name",
			Usage:    `Optional display name for the copied function. If not provided, defaults to the source function's display name with " (Copy)" appended.`,
			BodyPath: "targetDisplayName",
		},
		&requestflag.Flag[string]{
			Name:     "target-environment",
			Usage:    "Optional environment name to copy the function to. If not provided, the function will be copied within the same environment.",
			BodyPath: "targetEnvironment",
		},
	},
	Action:          handleFunctionsCopyCreate,
	HideHelpCommand: true,
}

func handleFunctionsCopyCreate(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := bem.FunctionCopyNewParams{}

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
	_, err = client.Functions.Copy.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "functions:copy create", obj, format, transform)
}
