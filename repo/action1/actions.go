package action1

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func CalcWordCount(t *testing.T, word string, expCount int) {
	url := "http://localhost:3000/wordcount"
	payload := map[string]string{"word": word}
	jsonData, err := json.Marshal(payload)
	require.NoError(t, err)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	var result map[string]int
	err = json.Unmarshal(body, &result)
	require.NoError(t, err)

	require.Equal(t, expCount, result["count"])
}
