package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"

	"github.com/radekg/google-font-download/downloader"
	"github.com/spf13/pflag"
)

var opts struct {
	url    *string
	target *string
}

func setupFlags() {
	opts.url = pflag.String("url", "", "Font URL")
	opts.target = pflag.String("target", "", "Target directory to start outputting files into")
}

func errorExit(e error) {
	fmt.Fprintln(os.Stderr, e.Error())
	os.Exit(1)
}

func main() {
	setupFlags()
	pflag.Parse()

	target := *opts.target

	if target != "" {
		abs, err := filepath.Abs(target)
		if err != nil {
			errorExit(err)
		}
		stat, statErr := os.Stat(abs)
		if statErr != nil {
			errorExit(statErr)
		}
		if !stat.IsDir() {
			errorExit(fmt.Errorf("expected directory at target: %s", target))
		}
		target = abs
	}

	parsedURL, err := url.Parse(*opts.url)
	if err != nil {
		errorExit(err)
	}

	fontData, err := downloader.Download(*parsedURL)
	if err != nil {
		errorExit(err)
	}

	for _, fd := range fontData {
		fontBytes, readErr := ioutil.ReadAll(fd.Reader())
		if readErr != nil {
			fmt.Fprintln(os.Stderr, fmt.Errorf("error: failed reading font data: %v", err))
		} else {
			if target == "" {
				encodedData := base64.StdEncoding.EncodeToString(fontBytes)
				fmt.Println("Font data for", fd.URL().String(), encodedData)
			} else {
				fullPath := filepath.Join(target, fd.URL().Path)
				dir := filepath.Dir(fullPath)
				if err := os.MkdirAll(dir, 0777); err == nil {
					if err := ioutil.WriteFile(fullPath, fontBytes, 0777); err != nil {
						fmt.Fprintln(os.Stderr, fmt.Errorf("error: failed writing font data at %s font data: %v", fullPath, err))
					}
				} else {
					fmt.Fprintln(os.Stderr, fmt.Errorf("error: failed creating target folder for font data at %s font data: %v", fullPath, err))
				}
			}
		}
	}

}
