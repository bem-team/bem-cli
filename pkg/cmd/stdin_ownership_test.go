package cmd

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/bem-team/bem-cli/internal/apiquery"
	"github.com/bem-team/bem-cli/internal/requestflag"
	bem "github.com/bem-team/bem-go-sdk"
	"github.com/bem-team/bem-go-sdk/option"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v3"
)

func TestReferencesStdinFile(t *testing.T) {
	for _, path := range []string{"-", "/dev/stdin", "/dev/fd/0"} {
		for _, prefix := range []string{"@", "@data://", "@file://"} {
			ref := prefix + path
			require.True(t, referencesStdinFile(map[string]any{"nested": []any{nil, map[string]string{"file": ref}}}), ref)
			require.False(t, referencesStdinFile(`\`+ref), ref)
		}
		require.True(t, referencesStdinFile([1]FilePathValue{FilePathValue(path)}))
		require.False(t, referencesStdinFile(path), "untyped paths stay literal")
	}
	for _, value := range []any{nil, 42, "", "@username", "@-suffix", "@data://invoice.pdf", "@file://invoice.txt",
		FilePathValue("@-"), FilePathValue(""), map[string]any{"@-": nil}, []any{}, map[string]any(nil)} {
		require.False(t, referencesStdinFile(value), "%#v", value)
	}
}

func TestFlagOptionsStdinOwnership(t *testing.T) {
	for _, location := range []string{"body", "header", "query"} {
		for _, path := range []string{"-", "/dev/stdin", "/dev/fd/0"} {
			for _, prefix := range []string{"@", "@data://", "@file://", "FileInput"} {
				t.Run(location+"/"+prefix+path, func(t *testing.T) {
					value := any(prefix + path)
					fileInput := prefix == "FileInput"
					if fileInput {
						value = path
					} else if location == "body" {
						value = map[string]any{"nested": []any{value}}
					}
					cmd := ownershipCommand(t, location, fileInput, value)
					options, err := ownershipOptions(t, cmd, "stdin text", ApplicationJSON, false)
					require.NoError(t, err)
					req, body := ownershipRequest(t, options)
					want := "stdin text"
					if prefix == "@data://" {
						want = base64.StdEncoding.EncodeToString([]byte(want))
					}
					switch location {
					case "body":
						var expected any = want
						if !fileInput {
							expected = map[string]any{"nested": []any{want}}
						}
						require.Equal(t, ownershipJSON(t, map[string]any{"value": expected}), body)
					case "header":
						require.Equal(t, want, req.Header.Get("X-Value"))
					case "query":
						require.Equal(t, want, req.URL.Query().Get("value"))
					}
				})
			}
		}
	}
}

func TestFlagOptionsStdinDocumentBytes(t *testing.T) {
	pdf := "%PDF-1.7\n\x00\xff\r\n%%EOF\n"
	for _, prefix := range []string{"@", "@data://", "@file://", "FileInput"} {
		for _, encoding := range []struct {
			name     string
			bodyType BodyContentType
		}{{"json", ApplicationJSON}, {"multipart", MultipartFormEncoded}} {
			t.Run(prefix+"/"+encoding.name, func(t *testing.T) {
				bodyType := encoding.bodyType
				// JSON FileInput and forced text preserve text rather than encoding binary.
				input := pdf
				if bodyType == ApplicationJSON && (prefix == "FileInput" || prefix == "@file://") {
					input = "document\r\nsecond line\n"
				}
				ref := prefix + "-"
				if prefix == "FileInput" {
					ref = "-"
				}
				cmd := ownershipCommand(t, "body", prefix == "FileInput", ref)
				options, err := ownershipOptions(t, cmd, input, bodyType, false)
				require.NoError(t, err)
				req, body := ownershipRequest(t, options)
				if bodyType == MultipartFormEncoded {
					mediaType, params, err := mime.ParseMediaType(req.Header.Get("Content-Type"))
					require.NoError(t, err)
					require.Equal(t, "multipart/form-data", mediaType)
					reader := multipart.NewReader(bytes.NewReader(body), params["boundary"])
					part, err := reader.NextPart()
					require.NoError(t, err)
					require.Equal(t, "value", part.FormName())
					data, err := io.ReadAll(part)
					require.NoError(t, err)
					require.Equal(t, []byte(input), data)
					_, err = reader.NextPart()
					require.ErrorIs(t, err, io.EOF)
				} else {
					want := input
					if prefix == "@" || prefix == "@data://" {
						want = base64.StdEncoding.EncodeToString([]byte(input))
					}
					require.Equal(t, ownershipJSON(t, map[string]any{"value": want}), body)
				}
			})
		}
	}
}

func TestFlagOptionsPipedConfiguration(t *testing.T) {
	path := filepath.Join(t.TempDir(), "document.txt")
	text := "document\r\nsecond line\n"
	require.NoError(t, os.WriteFile(path, []byte(text), 0600))
	for _, tc := range []struct {
		name, value, want, config string
	}{
		{"dash", "-", "-", `{"extra":true}`},
		{"bare-dev-stdin", "/dev/stdin", "/dev/stdin", `{"extra":true}`},
		{"escaped", `\@data://-`, "@data://-", "extra: true\n"},
		{"filename", "document.pdf", "document.pdf", `{"extra":true}`},
		{"file-auto", "@" + path, text, `{"extra":true}`},
		{"file-text", "@file://" + path, text, `{"extra":true}`},
		{"file-base64", "@data://" + path, base64.StdEncoding.EncodeToString([]byte(text)), `{"extra":true}`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			cmd := ownershipCommand(t, "body", false, tc.value)
			options, err := ownershipOptions(t, cmd, tc.config, ApplicationJSON, false)
			require.NoError(t, err)
			_, body := ownershipRequest(t, options)
			require.Equal(t, ownershipJSON(t, map[string]any{"extra": true, "value": tc.want}), body)
		})
	}
	for _, location := range []string{"body", "header", "query"} {
		t.Run("piped-FileInput/"+location, func(t *testing.T) {
			cmd := ownershipCommand(t, location, true)
			key := "value"
			if location == "header" {
				key = "X-Value"
			}
			options, err := ownershipOptions(t, cmd, string(ownershipJSON(t, map[string]any{key: path})), ApplicationJSON, false)
			require.NoError(t, err)
			req, body := ownershipRequest(t, options)
			switch location {
			case "body":
				require.Equal(t, ownershipJSON(t, map[string]any{"value": text}), body)
			case "header":
				require.Equal(t, text, req.Header.Get("X-Value"))
			case "query":
				require.Equal(t, text, req.URL.Query().Get("value"))
			}
		})
	}
}

func TestFlagOptionsStdinConflicts(t *testing.T) {
	t.Run("two-body-consumers", func(t *testing.T) {
		cmd := ownershipCommand(t, "body", false, []any{"@data://-", "@file:///dev/stdin"})
		_, err := ownershipOptions(t, cmd, "document", ApplicationJSON, false)
		require.ErrorContains(t, err, "it can only be read once")
	})
	t.Run("body-and-header-consumers", func(t *testing.T) {
		cmd := ownershipCommand(t, "body", false, "@-")
		header := ownershipCommand(t, "header", false, "@file:///dev/fd/0")
		cmd.Flags = append(cmd.Flags, header.Flags...)
		_, err := ownershipOptions(t, cmd, "document", ApplicationJSON, false)
		require.ErrorContains(t, err, "it can only be read once")
	})
	for _, fileInput := range []bool{false, true} {
		t.Run("config-already-owns-stdin/"+ownershipBool(fileInput), func(t *testing.T) {
			cmd := ownershipCommand(t, "body", fileInput)
			ref := "@-"
			if fileInput {
				ref = "-"
			}
			_, err := ownershipOptions(t, cmd, string(ownershipJSON(t, map[string]any{"value": ref})), ApplicationJSON, false)
			require.ErrorContains(t, err, "stdin was already consumed by piped YAML/JSON input")
		})
		t.Run("ignoreStdin/"+ownershipBool(fileInput), func(t *testing.T) {
			ref := "@-"
			if fileInput {
				ref = "-"
			}
			cmd := ownershipCommand(t, "body", fileInput, ref)
			_, err := ownershipOptions(t, cmd, "document", ApplicationJSON, true)
			require.ErrorContains(t, err, "stdin is already being used for the request body")
			remaining, err := io.ReadAll(os.Stdin)
			require.NoError(t, err)
			require.Equal(t, "document", string(remaining))
		})
	}
	t.Run("ignoreStdin-without-reference", func(t *testing.T) {
		cmd := ownershipCommand(t, "body", false, "literal")
		options, err := ownershipOptions(t, cmd, "not: [valid", ApplicationJSON, true)
		require.NoError(t, err)
		_, body := ownershipRequest(t, options)
		require.Equal(t, []byte(`{"value":"literal"}`), body)
		remaining, err := io.ReadAll(os.Stdin)
		require.NoError(t, err)
		require.Equal(t, "not: [valid", string(remaining))
	})
}

func TestFileInputWrappingOwnership(t *testing.T) {
	flag := &requestflag.Flag[[]string]{Name: "files", BodyPath: "files", FileInput: true}
	require.NoError(t, flag.Set("files", "-"))
	require.NoError(t, flag.Set("files", "invoice.pdf"))
	cmd := &cli.Command{Flags: []cli.Flag{flag}}
	contents := requestflag.ExtractRequestContents(cmd)
	wrapFileInputValues(cmd, &contents)
	want := map[string]any{"files": []any{FilePathValue("-"), FilePathValue("invoice.pdf")}}
	require.Equal(t, want, contents.Body)
	require.True(t, referencesStdinFile(contents.Body))
	wrapFileInputValues(cmd, &contents)
	require.Equal(t, want, contents.Body, "repeated wrapping preserves sentinels")
	require.Equal(t, []string{"-", "invoice.pdf"}, flag.Get(), "classification must not mutate flags")

	root := &requestflag.Flag[string]{Name: "root", FileInput: true, BodyRoot: true}
	require.NoError(t, root.Set("root", "-"))
	cmd = &cli.Command{Flags: []cli.Flag{root}}
	contents = requestflag.ExtractRequestContents(cmd)
	wrapFileInputValues(cmd, &contents)
	require.Equal(t, "-", contents.Body, "body-root uploads retain their separate ownership path")
}

// --- Test helpers: avoid sockets and give each composition test its own stdin. ---

func ownershipCommand(t *testing.T, location string, fileInput bool, values ...any) *cli.Command {
	t.Helper()
	flag := &requestflag.Flag[any]{Name: "value", FileInput: fileInput}
	switch location {
	case "body":
		flag.BodyPath = "value"
	case "header":
		flag.HeaderPath = "X-Value"
	case "query":
		flag.QueryPath = "value"
	}
	for _, value := range values {
		require.NoError(t, flag.Set("value", string(ownershipJSON(t, value))))
	}
	return &cli.Command{Name: "ownership-test", Flags: []cli.Flag{flag}}
}

func ownershipJSON(t *testing.T, value any) []byte {
	t.Helper()
	data, err := json.Marshal(value)
	require.NoError(t, err)
	return data
}

// ownershipOptions exercises pipe detection without depending on a terminal or a loopback listener.
func ownershipOptions(t *testing.T, cmd *cli.Command, input string, bodyType BodyContentType, ignoreStdin bool) ([]option.RequestOption, error) {
	t.Helper()
	file, err := os.CreateTemp(t.TempDir(), "stdin")
	require.NoError(t, err)
	_, err = file.WriteString(input)
	require.NoError(t, err)
	_, err = file.Seek(0, io.SeekStart)
	require.NoError(t, err)
	previous := os.Stdin
	os.Stdin = file
	t.Cleanup(func() {
		os.Stdin = previous
		require.NoError(t, file.Close())
	})
	return flagOptions(cmd, apiquery.NestedQueryFormatDots, apiquery.ArrayQueryFormatRepeat, bodyType, ignoreStdin)
}

type ownershipHTTPClient func(*http.Request) (*http.Response, error)

func (f ownershipHTTPClient) Do(req *http.Request) (*http.Response, error) { return f(req) }

// ownershipRequest checks serialized options in memory; native command proof is a separate replay.
func ownershipRequest(t *testing.T, options []option.RequestOption) (*http.Request, []byte) {
	t.Helper()
	var request *http.Request
	var body []byte
	client := bem.Client{}
	options = append(options, option.WithBaseURL("http://stdin.invalid"), option.WithAPIKey("local-dummy-key"),
		option.WithHTTPClient(ownershipHTTPClient(func(req *http.Request) (*http.Response, error) {
			request = req
			var err error
			body, err = io.ReadAll(req.Body)
			return &http.Response{StatusCode: http.StatusOK, Header: make(http.Header), Body: io.NopCloser(strings.NewReader(`{}`)), Request: req}, err
		})))
	require.NoError(t, client.Post(context.Background(), "/", nil, nil, options...))
	require.NotNil(t, request)
	return request, body
}

func ownershipBool(value bool) string {
	if value {
		return "FileInput"
	}
	return "reference"
}

// --- End test helpers. ---
