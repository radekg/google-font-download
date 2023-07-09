package reader

import (
	"io"
	"strings"

	"github.com/tdewolff/parse/v2"
	"github.com/tdewolff/parse/v2/css"
)

// Public types

type Font interface {
	Property(string) ([]Property, bool)
}

type Property interface {
	Type() css.TokenType
	Value() []byte
}

type Reader interface {
	OnNextFont() <-chan Font
	OnError() <-chan error
	OnDone() <-chan struct{}
}

// Helper types:

type font struct {
	nest       int
	properties map[string][]Property
}

func (f *font) Property(p string) ([]Property, bool) {
	props, ok := f.properties[p]
	return props, ok
}

type propertyToken struct {
	t css.TokenType
	v []byte
}

func (pt *propertyToken) Type() css.TokenType {
	return pt.t
}

func (pt *propertyToken) Value() []byte {
	return pt.v
}

// Implementation:

type defaultReader struct {
	lexer     *css.Lexer
	chanFont  chan Font
	chanError chan error
	chanDone  chan struct{}
}

func NewReader(input io.Reader) Reader {
	r := &defaultReader{
		lexer:     css.NewLexer(parse.NewInput(input)),
		chanFont:  make(chan Font),
		chanError: make(chan error, 1),
		chanDone:  make(chan struct{}),
	}
	go r.process()
	return r
}

func (dr *defaultReader) OnDone() <-chan struct{} {
	return dr.chanDone
}

func (dr *defaultReader) OnError() <-chan error {
	return dr.chanError
}

func (dr *defaultReader) OnNextFont() <-chan Font {
	return dr.chanFont
}

func (dr *defaultReader) process() {
	nest := 0
	var currentFont *font

loop:
	for {
		tt, text := dr.lexer.Next()
		switch tt {
		case css.ErrorToken:
			// error or EOF set in l.Err()
			lexerErr := dr.lexer.Err()
			if lexerErr == io.EOF {
				close(dr.chanDone)
			} else {
				dr.chanError <- &ErrUnexpectedError{
					ReceivedToken:  tt,
					CurrentContext: currentFont,
				}
			}
			break loop
		case css.AtKeywordToken:
			if string(text) == "@font-face" {
				ok, found := expectNextToken(dr.lexer, css.LeftBraceToken)
				if ok {
					nest = nest + 1
				} else {
					dr.chanError <- &ErrUnexpectedToken{
						ExpectedToken: css.LeftBraceToken,
						ReceivedToken: found,
					}
					break loop
				}
				// out of the inner loop
				currentFont = &font{
					nest:       nest,
					properties: map[string][]Property{},
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
						dr.chanFont <- currentFont
					}
				}
			case css.IdentToken:

				if currentFont != nil {
					propertyName := strings.TrimSpace(string(text))
					ok, found := expectNextToken(dr.lexer, css.ColonToken)
					if !ok {
						dr.chanError <- &ErrUnexpectedToken{
							ExpectedToken: css.ColonToken,
							ReceivedToken: found,
						}
						break loop
					}
					hadError, nextToken, nextTokenText := nextNonWhitespaceToken(dr.lexer)
					if hadError {
						dr.chanError <- &ErrUnexpectedError{
							ReceivedToken:  found,
							CurrentContext: currentFont,
						}
						break loop
					}

					currentFont.properties[propertyName] = []Property{
						&propertyToken{
							t: nextToken,
							v: nextTokenText,
						},
					}

					// now read any token until semicolon:
				aggregator:
					for {
						propToken, propTokenText := dr.lexer.Next()
						switch propToken {
						case css.ErrorToken:
							dr.chanError <- &ErrUnexpectedError{
								ReceivedToken:  propToken,
								CurrentContext: currentFont,
							}
							break loop
						case css.SemicolonToken:
							break aggregator
						default:
							currentFont.properties[propertyName] = append(currentFont.properties[propertyName], &propertyToken{
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
