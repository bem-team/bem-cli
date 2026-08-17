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

var entityTypesCreate = cli.Command{
	Name:    "create",
	Usage:   "Create an Entity Type",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Type name. Required and unique within the account+environment.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:     "attribute-schema",
			Usage:    "Optional per-type structured attribute metadata.",
			BodyPath: "attributeSchema",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "Optional description.",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "parent-type-id",
			Usage:    "Optional public ID (`ety_...`) of the parent type. Must belong to the\nsame account+environment.",
			BodyPath: "parentTypeID",
		},
	},
	Action:          handleEntityTypesCreate,
	HideHelpCommand: true,
}

var entityTypesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get an Entity Type",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "type-id",
			Required:  true,
			PathParam: "typeID",
		},
	},
	Action:          handleEntityTypesRetrieve,
	HideHelpCommand: true,
}

var entityTypesUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update an Entity Type",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "type-id",
			Required:  true,
			PathParam: "typeID",
		},
		&requestflag.Flag[any]{
			Name:     "attribute-schema",
			Usage:    "New per-type structured attribute metadata.",
			BodyPath: "attributeSchema",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "New description.",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "parent-type-id",
			Usage:    "New parent type public ID (`ety_...`), or an empty string to clear the\nparent (promote to top-level). Must belong to the same\naccount+environment and may not be the type itself.",
			BodyPath: "parentTypeID",
		},
	},
	Action:          handleEntityTypesUpdate,
	HideHelpCommand: true,
}

var entityTypesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete an Entity Type",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "type-id",
			Required:  true,
			PathParam: "typeID",
		},
	},
	Action:          handleEntityTypesDelete,
	HideHelpCommand: true,
}

func handleEntityTypesCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := bem.EntityTypeNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EntityTypes.New(ctx, params, options...)
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
		Title:          "entity-types create",
		Transform:      transform,
	})
}

func handleEntityTypesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("type-id") && len(unusedArgs) > 0 {
		cmd.Set("type-id", unusedArgs[0])
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
	_, err = client.EntityTypes.Get(ctx, cmd.Value("type-id").(string), options...)
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
		Title:          "entity-types retrieve",
		Transform:      transform,
	})
}

func handleEntityTypesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("type-id") && len(unusedArgs) > 0 {
		cmd.Set("type-id", unusedArgs[0])
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

	params := bem.EntityTypeUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EntityTypes.Update(
		ctx,
		cmd.Value("type-id").(string),
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
		Title:          "entity-types update",
		Transform:      transform,
	})
}

func handleEntityTypesDelete(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("type-id") && len(unusedArgs) > 0 {
		cmd.Set("type-id", unusedArgs[0])
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

	return client.EntityTypes.Delete(ctx, cmd.Value("type-id").(string), options...)
}
