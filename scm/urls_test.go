package scm

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUrlJoin(t *testing.T) {
	t.Run("return empty string when passing no parameter", func(t *testing.T) {
		got, err := UrlJoin()

		require.NoError(t, err)
		expectedResult := ""
		require.Equal(t, expectedResult, got)
	})
	t.Run("return the only passed none empty param", func(t *testing.T) {
		got, err := UrlJoin("test")

		require.NoError(t, err)
		expectedResult := "test"
		require.Equal(t, expectedResult, got)
	})
	t.Run("return all params combined", func(t *testing.T) {
		got, err := UrlJoin("test", "bernd", "kram")

		require.NoError(t, err)
		expectedResult := "test/bernd/kram"
		require.Equal(t, expectedResult, got)
	})
	t.Run("return url with valid protocol", func(t *testing.T) {
		got, err := UrlJoin("http://localhost", "bernd", "kram")

		require.NoError(t, err)
		expectedResult := "http://localhost/bernd/kram"
		require.Equal(t, expectedResult, got)
	})
	t.Run("returns error when passing the evil control character 0x7f as first parameter", func(t *testing.T) {
		got, err := UrlJoin(string(0x7f), "bernd", "kram")

		expectedResult := ""
		assert.Error(t, err)
		assert.Equal(t, expectedResult, got)
		expectedErrText := "failed to parse URL parts"
		assert.Contains(t, err.Error(), expectedErrText)
	})
	t.Run("returns combined URL with file", func(t *testing.T) {
		baseURL := "https://myurl.de/nexus/repo/plugins"
		fileName := "test.json"

		resultURL, _ := UrlJoin(baseURL, fileName)

		expectedResult := "https://myurl.de/nexus/repo/plugins/test.json"
		assert.Equal(t, expectedResult, resultURL)
	})
	t.Run("returns combined URL ignoring the passed empty string", func(t *testing.T) {
		baseURL := "https://myurl.de/nexus/repo/plugins"
		emptyStringParam := ""

		resultURL, _ := UrlJoin(baseURL, emptyStringParam)

		expectedResult := "https://myurl.de/nexus/repo/plugins"
		assert.Equal(t, expectedResult, resultURL)
	})
}
