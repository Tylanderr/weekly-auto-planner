package html

import (
	"bytes"
	"context"
	"github.com/a-h/templ"
)

func TemplString(t templ.Component) (string, error) {
    var b bytes.Buffer
    if err := t.Render(context.Background(), &b); err != nil {
        return "", err
    }
    return b.String(), nil
}
