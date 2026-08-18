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

var bucketsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a Bucket",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Bucket name. Required and unique within the account+environment.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "Optional description.",
			BodyPath: "description",
		},
	},
	Action:          handleBucketsCreate,
	HideHelpCommand: true,
}

var bucketsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a Bucket",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "bucket-id",
			Required:  true,
			PathParam: "bucketID",
		},
	},
	Action:          handleBucketsRetrieve,
	HideHelpCommand: true,
}

var bucketsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a Bucket",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "bucket-id",
			Required:  true,
			PathParam: "bucketID",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "New description.",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "New name.",
			BodyPath: "name",
		},
	},
	Action:          handleBucketsUpdate,
	HideHelpCommand: true,
}

var bucketsList = cli.Command{
	Name:    "list",
	Usage:   "List Buckets",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "ending-before",
			Usage:     "Cursor: return buckets whose `bucketID` sorts before this value.",
			QueryPath: "endingBefore",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of buckets to return (default 50, max 200).",
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "name-substring",
			Usage:     "Case-insensitive substring match on the bucket name.",
			QueryPath: "nameSubstring",
		},
		&requestflag.Flag[string]{
			Name:      "starting-after",
			Usage:     "Cursor: return buckets whose `bucketID` sorts after this value.",
			QueryPath: "startingAfter",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleBucketsList,
	HideHelpCommand: true,
}

var bucketsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a Bucket",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "bucket-id",
			Required:  true,
			PathParam: "bucketID",
		},
		&requestflag.Flag[bool]{
			Name:      "cascade",
			Usage:     "When `true`, delete the bucket even if it still contains entities\n(the entities are removed along with it). When omitted or `false`, the\nrequest is rejected with `409 Conflict` if the bucket is non-empty.\n\nThe default bucket can never be deleted regardless of this flag.",
			QueryPath: "cascade",
		},
	},
	Action:          handleBucketsDelete,
	HideHelpCommand: true,
}

func handleBucketsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := bem.BucketNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Buckets.New(ctx, params, options...)
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
		Title:          "buckets create",
		Transform:      transform,
	})
}

func handleBucketsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("bucket-id") && len(unusedArgs) > 0 {
		cmd.Set("bucket-id", unusedArgs[0])
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
	_, err = client.Buckets.Get(ctx, cmd.Value("bucket-id").(string), options...)
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
		Title:          "buckets retrieve",
		Transform:      transform,
	})
}

func handleBucketsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("bucket-id") && len(unusedArgs) > 0 {
		cmd.Set("bucket-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
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

	params := bem.BucketUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Buckets.Update(
		ctx,
		cmd.Value("bucket-id").(string),
		params,
		options...,
	)
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
		Title:          "buckets update",
		Transform:      transform,
	})
}

func handleBucketsList(ctx context.Context, cmd *cli.Command) error {
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

	params := bem.BucketListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Buckets.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "buckets list",
			Transform:      transform,
		})
	} else {
		iter := client.Buckets.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "buckets list",
			Transform:      transform,
		})
	}
}

func handleBucketsDelete(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("bucket-id") && len(unusedArgs) > 0 {
		cmd.Set("bucket-id", unusedArgs[0])
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

	params := bem.BucketDeleteParams{}

	return client.Buckets.Delete(
		ctx,
		cmd.Value("bucket-id").(string),
		params,
		options...,
	)
}
