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

var collectionsItemsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a Collection",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "collection-name",
			Usage:     "The name/path of the collection. Must use only letters, digits, underscores, and dots.\nEach segment must start with a letter or underscore.",
			Required:  true,
			QueryPath: "collectionName",
		},
		&requestflag.Flag[bool]{
			Name:      "include-subcollections",
			Usage:     "When true, includes items from all subcollections under the specified collection path.\nFor example, querying \"customers\" with this flag will return items from \"customers\",\n\"customers.premium\", \"customers.premium.vip\", etc.",
			QueryPath: "includeSubcollections",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Number of items per page",
			Default:   50,
			QueryPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:      "page",
			Usage:     "Page number for pagination",
			Default:   1,
			QueryPath: "page",
		},
	},
	Action:          handleCollectionsItemsRetrieve,
	HideHelpCommand: true,
}

var collectionsItemsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update existing items in a Collection",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "collection-name",
			Usage:    "The name/path of the collection. Must use only letters, digits, underscores, and dots.\nEach segment must start with a letter or underscore.",
			Required: true,
			BodyPath: "collectionName",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "item",
			Usage:    "Array of items to update (maximum 100 items per request)",
			Required: true,
			BodyPath: "items",
		},
	},
	Action:          handleCollectionsItemsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"item": {
		&requestflag.InnerFlag[string]{
			Name:       "item.collection-item-id",
			Usage:      "Unique identifier for the item to update",
			InnerField: "collectionItemID",
		},
		&requestflag.InnerFlag[any]{
			Name:       "item.data",
			Usage:      "The updated data to be embedded and stored (string or JSON object)",
			InnerField: "data",
		},
	},
})

var collectionsItemsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete an item from a Collection",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "collection-item-id",
			Usage:     "The unique identifier of the item to delete",
			Required:  true,
			QueryPath: "collectionItemID",
		},
		&requestflag.Flag[string]{
			Name:      "collection-name",
			Usage:     "The name/path of the collection. Must use only letters, digits, underscores, and dots.\nEach segment must start with a letter or underscore.",
			Required:  true,
			QueryPath: "collectionName",
		},
	},
	Action:          handleCollectionsItemsDelete,
	HideHelpCommand: true,
}

var collectionsItemsAdd = requestflag.WithInnerFlags(cli.Command{
	Name:    "add",
	Usage:   "Add new items to a Collection",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "collection-name",
			Usage:    "The name/path of the collection. Must use only letters, digits, underscores, and dots.\nEach segment must start with a letter or underscore.",
			Required: true,
			BodyPath: "collectionName",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "item",
			Usage:    "Array of items to add (maximum 100 items per request)",
			Required: true,
			BodyPath: "items",
		},
	},
	Action:          handleCollectionsItemsAdd,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"item": {
		&requestflag.InnerFlag[any]{
			Name:       "item.data",
			Usage:      "The data to be embedded and stored (string or JSON object)",
			InnerField: "data",
		},
	},
})

func handleCollectionsItemsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := bem.CollectionItemGetParams{}

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
	_, err = client.Collections.Items.Get(ctx, params, options...)
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
		Title:          "collections:items retrieve",
		Transform:      transform,
	})
}

func handleCollectionsItemsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := bem.CollectionItemUpdateParams{}

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
	_, err = client.Collections.Items.Update(ctx, params, options...)
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
		Title:          "collections:items update",
		Transform:      transform,
	})
}

func handleCollectionsItemsDelete(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := bem.CollectionItemDeleteParams{}

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

	return client.Collections.Items.Delete(ctx, params, options...)
}

func handleCollectionsItemsAdd(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := bem.CollectionItemAddParams{}

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
	_, err = client.Collections.Items.Add(ctx, params, options...)
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
		Title:          "collections:items add",
		Transform:      transform,
	})
}
