package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func makeMockContext(
	e *echo.Echo,
	method,
	path string,
	body interface{},
	pathParams map[string]string,
	queryParams map[string]string,
) (echo.Context, *bytes.Buffer, error) {
	var req *http.Request

	if body != nil {
		byteData, err := json.Marshal(body)
		if err != nil {
			return nil, nil, err
		}
		req = httptest.NewRequest(method, path, bytes.NewReader(byteData))
	} else {
		req = httptest.NewRequest(method, path, nil)
	}

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	buf := new(bytes.Buffer)
	rec.Body = buf

	ctx := e.NewContext(req, rec)

	for k, v := range pathParams {
		ctx.SetParamNames(k)
		ctx.SetParamValues(v)
	}

	q := url.Values{}
	for k, v := range queryParams {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	return ctx, buf, nil
}

func makeGoldenFileName(t *testing.T, dir string) string {
	return "testdata/" + dir + "/" + strings.ReplaceAll(t.Name(), "/", "_") + ".golden"
}

func writeGoldenFile(t *testing.T, filename string, body []byte) {
	t.Helper()

	formattedJSON, err := formatJSON(body)
	if err != nil {
		t.Fatalf("Failed to formatJSON: %s", err)
	}

	if err := os.WriteFile(filename, formattedJSON, 0644); err != nil {
		t.Fatalf("Failed to write golden file: %s", err)
	}
}

func assertGoldenFile(t *testing.T, goldenFilePath string, body []byte) {
	t.Helper()

	expected, err := os.ReadFile(goldenFilePath)
	if err != nil {
		t.Fatalf("Failed to read golden file: %s", err)
	}

	formattedJSON, err := formatJSON(body)
	if err != nil {
		t.Fatalf("Failed to formatJSON: %s", err)
	}

	assert.Equal(t, string(expected), string(formattedJSON))
}

func formatJSON(body []byte) ([]byte, error) {
	a := string(body)
	fmt.Println(a)

	jsonType := determineJSONType(body)

	switch jsonType {
	case "object":
		var result map[string]interface{}
		err := json.Unmarshal(body, &result)
		if err != nil {
			return nil, err
		}
		return json.MarshalIndent(result, "", "    ")
	case "array":
		var result []interface{}
		err := json.Unmarshal(body, &result)
		if err != nil {
			return nil, err
		}
		return json.MarshalIndent(result, "", "    ")
	default:
		return nil, fmt.Errorf("failed to json format jsonType: %s", jsonType)
	}
}

func determineJSONType(data []byte) string {
	// 先頭のホワイトスペースをスキップ
	data = bytes.TrimSpace(data)

	if len(data) == 0 {
		return "empty"
	}

	switch data[0] {
	case '{':
		return "object"
	case '[':
		return "array"
	default:
		return "unknown"
	}
}
