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

var reviewQueueList = cli.Command{
	Name:    "list",
	Usage:   "**List entities awaiting curation, for a human reviewer's queue.**",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "assigned-to",
			Usage:     "`me` or a `usr_...` ID — restrict to entities whose effective type that user reviews.",
			QueryPath: "assignedTo",
		},
		&requestflag.Flag[string]{
			Name:      "bucket",
			Usage:     "Optional bucket public ID (`bkt_...`) to scope to. Omit for all buckets.",
			QueryPath: "bucket",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Cursor — an `entityID` defining your place in the list.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Default:   50,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "since",
			Usage:     "RFC3339 timestamp — restrict to entities created at or after this time.",
			QueryPath: "since",
		},
		&requestflag.Flag[[]string]{
			Name:      "status",
			Usage:     "Restrict to these lifecycle states. Defaults to `extracted` + `proposed`.",
			QueryPath: "status",
		},
		&requestflag.Flag[[]string]{
			Name:      "type",
			Usage:     "Restrict to entities whose effective type is one of these `ety_...` IDs.",
			QueryPath: "type",
		},
	},
	Action:          handleReviewQueueList,
	HideHelpCommand: true,
}

func handleReviewQueueList(ctx context.Context, cmd *cli.Command) error {
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

	params := bem.ReviewQueueListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.ReviewQueue.List(ctx, params, options...)
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
		Title:          "review-queue list",
		Transform:      transform,
	})
}
