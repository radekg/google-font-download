package downloader

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDownloader(t *testing.T) {
	var boundURL string

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/fetch" {
			w.Write([]byte(setupFontData(boundURL)))
			return
		}
		if strings.HasPrefix(r.URL.Path, "/s/sourcecodepro") {
			w.Write([]byte(fakeFontData))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))

	defer svr.Close()
	boundURL = svr.URL

	parsedURL, err := url.Parse(fmt.Sprintf("%s/fetch", boundURL))
	assert.Nil(t, err)

	fontData, err := Download(*parsedURL)
	assert.Nil(t, err)
	assert.Greater(t, len(fontData), 0)

	for _, fd := range fontData {
		responseData, err := ioutil.ReadAll(fd.Reader())
		fd.Reader().Close()
		assert.Nil(t, err)
		assert.Equal(t, fakeFontData, string(responseData))
	}

}

func setupFontData(prefix string) string {
	return fmt.Sprintf(format, prefix, prefix)
}

const fakeFontData = "binary font data would be here"

const format = `
@font-face {
	font-family: 'Source Code Pro';
	font-style: italic;
	font-weight: 700;
	font-display: swap;
	src: url(%s/s/sourcecodepro/v22/HI_QiYsKILxRpg3hIP6sJ7fM7PqlONvYlMIFxGC8NAU.woff2) format('woff2');
	unicode-range: U+1F00-1FFF;
}

/* latin */
@font-face {
	font-family: 'Source Code Pro';
	font-style: italic;
	font-weight: 700;
	font-display: swap;
	src: url(%s/s/sourcecodepro/v22/HI_QiYsKILxRpg3hIP6sJ7fM7PqlONvUlMIFxGC8.woff2) format('woff2');
	unicode-range: U+0000-00FF, U+0131, U+0152-0153, U+02BB-02BC, U+02C6, U+02DA, U+02DC, U+0304, U+0308, U+0329, U+2000-206F, U+2074, U+20AC, U+2122, U+2191, U+2193, U+2212, U+2215, U+FEFF, U+FFFD;
}
`
