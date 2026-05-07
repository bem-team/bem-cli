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

var connectorsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a Connector",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Human-friendly name for this connector.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "Connector type.",
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[string]{
			Name:     "box-client-id",
			Usage:    "Box client ID (from your Box application).",
			BodyPath: "boxClientID",
		},
		&requestflag.Flag[string]{
			Name:     "box-client-secret",
			Usage:    "Box client secret (from your Box application).",
			BodyPath: "boxClientSecret",
		},
		&requestflag.Flag[string]{
			Name:     "box-enterprise-id",
			Usage:    "Box enterprise ID.",
			BodyPath: "boxEnterpriseID",
		},
		&requestflag.Flag[string]{
			Name:     "box-folder-id",
			Usage:    "Box folder ID to watch for new uploads.",
			BodyPath: "boxFolderID",
		},
		&requestflag.Flag[any]{
			Name:     "paragon-configuration",
			Usage:    "Configuration specific to the type of integration.",
			BodyPath: "paragonConfiguration",
		},
		&requestflag.Flag[string]{
			Name:     "paragon-integration",
			Usage:    `Paragon integration, eg. "googledrive".`,
			BodyPath: "paragonIntegration",
		},
		&requestflag.Flag[string]{
			Name:     "workflow-id",
			Usage:    "One of `workflowID` or `workflowName` must be provided.\n\nIf both are provided, they must refer to the same workflow.",
			BodyPath: "workflowID",
		},
		&requestflag.Flag[string]{
			Name:     "workflow-name",
			Usage:    "One of `workflowID` or `workflowName` must be provided.\n\nIf both are provided, they must refer to the same workflow.",
			BodyPath: "workflowName",
		},
	},
	Action:          handleConnectorsCreate,
	HideHelpCommand: true,
}

var connectorsList = cli.Command{
	Name:    "list",
	Usage:   "List Connectors",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "workflow-id",
			Usage:     "Filter connectors by workflow API ID (e.g. `wf_...`).\n\nIf both `workflowID` and `workflowName` are provided, results must match both.",
			QueryPath: "workflowID",
		},
		&requestflag.Flag[string]{
			Name:      "workflow-name",
			Usage:     "Filter connectors by workflow name (exact match).\n\nIf both `workflowID` and `workflowName` are provided, results must match both.",
			QueryPath: "workflowName",
		},
	},
	Action:          handleConnectorsList,
	HideHelpCommand: true,
}

func handleConnectorsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := bem.ConnectorNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Connectors.New(ctx, params, options...)
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
		Title:          "connectors create",
		Transform:      transform,
	})
}

func handleConnectorsList(ctx context.Context, cmd *cli.Command) error {
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

	params := bem.ConnectorListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Connectors.List(ctx, params, options...)
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
		Title:          "connectors list",
		Transform:      transform,
	})
}
