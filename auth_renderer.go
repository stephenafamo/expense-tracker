// To satisfy authboss's renderer interface
package main

import (
	"context"

	"github.com/volatiletech/authboss"
)

type AuthRenderer struct {
	*server
	Base string
}

func (AuthRenderer) Load(names ...string) error {
	return nil
}

func (a AuthRenderer) Render(ctx context.Context, page string, data authboss.HTMLData) (output []byte, contentType string, err error) {
	renderedString, err := a.render(a.Base+"/"+page, data)
	return []byte(renderedString), "text/html", err
}
