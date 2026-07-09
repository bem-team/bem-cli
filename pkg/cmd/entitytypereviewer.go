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

var entityTypesReviewersList = cli.Command{
	Name:    "list",
	Usage:   "List Reviewers",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "type-id",
			Required:  true,
			PathParam: "typeID",
		},
	},
	Action:          handleEntityTypesReviewersList,
	HideHelpCommand: true,
}

var entityTypesReviewersAssign = cli.Command{
	Name:    "assign",
	Usage:   "Assign a Reviewer",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "type-id",
			Required:  true,
			PathParam: "typeID",
		},
		&requestflag.Flag[string]{
			Name:     "user-id",
			Usage:    "Public ID (`usr_...`) of the user to assign. Must belong to the account.",
			Required: true,
			BodyPath: "userID",
		},
	},
	Action:          handleEntityTypesReviewersAssign,
	HideHelpCommand: true,
}

var entityTypesReviewersRemove = cli.Command{
	Name:    "remove",
	Usage:   "Remove a Reviewer",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "type-id",
			Required:  true,
			PathParam: "typeID",
		},
		&requestflag.Flag[string]{
			Name:      "user-id",
			Required:  true,
			PathParam: "userID",
		},
	},
	Action:          handleEntityTypesReviewersRemove,
	HideHelpCommand: true,
}

func handleEntityTypesReviewersList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.EntityTypes.Reviewers.List(ctx, cmd.Value("type-id").(string), options...)
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
		Title:          "entity-types:reviewers list",
		Transform:      transform,
	})
}

func handleEntityTypesReviewersAssign(ctx context.Context, cmd *cli.Command) error {
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

	params := bem.EntityTypeReviewerAssignParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.EntityTypes.Reviewers.Assign(
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
		Title:          "entity-types:reviewers assign",
		Transform:      transform,
	})
}

func handleEntityTypesReviewersRemove(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("user-id") && len(unusedArgs) > 0 {
		cmd.Set("user-id", unusedArgs[0])
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

	params := bem.EntityTypeReviewerRemoveParams{
		TypeID: cmd.Value("type-id").(string),
	}

	return client.EntityTypes.Reviewers.Remove(
		ctx,
		cmd.Value("user-id").(string),
		params,
		options...,
	)
}
