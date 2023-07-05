package reader

import (
	"fmt"

	"github.com/tdewolff/parse/v2/css"
)

type ErrUnexpectedToken struct {
	ExpectedToken css.TokenType
	ReceivedToken css.TokenType
}

func (e *ErrUnexpectedToken) Error() string {
	return fmt.Sprintf("font reader: unexpected token, expected: %s, received: %s", e.ExpectedToken.String(), e.ReceivedToken.String())
}

type ErrUnexpectedError struct {
	ReceivedToken  css.TokenType
	CurrentContext Font
}

func (e *ErrUnexpectedError) Error() string {
	return fmt.Sprintf("font reader: non-whitespace non-error token, got error token, received: %s", e.ReceivedToken.String())
}
