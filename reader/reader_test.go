package reader

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tdewolff/parse/v2/css"
)

func TestValidInput(t *testing.T) {
	reader := NewReader(bytes.NewReader([]byte(testInputValid)))

	var readerError error
	fonts := []Font{}

readerloop:
	for {
		select {
		case <-reader.OnDone():
			break readerloop
		case err := <-reader.OnError():
			readerError = err
			break readerloop
		case f := <-reader.OnNextFont():
			fonts = append(fonts, f)
		}
	}

	assert.Nil(t, readerError)
	assert.Equal(t, 2, len(fonts))

	properties, ok := fonts[0].Property("src")

	assert.True(t, ok)
	assert.Eventually(t, func() bool {
		for _, prop := range properties {
			if prop.Type() == css.URLToken {
				return string(prop.Value()) == "url(https://fonts.gstatic.com/s/sourcecodepro/v22/HI_QiYsKILxRpg3hIP6sJ7fM7PqlONvYlMIFxGC8NAU.woff2)"
			}
		}
		return false
	}, 100*time.Millisecond, 10*time.Millisecond)
}

func TestInvalidInput(t *testing.T) {
	reader := NewReader(bytes.NewReader([]byte(testInputInvalid)))

	var readerError error

readerloop:
	for {
		select {
		case <-reader.OnDone():
			break readerloop
		case err := <-reader.OnError():
			readerError = err
			break readerloop
		case <-reader.OnNextFont():
		}
	}

	assert.NotNil(t, readerError)
}

const testInputValid = `
	.some-rule {
		whatever: 100px;
	}
	@font-face {
		font-family: 'Source Code Pro';
		font-style: italic;
		font-weight: 700;
		font-display: swap;
		src: url(https://fonts.gstatic.com/s/sourcecodepro/v22/HI_QiYsKILxRpg3hIP6sJ7fM7PqlONvYlMIFxGC8NAU.woff2) format('woff2');
		unicode-range: U+1F00-1FFF;
	}

	/* latin */
	@font-face {
		font-family: 'Source Code Pro';
		font-style: italic;
		font-weight: 700;
		font-display: swap;
		src: url(https://fonts.gstatic.com/s/sourcecodepro/v22/HI_QiYsKILxRpg3hIP6sJ7fM7PqlONvUlMIFxGC8.woff2) format('woff2');
		unicode-range: U+0000-00FF, U+0131, U+0152-0153, U+02BB-02BC, U+02C6, U+02DA, U+02DC, U+0304, U+0308, U+0329, U+2000-206F, U+2074, U+20AC, U+2122, U+2191, U+2193, U+2212, U+2215, U+FEFF, U+FFFD;
	}
`

const testInputInvalid = `
	@font-face unexpected {
		font-family: 'Source Code Pro';
		font-style: italic;
		font-weight: 700;
		font-display: swap;
		src: url(https://fonts.gstatic.com/s/sourcecodepro/v22/HI_QiYsKILxRpg3hIP6sJ7fM7PqlONvYlMIFxGC8NAU.woff2) format('woff2');
		unicode-range: U+1F00-1FFF;
	}
`
