package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/tdewolff/parse/v2"
	"github.com/tdewolff/parse/v2/css"
)

type font struct {
	nest       int
	properties map[string][]propertyToken
}

type propertyToken struct {
	t css.TokenType
	v []byte
}

func nextNonWhitespaceToken(lexer *css.Lexer) (bool, css.TokenType, []byte) {
	for {
		nextToken, text := lexer.Next()
		switch nextToken {
		case css.WhitespaceToken:
		case css.ErrorToken:
			return true, nextToken, text
		default:
			return false, nextToken, text
		}
	}
}

func expectNextToken(lexer *css.Lexer, expected css.TokenType) (bool, css.TokenType) {
	for {
		nextToken, _ := lexer.Next()
		switch nextToken {
		case css.WhitespaceToken:
		case expected:
			return true, expected
		default:
			return false, nextToken
		}
	}
}

func main() {
	input := `.some-rule {
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
	r := bytes.NewReader([]byte(input))
	lexer := css.NewLexer(parse.NewInput(r))

	nest := 0

	var currentFont *font

loop:
	for {
		tt, text := lexer.Next()
		switch tt {
		case css.ErrorToken:
			// error or EOF set in l.Err()
			break loop
		case css.AtKeywordToken:
			if string(text) == "@font-face" {
				ok, found := expectNextToken(lexer, css.LeftBraceToken)
				if ok {
					nest = nest + 1
				} else {
					fmt.Println("ERROR: Expected token", css.LeftBraceToken, "got", found)
					break loop
				}
				// out of the inner loop
				currentFont = &font{
					nest:       nest,
					properties: map[string][]propertyToken{},
				}

			}
		default:

			switch tt {
			case css.WhitespaceToken:
				// ignore...
			case css.LeftBraceToken:
				nest = nest + 1
			case css.RightBraceToken:
				nest = nest - 1
				if currentFont != nil {
					if currentFont.nest > nest {
						fmt.Println("Font definition completed:", currentFont)
					}
				}
			case css.IdentToken:

				if currentFont != nil {
					propertyName := strings.TrimSpace(string(text))
					ok, found := expectNextToken(lexer, css.ColonToken)
					if !ok {
						fmt.Println("ERROR: Expected token", css.ColonToken, "got", found)
						break loop
					}
					hadError, nextToken, nextTokenText := nextNonWhitespaceToken(lexer)
					if hadError {
						fmt.Println("ERROR: Expected non-whitespace non-error token, got error token", nextToken)
						break loop
					}

					currentFont.properties[propertyName] = []propertyToken{
						{
							t: nextToken,
							v: nextTokenText,
						},
					}

					// now read any token until semicolon:
				aggregator:
					for {
						propToken, propTokenText := lexer.Next()
						switch propToken {
						case css.ErrorToken:
							fmt.Println("ERROR: Expected non-error property token, got error token", propToken)
							break loop
						case css.SemicolonToken:
							break aggregator
						default:
							currentFont.properties[propertyName] = append(currentFont.properties[propertyName], propertyToken{
								t: propToken,
								v: propTokenText,
							})
						}
					}

				}
			default:
				// ignore anything else
			}

		}
	}

}
