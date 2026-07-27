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

var functionsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "**Create a function.**",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "function-name",
			Usage:    "Name of function. Must be UNIQUE on a per-environment basis.",
			Required: true,
			BodyPath: "functionName",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    `Allowed values: "extract".`,
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[string]{
			Name:     "display-name",
			Usage:    "Display name of function. Human-readable name to help you identify the function.",
			BodyPath: "displayName",
		},
		&requestflag.Flag[bool]{
			Name:     "enable-bounding-boxes",
			Usage:    "Whether bounding box extraction is enabled. Applies to vision input types\n(pdf, png, jpeg, heic, heif, webp) that dispatch through the analyze path.\nWhen true, the function returns the document regions (page, coordinates) from which each\nfield was extracted. Enabling this automatically configures the function to use the bounding\nbox model. Disabling resets to the default.",
			BodyPath: "enableBoundingBoxes",
		},
		&requestflag.Flag[any]{
			Name:     "output-schema",
			Usage:    "Desired output structure defined in standard JSON Schema convention.",
			BodyPath: "outputSchema",
		},
		&requestflag.Flag[string]{
			Name:     "output-schema-name",
			Usage:    "Name of output schema object.",
			BodyPath: "outputSchemaName",
		},
		&requestflag.Flag[bool]{
			Name:     "pre-count",
			Usage:    "Reducing the risk of the model stopping early on long documents.\nTrade-off: Increases total latency. Compatible with\n`enableBoundingBoxes`.",
			BodyPath: "preCount",
		},
		&requestflag.Flag[bool]{
			Name:     "tabular-chunking-enabled",
			Usage:    "Whether tabular chunking is enabled. When true, tables in CSV/Excel files are processed\nin row batches rather than all at once.",
			BodyPath: "tabularChunkingEnabled",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Array of tags to categorize and organize functions.",
			BodyPath: "tags",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "classification",
			Usage:    "List of classifications a classify function can produce. Shares the underlying route list shape.",
			BodyPath: "classifications",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "Description of classifier. Can be used to provide additional context on classifier's purpose and expected inputs.",
			BodyPath: "description",
		},
		&requestflag.Flag[bool]{
			Name:     "native-visual-input",
			Usage:    "When true, image and PDF inputs are sent directly to the model for\nrouting instead of being OCR'd to text first. Defaults to true for new\nclassify functions and false for the legacy route type.",
			BodyPath: "nativeVisualInput",
		},
		&requestflag.Flag[string]{
			Name:     "destination-type",
			Usage:    "Destination type for a Send function.",
			BodyPath: "destinationType",
		},
		&requestflag.Flag[string]{
			Name:     "google-drive-folder-id",
			Usage:    "Google Drive folder ID. Required when destinationType is google_drive. Managed via Paragon OAuth.",
			BodyPath: "googleDriveFolderId",
		},
		&requestflag.Flag[string]{
			Name:     "s3-bucket",
			Usage:    "S3 bucket to upload the payload to. Required when destinationType is s3.",
			BodyPath: "s3Bucket",
		},
		&requestflag.Flag[string]{
			Name:     "s3-prefix",
			Usage:    "Optional S3 key prefix (folder path).",
			BodyPath: "s3Prefix",
		},
		&requestflag.Flag[bool]{
			Name:     "webhook-signing-enabled",
			Usage:    "Whether to sign webhook deliveries with an HMAC-SHA256 `bem-signature` header.\nDefaults to `true` when omitted — signing is on by default for new send functions.\nSet explicitly to `false` to disable.",
			BodyPath: "webhookSigningEnabled",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-url",
			Usage:    "Webhook URL to POST the payload to. Required when destinationType is webhook.",
			BodyPath: "webhookUrl",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "print-page-split-config",
			BodyPath: "printPageSplitConfig",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "semantic-page-split-config",
			BodyPath: "semanticPageSplitConfig",
		},
		&requestflag.Flag[string]{
			Name:     "split-type",
			Usage:    `Allowed values: "print_page", "semantic_page".`,
			BodyPath: "splitType",
		},
		&requestflag.Flag[string]{
			Name:     "join-type",
			Usage:    "The type of join to perform.",
			BodyPath: "joinType",
		},
		&requestflag.Flag[string]{
			Name:     "shaping-schema",
			Usage:    "JMESPath expression that defines how to transform and customize the input payload structure.\nPayload shaping allows you to extract, reshape, and reorganize data from complex input payloads\ninto a simplified, standardized output format. Use JMESPath syntax to select specific fields,\nperform calculations, and create new data structures tailored to your needs.",
			BodyPath: "shapingSchema",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "config",
			Usage:    "Configuration for an enrich function.\n\n**How Enrich Functions Work:**\n\nEnrich functions augment JSON input with data from external sources. They take JSON input\n(typically from a previous function), extract specified fields, fetch or search for matching\ndata, and inject the results back into the JSON.\n\n**Data Sources:**\n- **Collections** (`source: \"collection\"`): Vector/keyword search against a BEM collection.\nBest for semantic matching against pre-indexed documents.\n- **Endpoints** (`source: \"endpoint\"`): HTTP call to any user-provided REST API.\nBest for looking up live data from CRMs, ERPs, or other external systems.\nOptionally uses LLM agent reasoning to rank candidates returned by the endpoint.\n\n**Input Requirements:**\n- Must receive JSON input (typically from a previous function's output)\n\n**Example Use Cases:**\n- Match product descriptions to SKU codes from a product catalog collection\n- Enrich customer data with account details from a CRM endpoint\n- Use LLM agent reasoning to fuzzy-match line item descriptions to catalog products\n\n**Configuration:**\n- Define named endpoints (for endpoint-source steps)\n- Define one or more enrichment steps; steps are executed sequentially",
			BodyPath: "config",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "extra-config",
			Usage:    "Cross-cutting toggles for Parse functions. Mirrors the `extraConfig`\nsurface on Extract / Join — separated from `parseConfig` so the per-call\nParse output shape stays distinct from operator-level execution flags.",
			BodyPath: "extraConfig",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "parse-config",
			Usage:    "Per-version configuration for a Parse function.\n\nParse renders document pages (PDF, image) via vision LLM and emits\nstructured JSON. The two toggles below independently control entity\nextraction (a per-call output concern) and cross-document memory\nlinking (an environment-wide concern).",
			BodyPath: "parseConfig",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "render-config",
			Usage:    "Request-side render configuration. Carries the template document as\nbase64-encoded `.docx` bytes: the server validates them, stores the template,\nand derives the placeholder/style-id contract at create/update time, so\nclients never submit `placeholders` or `styleIds`. The response shape\n(`RenderConfig`) returns the derived contract.",
			BodyPath: "renderConfig",
		},
	},
	Action:          handleFunctionsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"classification": {
		&requestflag.InnerFlag[string]{
			Name:       "classification.name",
			InnerField: "name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "classification.description",
			InnerField: "description",
		},
		&requestflag.InnerFlag[string]{
			Name:       "classification.function-id",
			InnerField: "functionID",
		},
		&requestflag.InnerFlag[string]{
			Name:       "classification.function-name",
			InnerField: "functionName",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "classification.is-error-fallback",
			InnerField: "isErrorFallback",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "classification.origin",
			InnerField: "origin",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "classification.regex",
			InnerField: "regex",
		},
	},
	"print-page-split-config": {
		&requestflag.InnerFlag[string]{
			Name:       "print-page-split-config.next-function-id",
			InnerField: "nextFunctionID",
		},
		&requestflag.InnerFlag[string]{
			Name:       "print-page-split-config.next-function-name",
			InnerField: "nextFunctionName",
		},
	},
	"semantic-page-split-config": {
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "semantic-page-split-config.item-classes",
			InnerField: "itemClasses",
		},
	},
	"config": {
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "config.steps",
			Usage:      "Array of enrichment steps to execute sequentially.",
			InnerField: "steps",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "config.endpoints",
			Usage:      "Named HTTP endpoints available to endpoint-source steps.\nEach endpoint must have a unique `name` referenced by the step's `endpointName`.\nRequired when any step uses `source: \"endpoint\"`.",
			InnerField: "endpoints",
		},
	},
	"extra-config": {
		&requestflag.InnerFlag[bool]{
			Name:       "extra-config.enable-bounding-boxes",
			Usage:      "When true, return per-section and per-entity-mention coordinates in\nthe parse event's `fieldBoundingBoxes` map (same shape as Extract:\nJSON Pointer key → array of `{page, left, top, width, height}` with\ncoordinates normalized to [0, 1]). Keys are `/sections/{N}` and\n`/entities/{N}/occurrences/{M}` into the parse output. Only applies\nto the open-ended discovery path (no `schema`) and to vision input\ntypes. Bedrock-backed parse functions silently return an empty map\n(no native bbox support). Defaults to false.",
			InnerField: "enableBoundingBoxes",
		},
	},
	"parse-config": {
		&requestflag.InnerFlag[string]{
			Name:       "parse-config.default-bucket",
			Usage:      "Optional bucket NAME that parse-extracted entities land in when no\ncall-level bucket is supplied. Lower precedence than a call-level bucket,\nhigher than the account+environment default.",
			InnerField: "defaultBucket",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "parse-config.extract-entities",
			Usage:      "When true, extract named entities (people, organizations, products,\nstudies, identifiers, etc.) and the relationships between them, and\ndedupe by canonical name within the document. When false, only\n`sections[]` is extracted; `entities[]` and `relationships[]` come\nback empty in the parse output. Defaults to true.",
			InnerField: "extractEntities",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "parse-config.link-across-documents",
			Usage:      "When true, link this document's entities to entities seen in earlier\ndocuments in this environment, building one canonical record per\nreal-world thing across the corpus. Visible in the Memory tab and\nqueryable via `POST /v3/fs` (op=find / open / xref). Doesn't change\nthis call's parse output. Requires `extractEntities=true`. Defaults\nto true.",
			InnerField: "linkAcrossDocuments",
		},
		&requestflag.InnerFlag[any]{
			Name:       "parse-config.schema",
			Usage:      "Optional JSONSchema. When provided, each chunk performs schema-guided\nextraction. When absent, chunks perform open-ended discovery and\nreturn sections, entities, and relationships per the discovery\nschema.",
			InnerField: "schema",
		},
	},
	"render-config": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "render-config.template",
			InnerField: "template",
		},
	},
})

var functionsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "**Retrieve a function's current version by name.**",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "function-name",
			Required:  true,
			PathParam: "functionName",
		},
	},
	Action:          handleFunctionsRetrieve,
	HideHelpCommand: true,
}

var functionsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "**Update a function. Updates create a new version.**",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "path-function-name",
			Required:  true,
			PathParam: "functionName",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    `Allowed values: "extract".`,
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[string]{
			Name:     "display-name",
			Usage:    "Display name of function. Human-readable name to help you identify the function.",
			BodyPath: "displayName",
		},
		&requestflag.Flag[bool]{
			Name:     "enable-bounding-boxes",
			Usage:    "Whether bounding box extraction is enabled. Applies to vision input types\n(pdf, png, jpeg, heic, heif, webp) that dispatch through the analyze path.\nWhen true, the function returns the document regions (page, coordinates) from which each\nfield was extracted. Enabling this automatically configures the function to use the bounding\nbox model. Disabling resets to the default.",
			BodyPath: "enableBoundingBoxes",
		},
		&requestflag.Flag[string]{
			Name:     "function-name",
			Usage:    "Name of function. Must be UNIQUE on a per-environment basis.",
			BodyPath: "functionName",
		},
		&requestflag.Flag[any]{
			Name:     "output-schema",
			Usage:    "Desired output structure defined in standard JSON Schema convention.",
			BodyPath: "outputSchema",
		},
		&requestflag.Flag[string]{
			Name:     "output-schema-name",
			Usage:    "Name of output schema object.",
			BodyPath: "outputSchemaName",
		},
		&requestflag.Flag[bool]{
			Name:     "pre-count",
			Usage:    "Reducing the risk of the model stopping early on long documents.\nTrade-off: Increases total latency. Compatible with\n`enableBoundingBoxes`.",
			BodyPath: "preCount",
		},
		&requestflag.Flag[bool]{
			Name:     "tabular-chunking-enabled",
			Usage:    "Whether tabular chunking is enabled. When true, tables in CSV/Excel files are processed\nin row batches rather than all at once.",
			BodyPath: "tabularChunkingEnabled",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Array of tags to categorize and organize functions.",
			BodyPath: "tags",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "classification",
			Usage:    "List of classifications a classify function can produce. Shares the underlying route list shape.",
			BodyPath: "classifications",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "Description of classifier. Can be used to provide additional context on classifier's purpose and expected inputs.",
			BodyPath: "description",
		},
		&requestflag.Flag[bool]{
			Name:     "native-visual-input",
			Usage:    "When true, image and PDF inputs are sent directly to the model for\nrouting instead of being OCR'd to text first. Defaults to true for new\nclassify functions and false for the legacy route type.",
			BodyPath: "nativeVisualInput",
		},
		&requestflag.Flag[string]{
			Name:     "destination-type",
			Usage:    "Destination type for a Send function.",
			BodyPath: "destinationType",
		},
		&requestflag.Flag[string]{
			Name:     "google-drive-folder-id",
			Usage:    "Google Drive folder ID. Required when destinationType is google_drive. Managed via Paragon OAuth.",
			BodyPath: "googleDriveFolderId",
		},
		&requestflag.Flag[string]{
			Name:     "s3-bucket",
			Usage:    "S3 bucket to upload the payload to. Required when destinationType is s3.",
			BodyPath: "s3Bucket",
		},
		&requestflag.Flag[string]{
			Name:     "s3-prefix",
			Usage:    "Optional S3 key prefix (folder path).",
			BodyPath: "s3Prefix",
		},
		&requestflag.Flag[bool]{
			Name:     "webhook-signing-enabled",
			Usage:    "Whether to sign webhook deliveries with an HMAC-SHA256 `bem-signature` header.\nDefaults to `true` when omitted — signing is on by default for new send functions.\nSet explicitly to `false` to disable.",
			BodyPath: "webhookSigningEnabled",
		},
		&requestflag.Flag[string]{
			Name:     "webhook-url",
			Usage:    "Webhook URL to POST the payload to. Required when destinationType is webhook.",
			BodyPath: "webhookUrl",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "print-page-split-config",
			BodyPath: "printPageSplitConfig",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "semantic-page-split-config",
			BodyPath: "semanticPageSplitConfig",
		},
		&requestflag.Flag[string]{
			Name:     "split-type",
			Usage:    `Allowed values: "print_page", "semantic_page".`,
			BodyPath: "splitType",
		},
		&requestflag.Flag[string]{
			Name:     "join-type",
			Usage:    "The type of join to perform.",
			BodyPath: "joinType",
		},
		&requestflag.Flag[string]{
			Name:     "shaping-schema",
			Usage:    "JMESPath expression that defines how to transform and customize the input payload structure.\nPayload shaping allows you to extract, reshape, and reorganize data from complex input payloads\ninto a simplified, standardized output format. Use JMESPath syntax to select specific fields,\nperform calculations, and create new data structures tailored to your needs.",
			BodyPath: "shapingSchema",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "config",
			Usage:    "Configuration for an enrich function.\n\n**How Enrich Functions Work:**\n\nEnrich functions augment JSON input with data from external sources. They take JSON input\n(typically from a previous function), extract specified fields, fetch or search for matching\ndata, and inject the results back into the JSON.\n\n**Data Sources:**\n- **Collections** (`source: \"collection\"`): Vector/keyword search against a BEM collection.\nBest for semantic matching against pre-indexed documents.\n- **Endpoints** (`source: \"endpoint\"`): HTTP call to any user-provided REST API.\nBest for looking up live data from CRMs, ERPs, or other external systems.\nOptionally uses LLM agent reasoning to rank candidates returned by the endpoint.\n\n**Input Requirements:**\n- Must receive JSON input (typically from a previous function's output)\n\n**Example Use Cases:**\n- Match product descriptions to SKU codes from a product catalog collection\n- Enrich customer data with account details from a CRM endpoint\n- Use LLM agent reasoning to fuzzy-match line item descriptions to catalog products\n\n**Configuration:**\n- Define named endpoints (for endpoint-source steps)\n- Define one or more enrichment steps; steps are executed sequentially",
			BodyPath: "config",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "extra-config",
			Usage:    "Cross-cutting toggles for Parse functions. Mirrors the `extraConfig`\nsurface on Extract / Join — separated from `parseConfig` so the per-call\nParse output shape stays distinct from operator-level execution flags.",
			BodyPath: "extraConfig",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "parse-config",
			Usage:    "Per-version configuration for a Parse function.\n\nParse renders document pages (PDF, image) via vision LLM and emits\nstructured JSON. The two toggles below independently control entity\nextraction (a per-call output concern) and cross-document memory\nlinking (an environment-wide concern).",
			BodyPath: "parseConfig",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "render-config",
			Usage:    "Request-side render configuration. Carries the template document as\nbase64-encoded `.docx` bytes: the server validates them, stores the template,\nand derives the placeholder/style-id contract at create/update time, so\nclients never submit `placeholders` or `styleIds`. The response shape\n(`RenderConfig`) returns the derived contract.",
			BodyPath: "renderConfig",
		},
	},
	Action:          handleFunctionsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"classification": {
		&requestflag.InnerFlag[string]{
			Name:       "classification.name",
			InnerField: "name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "classification.description",
			InnerField: "description",
		},
		&requestflag.InnerFlag[string]{
			Name:       "classification.function-id",
			InnerField: "functionID",
		},
		&requestflag.InnerFlag[string]{
			Name:       "classification.function-name",
			InnerField: "functionName",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "classification.is-error-fallback",
			InnerField: "isErrorFallback",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "classification.origin",
			InnerField: "origin",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "classification.regex",
			InnerField: "regex",
		},
	},
	"print-page-split-config": {
		&requestflag.InnerFlag[string]{
			Name:       "print-page-split-config.next-function-id",
			InnerField: "nextFunctionID",
		},
		&requestflag.InnerFlag[string]{
			Name:       "print-page-split-config.next-function-name",
			InnerField: "nextFunctionName",
		},
	},
	"semantic-page-split-config": {
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "semantic-page-split-config.item-classes",
			InnerField: "itemClasses",
		},
	},
	"config": {
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "config.steps",
			Usage:      "Array of enrichment steps to execute sequentially.",
			InnerField: "steps",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "config.endpoints",
			Usage:      "Named HTTP endpoints available to endpoint-source steps.\nEach endpoint must have a unique `name` referenced by the step's `endpointName`.\nRequired when any step uses `source: \"endpoint\"`.",
			InnerField: "endpoints",
		},
	},
	"extra-config": {
		&requestflag.InnerFlag[bool]{
			Name:       "extra-config.enable-bounding-boxes",
			Usage:      "When true, return per-section and per-entity-mention coordinates in\nthe parse event's `fieldBoundingBoxes` map (same shape as Extract:\nJSON Pointer key → array of `{page, left, top, width, height}` with\ncoordinates normalized to [0, 1]). Keys are `/sections/{N}` and\n`/entities/{N}/occurrences/{M}` into the parse output. Only applies\nto the open-ended discovery path (no `schema`) and to vision input\ntypes. Bedrock-backed parse functions silently return an empty map\n(no native bbox support). Defaults to false.",
			InnerField: "enableBoundingBoxes",
		},
	},
	"parse-config": {
		&requestflag.InnerFlag[string]{
			Name:       "parse-config.default-bucket",
			Usage:      "Optional bucket NAME that parse-extracted entities land in when no\ncall-level bucket is supplied. Lower precedence than a call-level bucket,\nhigher than the account+environment default.",
			InnerField: "defaultBucket",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "parse-config.extract-entities",
			Usage:      "When true, extract named entities (people, organizations, products,\nstudies, identifiers, etc.) and the relationships between them, and\ndedupe by canonical name within the document. When false, only\n`sections[]` is extracted; `entities[]` and `relationships[]` come\nback empty in the parse output. Defaults to true.",
			InnerField: "extractEntities",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "parse-config.link-across-documents",
			Usage:      "When true, link this document's entities to entities seen in earlier\ndocuments in this environment, building one canonical record per\nreal-world thing across the corpus. Visible in the Memory tab and\nqueryable via `POST /v3/fs` (op=find / open / xref). Doesn't change\nthis call's parse output. Requires `extractEntities=true`. Defaults\nto true.",
			InnerField: "linkAcrossDocuments",
		},
		&requestflag.InnerFlag[any]{
			Name:       "parse-config.schema",
			Usage:      "Optional JSONSchema. When provided, each chunk performs schema-guided\nextraction. When absent, chunks perform open-ended discovery and\nreturn sections, entities, and relationships per the discovery\nschema.",
			InnerField: "schema",
		},
	},
	"render-config": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "render-config.template",
			InnerField: "template",
		},
	},
})

var functionsList = cli.Command{
	Name:    "list",
	Usage:   "**List functions in the current environment.**",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "display-name",
			QueryPath: "displayName",
		},
		&requestflag.Flag[string]{
			Name:      "ending-before",
			QueryPath: "endingBefore",
		},
		&requestflag.Flag[[]string]{
			Name:      "function-id",
			QueryPath: "functionIDs",
		},
		&requestflag.Flag[[]string]{
			Name:      "function-name",
			QueryPath: "functionNames",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Default:   50,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "sort-order",
			Usage:     `Allowed values: "asc", "desc".`,
			Default:   "asc",
			QueryPath: "sortOrder",
		},
		&requestflag.Flag[string]{
			Name:      "starting-after",
			QueryPath: "startingAfter",
		},
		&requestflag.Flag[[]string]{
			Name:      "tag",
			QueryPath: "tags",
		},
		&requestflag.Flag[[]string]{
			Name:      "type",
			QueryPath: "types",
		},
		&requestflag.Flag[[]string]{
			Name:      "workflow-id",
			QueryPath: "workflowIDs",
		},
		&requestflag.Flag[[]string]{
			Name:      "workflow-name",
			QueryPath: "workflowNames",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleFunctionsList,
	HideHelpCommand: true,
}

var functionsDelete = cli.Command{
	Name:    "delete",
	Usage:   "**Delete a function and every one of its versions.**",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "function-name",
			Required:  true,
			PathParam: "functionName",
		},
	},
	Action:          handleFunctionsDelete,
	HideHelpCommand: true,
}

var functionsCompareMetrics = cli.Command{
	Name:    "compare-metrics",
	Usage:   "**Compare metrics between two function versions.**",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "function-name",
			Usage:    "Name of the function to compare versions for",
			Required: true,
			BodyPath: "functionName",
		},
		&requestflag.Flag[int64]{
			Name:     "baseline-version-num",
			Usage:    "**Baseline version number for comparison**\n\nIf not provided, defaults to the previous version (current - 1).",
			BodyPath: "baselineVersionNum",
		},
		&requestflag.Flag[int64]{
			Name:     "comparison-version-num",
			Usage:    "**Comparison version number**\n\nIf not provided, defaults to the current version.",
			BodyPath: "comparisonVersionNum",
		},
		&requestflag.Flag[bool]{
			Name:     "is-regression",
			Usage:    "**Whether to compare regression test data only**\n\nIf true, only compares transformations marked as regression tests.",
			BodyPath: "isRegression",
		},
	},
	Action:          handleFunctionsCompareMetrics,
	HideHelpCommand: true,
}

var functionsEstimateReviewRequirements = cli.Command{
	Name:    "estimate-review-requirements",
	Usage:   "**Estimate human review requirements for a function.**",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "function-name",
			Usage:    "Name of the function to analyze",
			Required: true,
			BodyPath: "functionName",
		},
		&requestflag.Flag[[]int64]{
			Name:     "confidence-level",
			Usage:    "Confidence levels for statistical analysis as integers representing percentages (e.g., [90, 95, 99] for 90%, 95%, 99%). IMPORTANT: Only integers are accepted, floats like 0.95 will be rejected.",
			Default:  []int64{95},
			BodyPath: "confidenceLevels",
		},
		&requestflag.Flag[string]{
			Name:     "confidence-method",
			Usage:    "Confidence interval calculation method (default \"wald\").\n\n- \"wald\": Normal approximation method (faster, standard)\n- \"wilson\": Wilson score interval (more robust for extreme rates)",
			Default:  "wald",
			BodyPath: "confidenceMethod",
		},
		&requestflag.Flag[string]{
			Name:     "evaluation-version",
			Usage:    `Optional evaluation version to filter evaluations by. Must be one of the supported versions. If not provided, defaults to "0.1.0-gemini".`,
			Default:  "0.1.0-gemini",
			BodyPath: "evaluationVersion",
		},
		&requestflag.Flag[int64]{
			Name:     "function-version-num",
			Usage:    "Optional function version number to analyze. If not provided, uses the latest/current version of the function.",
			BodyPath: "functionVersionNum",
		},
		&requestflag.Flag[bool]{
			Name:     "is-regression",
			Usage:    "Internal flag indicating if the request is from a regression test",
			BodyPath: "isRegression",
		},
		&requestflag.Flag[float64]{
			Name:     "margin-of-error",
			Usage:    "Margin of error for statistical calculations",
			Default:  0.05,
			BodyPath: "marginOfError",
		},
		&requestflag.Flag[float64]{
			Name:     "threshold-max",
			Usage:    "Maximum confidence threshold to analyze",
			Default:  1,
			BodyPath: "thresholdMax",
		},
		&requestflag.Flag[float64]{
			Name:     "threshold-min",
			Usage:    "Minimum confidence threshold to analyze",
			Default:  0.5,
			BodyPath: "thresholdMin",
		},
		&requestflag.Flag[float64]{
			Name:     "threshold-step",
			Usage:    "Step size for threshold analysis (smaller = more granular)",
			Default:  0.01,
			BodyPath: "thresholdStep",
		},
	},
	Action:          handleFunctionsEstimateReviewRequirements,
	HideHelpCommand: true,
}

var functionsGetMetrics = cli.Command{
	Name:    "get-metrics",
	Usage:   "**Retrieve performance metrics for functions based on labeled transformation\ndata.**",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "ending-before",
			Usage:     "Cursor — a `functionID` defining your place in the list.",
			QueryPath: "endingBefore",
		},
		&requestflag.Flag[[]string]{
			Name:      "function-id",
			QueryPath: "functionIDs",
		},
		&requestflag.Flag[[]string]{
			Name:      "function-name",
			QueryPath: "functionNames",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Default:   50,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "sort-order",
			Usage:     "Sort direction over the result set (default `asc`). Pagination works\nsymmetrically in both directions via `startingAfter` / `endingBefore`.",
			Default:   "asc",
			QueryPath: "sortOrder",
		},
		&requestflag.Flag[string]{
			Name:      "starting-after",
			Usage:     "Cursor — a `functionID` defining your place in the list.",
			QueryPath: "startingAfter",
		},
		&requestflag.Flag[[]string]{
			Name:      "type",
			QueryPath: "types",
		},
	},
	Action:          handleFunctionsGetMetrics,
	HideHelpCommand: true,
}

func handleFunctionsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := bem.FunctionNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Functions.New(ctx, params, options...)
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
		Title:          "functions create",
		Transform:      transform,
	})
}

func handleFunctionsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("function-name") && len(unusedArgs) > 0 {
		cmd.Set("function-name", unusedArgs[0])
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
	_, err = client.Functions.Get(ctx, cmd.Value("function-name").(string), options...)
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
		Title:          "functions retrieve",
		Transform:      transform,
	})
}

func handleFunctionsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("path-function-name") && len(unusedArgs) > 0 {
		cmd.Set("path-function-name", unusedArgs[0])
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

	params := bem.FunctionUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Functions.Update(
		ctx,
		cmd.Value("path-function-name").(string),
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
		Title:          "functions update",
		Transform:      transform,
	})
}

func handleFunctionsList(ctx context.Context, cmd *cli.Command) error {
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

	params := bem.FunctionListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Functions.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "functions list",
			Transform:      transform,
		})
	} else {
		iter := client.Functions.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "functions list",
			Transform:      transform,
		})
	}
}

func handleFunctionsDelete(ctx context.Context, cmd *cli.Command) error {
	client := bem.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("function-name") && len(unusedArgs) > 0 {
		cmd.Set("function-name", unusedArgs[0])
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

	return client.Functions.Delete(ctx, cmd.Value("function-name").(string), options...)
}

func handleFunctionsCompareMetrics(ctx context.Context, cmd *cli.Command) error {
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

	params := bem.FunctionCompareMetricsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Functions.CompareMetrics(ctx, params, options...)
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
		Title:          "functions compare-metrics",
		Transform:      transform,
	})
}

func handleFunctionsEstimateReviewRequirements(ctx context.Context, cmd *cli.Command) error {
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

	params := bem.FunctionEstimateReviewRequirementsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Functions.EstimateReviewRequirements(ctx, params, options...)
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
		Title:          "functions estimate-review-requirements",
		Transform:      transform,
	})
}

func handleFunctionsGetMetrics(ctx context.Context, cmd *cli.Command) error {
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

	params := bem.FunctionGetMetricsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Functions.GetMetrics(ctx, params, options...)
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
		Title:          "functions get-metrics",
		Transform:      transform,
	})
}
