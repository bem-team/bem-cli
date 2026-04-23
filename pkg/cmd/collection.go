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

var collectionsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a Collection",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "collection-name",
			Usage:    "Unique name/path for the collection. Supports dot notation for hierarchical paths.\n\n- Only letters (a-z, A-Z), digits (0-9), underscores (_), and dots (.) are allowed\n- Each segment (between dots) must start with a letter or underscore (not a digit)\n- Segments cannot consist only of digits\n- Each segment must be 1-256 characters\n- No leading, trailing, or consecutive dots\n- Invalid names are rejected with a 400 Bad Request error\n\n**Valid Examples:**\n- 'product_catalog'\n- 'orders.line_items.sku'\n- 'customer_data'\n- 'price_v2'\n\n**Invalid Examples:**\n- 'product-catalog' (contains hyphen)\n- '123items' (starts with digit)\n- 'items..data' (consecutive dots)\n- 'order#123' (contains invalid character #)",
			Required: true,
			BodyPath: "collectionName",
		},
	},
	Action:          handleCollectionsCreate,
	HideHelpCommand: true,
}

var collectionsList = cli.Command{
	Name:    "list",
	Usage:   "List Collections",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "collection-name-search",
			Usage:     "Optional substring search filter for collection names (case-insensitive).\nFor example, \"premium\" will match \"customers.premium\", \"products.premium\", etc.",
			QueryPath: "collectionNameSearch",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Number of collections per page",
			Default:   50,
			QueryPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:      "page",
			Usage:     "Page number for pagination",
			Default:   1,
			QueryPath: "page",
		},
		&requestflag.Flag[string]{
			Name:      "parent-collection-name",
			Usage:     "Optional filter to list only collections under a specific parent collection path.\nFor example, \"customers\" will return \"customers\", \"customers.premium\", \"customers.premium.vip\", etc.",
			QueryPath: "parentCollectionName",
		},
	},
	Action:          handleCollectionsList,
	HideHelpCommand: true,
}

var collectionsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a Collection",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "collection-name",
			Usage:     "The name/path of the collection to delete. Must use only letters, digits, underscores, and dots.\nEach segment must start with a letter or underscore.",
			Required:  true,
			QueryPath: "collectionName",
		},
	},
	Action:          handleCollectionsDelete,
	HideHelpCommand: true,
}

var collectionsCountTokens = cli.Command{
	Name:    "count-tokens",
	Usage:   "Count the number of tokens in the provided texts using the BGE M3 tokenizer.\nThis is useful for checking if texts will fit within the embedding model's token\nlimit (8,192 tokens per text) before sending them for embedding.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "text",
			Usage:    "One or more texts to tokenize.",
			Required: true,
			BodyPath: "texts",
		},
	},
	Action:          handleCollectionsCountTokens,
	HideHelpCommand: true,
}

func handleCollectionsCreate(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := bem.CollectionNewParams{}

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
	_, err = client.Collections.New(ctx, params, options...)
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
		Title:          "collections create",
		Transform:      transform,
	})
}

func handleCollectionsList(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := bem.CollectionListParams{}

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
	_, err = client.Collections.List(ctx, params, options...)
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
		Title:          "collections list",
		Transform:      transform,
	})
}

func handleCollectionsDelete(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := bem.CollectionDeleteParams{}

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

	return client.Collections.Delete(ctx, params, options...)
}

func handleCollectionsCountTokens(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := bem.CollectionCountTokensParams{}

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
	_, err = client.Collections.CountTokens(ctx, params, options...)
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
		Title:          "collections count-tokens",
		Transform:      transform,
	})
}
