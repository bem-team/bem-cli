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

var entitiesSynonymsAdd = cli.Command{
	Name:    "add",
	Usage:   "Add a Synonym to an Entity",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:     "text",
			Usage:    "The human-readable synonym surface form to attach (e.g. `Acme Corp`,\n`ACME`). It is normalized (lowercased, whitespace-folded) for the\nuniqueness key and the matcher's exact-match path.",
			Required: true,
			BodyPath: "text",
		},
		&requestflag.Flag[string]{
			Name:      "bucket",
			Usage:     "Optional bucket public ID (`bkt_...`) to scope the entity lookup to one\nbucket. Omit for the unscoped (all account+environment) view.",
			QueryPath: "bucket",
		},
		&requestflag.Flag[string]{
			Name:     "locale",
			Usage:    "Optional BCP 47 locale tag (e.g. `en-US`) for language-specific synonyms.",
			BodyPath: "locale",
		},
	},
	Action:          handleEntitiesSynonymsAdd,
	HideHelpCommand: true,
}

var entitiesSynonymsRemove = cli.Command{
	Name:    "remove",
	Usage:   "Remove a Synonym from an Entity",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "synonym-id",
			Required:  true,
			PathParam: "synonymID",
		},
		&requestflag.Flag[string]{
			Name:      "bucket",
			Usage:     "Optional bucket public ID (`bkt_...`) to scope the entity lookup to one\nbucket. Omit for the unscoped (all account+environment) view.",
			QueryPath: "bucket",
		},
	},
	Action:          handleEntitiesSynonymsRemove,
	HideHelpCommand: true,
}

func handleEntitiesSynonymsAdd(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
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

	params := bem.EntitySynonymAddParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Entities.Synonyms.Add(
		ctx,
		cmd.Value("id").(string),
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
		Title:          "entities:synonyms add",
		Transform:      transform,
	})
}

func handleEntitiesSynonymsRemove(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("synonym-id") && len(unusedArgs) > 0 {
		cmd.Set("synonym-id", unusedArgs[0])
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

	params := bem.EntitySynonymRemoveParams{
		ID: cmd.Value("id").(string),
	}

	return client.Entities.Synonyms.Remove(
		ctx,
		cmd.Value("synonym-id").(string),
		params,
		options...,
	)
}
