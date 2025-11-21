package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func getBaseURL() string {
	if url := os.Getenv("API_URL"); url != "" {
		return url
	}
	return "http://app:8080"
}

func decodeJSON[T any](t *testing.T, resp *http.Response, out *T) {
	defer resp.Body.Close()
	err := json.NewDecoder(resp.Body).Decode(out)
	require.NoError(t, err)
}

func doPost[T any](t *testing.T, url string, body any, expectedStatus int, out *T) {
	data, _ := json.Marshal(body)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	require.NoError(t, err)
	require.Equal(t, expectedStatus, resp.StatusCode)
	if out != nil {
		decodeJSON(t, resp, out)
	} else {
		resp.Body.Close()
	}
}

func doGet[T any](t *testing.T, url string, expectedStatus int, out *T) {
	resp, err := http.Get(url)
	require.NoError(t, err)
	require.Equal(t, expectedStatus, resp.StatusCode)
	if out != nil {
		decodeJSON(t, resp, out)
	} else {
		resp.Body.Close()
	}
}

func doDelete(t *testing.T, url string, expectedStatus int) {
	req, _ := http.NewRequest("DELETE", url, nil)
	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	require.Equal(t, expectedStatus, resp.StatusCode)
	resp.Body.Close()
}
