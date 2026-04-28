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

var fsNavigate = requestflag.WithInnerFlags(cli.Command{
	Name:    "navigate",
	Usage:   "**Navigate parsed documents and the cross-doc memory store via Unix-shell\nverbs.**",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "op",
			Usage:    "Operations exposed by `POST /v3/fs`.\n\nThe verbs and their flag names mirror Unix tools so an LLM agent's\nexisting vocabulary maps directly:\n\n- `ls` — list parsed documents\n- `cat` — read one parsed doc (optionally sliced by range / projected by select)\n- `grep` — substring or regex search across parse outputs\n- `head` — first N sections of one doc\n- `stat` — metadata only (page count, section count, parsed at, ...)\n- `find` — list canonical entities (cross-doc memory)\n- `open` — entity + mentions\n- `xref` — entity → sections across docs that mention it\n\nDoc-level ops (ls, cat, grep, head, stat) work on every parsed\ndocument, regardless of how the parse function was configured.\n\nMemory-level ops (find, open, xref) operate on the global entities\ntable which is only populated when the parse function had\n`linkAcrossDocuments: true`. On environments with no memory-linked\ndocs they return empty data with a hint pointing at the toggle.",
			Required: true,
			BodyPath: "op",
		},
		&requestflag.Flag[bool]{
			Name:     "count-only",
			Usage:    "When true, return only the hit count without snippet payload.\nCheaper than fetching matches when the agent only wants a yes/no.",
			BodyPath: "countOnly",
		},
		&requestflag.Flag[string]{
			Name:     "cursor",
			Usage:    "Pagination cursor. Pass the last item's ID from a previous response\n(`nextCursor`) to fetch the next page.",
			BodyPath: "cursor",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "filter",
			Usage:    "Filter options for `op=ls` and `op=find`.",
			BodyPath: "filter",
		},
		&requestflag.Flag[bool]{
			Name:     "ignore-case",
			Usage:    "When true (default), substring/regex matching is case-insensitive.",
			BodyPath: "ignoreCase",
		},
		&requestflag.Flag[int64]{
			Name:     "limit",
			Usage:    "Maximum results to return. Defaults vary per op (25–50).",
			BodyPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:     "n",
			Usage:    "First-N count for `op=head`. Defaults to 10.",
			BodyPath: "n",
		},
		&requestflag.Flag[string]{
			Name:     "path",
			Usage:    "Identifier for ops that operate on a single resource:\n- cat / head / stat: a parsed document, by `referenceID` or\n`transformationID`.\n- open / xref / stat: an entity, by `entityID`.",
			BodyPath: "path",
		},
		&requestflag.Flag[string]{
			Name:     "pattern",
			Usage:    "Substring or regex pattern for `op=grep`.",
			BodyPath: "pattern",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "range",
			Usage:    "Slice the parse output along page or section dimensions. Used with `op=cat`.",
			BodyPath: "range",
		},
		&requestflag.Flag[bool]{
			Name:     "regex",
			Usage:    "When true, `pattern` is interpreted as a Go regex. Default false.",
			BodyPath: "regex",
		},
		&requestflag.Flag[string]{
			Name:     "scope",
			Usage:    "Restricts grep to one part of the parse output. One of\n`\"sections\"`, `\"entities\"`, `\"relationships\"`, `\"all\"` (default).",
			BodyPath: "scope",
		},
		&requestflag.Flag[[]string]{
			Name:     "select",
			Usage:    "Project the parse output to specific dotted paths\n(e.g. `[\"sections.label\", \"sections.page\"]`), letting an agent map\na doc's structure cheaply before reading content. Used with\n`op=cat`.",
			BodyPath: "select",
		},
	},
	Action:          handleFsNavigate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"filter": {
		&requestflag.InnerFlag[string]{
			Name:       "filter.function-name",
			Usage:      "Match a parsed doc's source function name exactly.",
			InnerField: "functionName",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.search",
			Usage:      "Substring match on canonical name (entities) or `referenceID`\n(parsed docs). Case-insensitive.",
			InnerField: "search",
		},
		&requestflag.InnerFlag[any]{
			Name:       "filter.since",
			Usage:      "Restrict to resources created at or after this timestamp.",
			InnerField: "since",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.type",
			Usage:      "Match an entity's `type` field exactly (e.g. `\"drug\"`, `\"study\"`).",
			InnerField: "type",
		},
	},
	"range": {
		&requestflag.InnerFlag[int64]{
			Name:       "range.page",
			Usage:      "Restrict sections to one page (1-indexed).",
			InnerField: "page",
		},
		&requestflag.InnerFlag[[]int64]{
			Name:       "range.page-range",
			Usage:      "Restrict sections to an inclusive page range. Two-element array of\n`[from, to]` (both 1-indexed).",
			InnerField: "pageRange",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "range.section-types",
			Usage:      "Keep only sections whose `type` matches one of these (e.g.\n`[\"table\", \"list\"]`).",
			InnerField: "sectionTypes",
		},
	},
})

func handleFsNavigate(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := bem.FNavigateParams{}

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
	_, err = client.Fs.Navigate(ctx, params, options...)
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
		Title:          "fs navigate",
		Transform:      transform,
	})
}
