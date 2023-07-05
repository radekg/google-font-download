package downloader

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/radekg/google-font-download/reader"
	"github.com/tdewolff/parse/v2/css"
)

type FontData interface {
	URL() *url.URL
	Reader() io.ReadCloser
}

type defaultFontData struct {
	u *url.URL
	r io.ReadCloser
}

func (dfd *defaultFontData) Reader() io.ReadCloser {
	return dfd.r
}

func (dfd *defaultFontData) URL() *url.URL {
	return dfd.u
}

func Download(url url.URL) ([]FontData, error) {

	fonts := []reader.Font{}
	result := []FontData{}

	request, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}
	// we need something looking like a browser to have the woff2 files served:
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/114.0")
	httpClient := &http.Client{}

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	reader := reader.NewReader(response.Body)

readerloop:
	for {
		select {
		case <-reader.OnDone():
			break readerloop
		case err := <-reader.OnError():
			return nil, err
		case f := <-reader.OnNextFont():
			fonts = append(fonts, f)
		}
	}

	for _, font := range fonts {
		properties, ok := font.Property("src")
		if ok {
			for _, prop := range properties {
				if prop.Type() == css.URLToken {
					fontFileURLString := string(prop.Value())
					fontFileURLString = strings.TrimPrefix(fontFileURLString, "url(")
					fontFileURLString = strings.TrimSuffix(fontFileURLString, ")")
					fontFileURL, err := url.Parse(fontFileURLString)
					if err != nil {
						for _, x := range result {
							x.Reader().Close()
						}
						return nil, err
					}
					if fontFileURL.Scheme == "http" || fontFileURL.Scheme == "https" {

						fmt.Println(" ===========> ", fontFileURLString)
						fontDataResponse, err := http.Get(fontFileURLString)
						if err != nil {
							for _, x := range result {
								x.Reader().Close()
							}
							return nil, err
						}
						result = append(result, &defaultFontData{
							u: fontFileURL,
							r: fontDataResponse.Body,
						})

					}
				}
			}
		}
	}

	return result, nil

}
