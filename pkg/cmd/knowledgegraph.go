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

var knowledgeGraphRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve the Knowledge Graph",
	Suggest: true,
	Flags: []cli.Flag{
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
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of edges per page (default 50, max 200).",
			QueryPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:      "max-depth",
			Usage:     "Maximum hops from the center node. Only meaningful with `nodeID`. Defaults\nto 2 and is clamped down to a system maximum (5).",
			QueryPath: "maxDepth",
		},
		&requestflag.Flag[string]{
			Name:      "node-id",
			Usage:     "Center the graph on this entity (`ent_...`) and only return the subgraph\nwithin `maxDepth` hops of it; every node then carries its `depth` (hops\nfrom the center, center = 0). Omit for the uncentered whole-graph view.\n`rootNodeID` and `focusNodeID` are accepted as aliases.",
			QueryPath: "nodeID",
		},
		&requestflag.Flag[string]{
			Name:      "search",
			Usage:     "Case-insensitive substring match on canonical names. Both endpoints of an\nedge must match for the edge (and its nodes) to be returned.",
			QueryPath: "search",
		},
		&requestflag.Flag[any]{
			Name:      "since",
			Usage:     "Only edges created at/after this RFC 3339 timestamp.",
			QueryPath: "since",
		},
		&requestflag.Flag[[]string]{
			Name:      "type",
			Usage:     "Restrict to entities of these types. An edge is returned only when BOTH\nof its endpoints survive the type filter.",
			QueryPath: "type",
		},
	},
	Action:          handleKnowledgeGraphRetrieve,
	HideHelpCommand: true,
}

func handleKnowledgeGraphRetrieve(ctx context.Context, cmd *cli.Command) error {
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

	params := bem.KnowledgeGraphGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.KnowledgeGraph.Get(ctx, params, options...)
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
		Title:          "knowledge-graph retrieve",
		Transform:      transform,
	})
}
