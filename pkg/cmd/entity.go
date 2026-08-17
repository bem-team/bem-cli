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

var entitiesUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update Entity",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "bucket",
			Usage:     "Optional bucket public ID (`bkt_...`) to scope the lookup to. Omit for the default bucket.",
			QueryPath: "bucket",
		},
		&requestflag.Flag[[]string]{
			Name:     "add-synonym",
			Usage:    "Surface forms to attach as `customer_defined` synonyms.",
			BodyPath: "addSynonyms",
		},
		&requestflag.Flag[string]{
			Name:     "assigned-type-id",
			Usage:    "The `ety_...` public ID of the type to assign (overriding the bem-inferred\ntype). The empty string clears the assignment. Omit to leave unchanged.",
			BodyPath: "assignedTypeID",
		},
		&requestflag.Flag[string]{
			Name:     "canonical",
			Usage:    "Replace the entity's canonical surface form (re-derives its normalized form).",
			BodyPath: "canonical",
		},
		&requestflag.Flag[string]{
			Name:     "locale",
			Usage:    "Optional BCP 47 locale tag stamped on any added synonyms.",
			BodyPath: "locale",
		},
		&requestflag.Flag[[]string]{
			Name:     "remove-synonym-id",
			Usage:    "`esn_...` synonym IDs to soft-delete. Only `customer_defined` /\n`sme_approved` synonyms may be removed; an `extracted` synonym is rejected\nwith `409`.",
			BodyPath: "removeSynonymIDs",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    "Transition the entity's curation status. Only `approved` or `rejected` are\naccepted, and only from `extracted` or `proposed` (any other transition is\nrejected with `409`).",
			BodyPath: "status",
		},
	},
	Action:          handleEntitiesUpdate,
	HideHelpCommand: true,
}

var entitiesBulkCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "bulk-create",
	Usage:   "Bulk Seed Entities",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "entity",
			Usage:    "The entities to seed. Must be non-empty.",
			Required: true,
			BodyPath: "entities",
		},
		&requestflag.Flag[string]{
			Name:     "bucket",
			Usage:    "Optional bucket public ID (`bkt_...`) to seed into. Omit to use the\naccount+environment default bucket.",
			BodyPath: "bucket",
		},
		&requestflag.Flag[string]{
			Name:     "on-conflict",
			Usage:    "Conflict strategy for an entity that already exists. Only `merge` is\nsupported and it is the default: synonyms are added additively, a longer\ndescription replaces the old one, and attributes are merged with new keys\nwinning.",
			BodyPath: "onConflict",
		},
	},
	Action:          handleEntitiesBulkCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"entity": {
		&requestflag.InnerFlag[string]{
			Name:       "entity.canonical",
			Usage:      "The canonical (longest / most descriptive) surface form for the entity,\ne.g. `Acme Corporation`. Required. Normalized (lowercased,\nwhitespace-folded) for the uniqueness key.",
			InnerField: "canonical",
		},
		&requestflag.InnerFlag[string]{
			Name:       "entity.type",
			Usage:      "The entity type name, e.g. `instrument` or `organization`. Required.\nResolved against your taxonomy and created if it does not yet exist.",
			InnerField: "type",
		},
		&requestflag.InnerFlag[any]{
			Name:       "entity.attributes",
			Usage:      "Optional per-entity structured attribute values, e.g.\n`{ \"manufacturer\": \"Acme\", \"dosageMg\": 50 }`. When the entity's type\ndeclares an attribute schema, keys not present in that schema cause the row\nto be rejected.",
			InnerField: "attributes",
		},
		&requestflag.InnerFlag[string]{
			Name:       "entity.description",
			Usage:      "Optional free-form description of the entity.",
			InnerField: "description",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "entity.synonyms",
			Usage:      "Optional additional surface forms to attach as `customer_defined` synonyms.",
			InnerField: "synonyms",
		},
	},
})

var entitiesBulkValidate = cli.Command{
	Name:    "bulk-validate",
	Usage:   "Bulk Validate Entities",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "entity-id",
			Usage:    "The `ent_...` IDs to transition. Must be non-empty.",
			Required: true,
			BodyPath: "entityIDs",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    "Terminal status to apply to every entity.",
			Required: true,
			BodyPath: "status",
		},
		&requestflag.Flag[string]{
			Name:      "bucket",
			Usage:     "Optional bucket public ID (`bkt_...`) to scope the lookup to. Omit for the default bucket.",
			QueryPath: "bucket",
		},
	},
	Action:          handleEntitiesBulkValidate,
	HideHelpCommand: true,
}

var entitiesRetrieveRelations = cli.Command{
	Name:    "retrieve-relations",
	Usage:   "Get an Entity's Relations",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "bucket",
			Usage:     "Optional bucket public ID (`bkt_...`) to scope the read to one bucket.\nOmit for the unscoped (all account+environment) view.",
			QueryPath: "bucket",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Cursor: return edges whose KSUID sorts after this value.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[string]{
			Name:      "direction",
			Usage:     "Which edges to return relative to the entity. Defaults to `both`.",
			QueryPath: "direction",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of edges to return (default 50, max 200).",
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "relation-type",
			Usage:     "Exact-match filter on the relation label.",
			QueryPath: "relationType",
		},
	},
	Action:          handleEntitiesRetrieveRelations,
	HideHelpCommand: true,
}

var entitiesRetrieveSeedStatus = cli.Command{
	Name:    "retrieve-seed-status",
	Usage:   "Get Seed Job Status",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleEntitiesRetrieveSeedStatus,
	HideHelpCommand: true,
}

func handleEntitiesUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := bem.EntityUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Entities.Update(
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
		Title:          "entities update",
		Transform:      transform,
	})
}

func handleEntitiesBulkCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := bem.EntityBulkNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Entities.BulkNew(ctx, params, options...)
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
		Title:          "entities bulk-create",
		Transform:      transform,
	})
}

func handleEntitiesBulkValidate(ctx context.Context, cmd *cli.Command) error {
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

	params := bem.EntityBulkValidateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Entities.BulkValidate(ctx, params, options...)
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
		Title:          "entities bulk-validate",
		Transform:      transform,
	})
}

func handleEntitiesRetrieveRelations(ctx context.Context, cmd *cli.Command) error {
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
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := bem.EntityGetRelationsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Entities.GetRelations(
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
		Title:          "entities retrieve-relations",
		Transform:      transform,
	})
}

func handleEntitiesRetrieveSeedStatus(ctx context.Context, cmd *cli.Command) error {
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
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Entities.GetSeedStatus(ctx, cmd.Value("id").(string), options...)
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
		Title:          "entities retrieve-seed-status",
		Transform:      transform,
	})
}
