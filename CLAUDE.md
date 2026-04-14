# Bem CLI

CLI for the [Bem](https://docs.bem.ai) document processing API. Installed on this machine as `bem`.

Authentication is via the `BEM_API_KEY` environment variable (already configured).

## Sending Files Through the CLI

### Workflow Call API (`bem workflows call`)

Call a workflow to process a file. The `--input.single-file` flag takes a JSON object with `inputContent` (file data) and `inputType` (file extension).

Use the `@` prefix to embed file contents automatically. Binary files (PDF, images) are base64-encoded; text files are embedded as strings.

```bash
# Call a workflow with a single file (synchronous, waits up to 30s)
bem workflows call \
  --workflow-name <WORKFLOW_NAME> \
  --input.single-file '{"inputContent": "@path/to/file.pdf", "inputType": "pdf"}' \
  --wait

# Without --wait, returns immediately with a callID to poll later
bem workflows call \
  --workflow-name <WORKFLOW_NAME> \
  --input.single-file '{"inputContent": "@path/to/file.pdf", "inputType": "pdf"}'
```

**Supported `inputType` values:** `csv`, `docx`, `email`, `heic`, `html`, `jpeg`, `json`, `heif`, `m4a`, `mp3`, `pdf`, `png`, `text`, `wav`, `webp`, `xls`, `xlsx`, `xml`

**Batch files** (multiple files in one call):

```bash
bem workflows call \
  --workflow-name <WORKFLOW_NAME> \
  --input.batch-files '{"inputs": [{"inputContent": "@file1.pdf", "inputType": "pdf"}, {"inputContent": "@file2.png", "inputType": "png"}]}'
```

**Alternative: pass the full `--input` flag as JSON:**

```bash
bem workflows call \
  --workflow-name <WORKFLOW_NAME> \
  --input '{"singleFile": {"inputContent": "@file.pdf", "inputType": "pdf"}}' \
  --wait
```

**Alternative: pipe YAML/JSON via stdin:**

```bash
cat <<'YAML' | bem workflows call --workflow-name <WORKFLOW_NAME> --wait
input:
  singleFile:
    inputContent: "@path/to/file.pdf"
    inputType: "pdf"
YAML
```

### Infer Schema API (`bem infer-schema create`)

Analyze a file and infer a JSON Schema from its contents. Uses multipart form upload.

```bash
bem infer-schema create --file @path/to/file.pdf
```

### Checking Call Status and Results

```bash
# Retrieve a specific call by ID
bem calls retrieve --call-id <CALL_ID>

# List recent calls, optionally filtered
bem calls list --workflow-name <WORKFLOW_NAME> --status completed

# List outputs for a call
bem outputs list --call-id <CALL_ID>

# List errors for a call
bem errors list --call-id <CALL_ID>
```

### Listing Available Workflows and Functions

```bash
bem workflows list
bem functions list
```

## Common Gotchas

- **`--wait` is a boolean flag.** Use `--wait` or `--wait=true`. Do NOT use `--wait true` (with a space) -- the `true` will be parsed as an unexpected positional argument.
- **The `@` prefix is required** for file embedding in JSON string values. `"@file.pdf"` reads and embeds the file; `"file.pdf"` sends the literal string.
- **Binary files are auto-base64-encoded.** When using `@file.pdf` inside a JSON body, the CLI detects the file type and base64-encodes binary files automatically. No need to manually encode.
- **Use `@data://` to force base64 encoding** for text files: `"@data://file.txt"`.
- **Use `@file://` to force string embedding** for files that might be detected as binary: `"@file://file.bin"`.
- **Escape literal `@` signs** with a backslash: `'\@username'`.

## Output Formatting

```bash
# JSON output (default)
bem workflows call ... --format json

# Extract specific fields with GJSON
bem workflows call ... --transform 'call.outputs.0.transformedContent'

# YAML output
bem workflows call ... --format yaml
```

## Project Structure

- `cmd/bem/main.go` - Entry point
- `pkg/cmd/` - All command implementations (workflow.go, function.go, call.go, etc.)
- `pkg/cmd/flagoptions.go` - File embedding and request building logic
- `internal/requestflag/` - CLI flag parsing
- Generated from OpenAPI spec by [Stainless](https://www.stainless.com/)
