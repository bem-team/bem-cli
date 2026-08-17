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

var subscriptionsCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a new subscription to listen to transform or error events.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of subscription.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "Type of subscription.",
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[string]{
			Name:     "collection-id",
			Usage:    "Unique identifier of collection this subscription listens to (alternative to collectionName).",
			BodyPath: "collectionID",
		},
		&requestflag.Flag[string]{
			Name:     "collection-name",
			Usage:    "Name of collection this subscription listens to (required for collection-based subscriptions).",
			BodyPath: "collectionName",
		},
		&requestflag.Flag[bool]{
			Name:     "disabled",
			Usage:    "Toggles whether subscription is active or not.",
			BodyPath: "disabled",
		},
		&requestflag.Flag[string]{
			Name:     "function-id",
			Usage:    "Unique identifier of function this subscription listens to (alternative to functionName).",
			BodyPath: "functionID",
		},
		&requestflag.Flag[string]{
			Name:     "function-name",
			Usage:    "Unique name of function this subscription listens to (required for function-based subscriptions).",
			BodyPath: "functionName",
		},
		&requestflag.Flag[string]{
			Name:     "google-drive-folder-id",
			Usage:    "Google Drive folder ID for syncing output data to Google Drive.",
			BodyPath: "googleDriveFolderID",
		},
		&requestflag.Flag[string]{
			Name:     "s3-bucket",
			Usage:    "S3 bucket name for syncing output data to AWS S3.",
			BodyPath: "s3Bucket",
		},
		&requestflag.Flag[string]{
			Name:     "s3-file-path",
			Usage:    "S3 file path for syncing output data to AWS S3.",
			BodyPath: "s3FilePath",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-url",
			Usage:    "URL bem will send webhook requests to.",
			BodyPath: "webhookURL",
		},
	},
	Action:          handleSubscriptionsCreate,
	HideHelpCommand: true,
}

var subscriptionsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get a Subscription",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "subscription-id",
			Required:  true,
			PathParam: "subscriptionID",
		},
	},
	Action:          handleSubscriptionsRetrieve,
	HideHelpCommand: true,
}

var subscriptionsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Updates an existing subscription. Follow conventional PATCH behavior, so only\nincluded fields will be updated.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "subscription-id",
			Required:  true,
			PathParam: "subscriptionID",
		},
		&requestflag.Flag[bool]{
			Name:     "disabled",
			Usage:    "Toggles whether subscription is active or not.",
			BodyPath: "disabled",
		},
		&requestflag.Flag[string]{
			Name:     "function-name",
			Usage:    "Unique name of function this subscription listens to.",
			BodyPath: "functionName",
		},
		&requestflag.Flag[string]{
			Name:     "google-drive-folder-id",
			Usage:    "Google Drive folder ID for syncing output data to Google Drive.",
			BodyPath: "googleDriveFolderID",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of subscription.",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "s3-bucket",
			Usage:    "S3 bucket name for syncing output data to AWS S3.",
			BodyPath: "s3Bucket",
		},
		&requestflag.Flag[string]{
			Name:     "s3-file-path",
			Usage:    "S3 file path for syncing output data to AWS S3.",
			BodyPath: "s3FilePath",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "Type of subscription.",
			BodyPath: "type",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-url",
			Usage:    "URL bem will send webhook requests to.",
			BodyPath: "webhookURL",
		},
	},
	Action:          handleSubscriptionsUpdate,
	HideHelpCommand: true,
}

var subscriptionsList = cli.Command{
	Name:    "list",
	Usage:   "List Subscriptions",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "ending-before",
			Usage:     "A cursor to use in pagination. `endingBefore` is a task ID that defines your place in the list. For example, if you make a list request and receive 50 objects, starting with `sub_2c9AXIj48cUYJtCuv1gsQtHGDzK`, your subsequent call can include `endingBefore=sub_2c9AXIj48cUYJtCuv1gsQtHGDzK` to fetch the previous page of the list.",
			QueryPath: "endingBefore",
		},
		&requestflag.Flag[[]string]{
			Name:      "function-name",
			Usage:     "Filters to subscriptions linked to included array of function names.",
			QueryPath: "functionNames",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "This specifies a limit on the number of objects to return, ranging between 1 and 100.",
			Default:   50,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "starting-after",
			Usage:     "A cursor to use in pagination. `startingAfter` is a task ID that defines your place in the list. For example, if you make a list request and receive 50 objects, ending with `sub_2c9AXIj48cUYJtCuv1gsQtHGDzK`, your subsequent call can include `startingAfter=sub_2c9AXIj48cUYJtCuv1gsQtHGDzK` to fetch the next page of the list.",
			QueryPath: "startingAfter",
		},
	},
	Action:          handleSubscriptionsList,
	HideHelpCommand: true,
}

var subscriptionsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes an existing subscription.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "subscription-id",
			Required:  true,
			PathParam: "subscriptionID",
		},
	},
	Action:          handleSubscriptionsDelete,
	HideHelpCommand: true,
}

func handleSubscriptionsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := bem.SubscriptionNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Subscriptions.New(ctx, params, options...)
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
		Title:          "subscriptions create",
		Transform:      transform,
	})
}

func handleSubscriptionsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("subscription-id") && len(unusedArgs) > 0 {
		cmd.Set("subscription-id", unusedArgs[0])
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
	_, err = client.Subscriptions.Get(ctx, cmd.Value("subscription-id").(string), options...)
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
		Title:          "subscriptions retrieve",
		Transform:      transform,
	})
}

func handleSubscriptionsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("subscription-id") && len(unusedArgs) > 0 {
		cmd.Set("subscription-id", unusedArgs[0])
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

	params := bem.SubscriptionUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Subscriptions.Update(
		ctx,
		cmd.Value("subscription-id").(string),
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
		Title:          "subscriptions update",
		Transform:      transform,
	})
}

func handleSubscriptionsList(ctx context.Context, cmd *cli.Command) error {
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

	params := bem.SubscriptionListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Subscriptions.List(ctx, params, options...)
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
		Title:          "subscriptions list",
		Transform:      transform,
	})
}

func handleSubscriptionsDelete(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("subscription-id") && len(unusedArgs) > 0 {
		cmd.Set("subscription-id", unusedArgs[0])
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

	return client.Subscriptions.Delete(ctx, cmd.Value("subscription-id").(string), options...)
}
