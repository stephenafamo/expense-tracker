package main

import (
	// "context"
	"net/url"

	"github.com/spf13/viper"
	// "github.com/volatiletech/authboss"
	// models "github.com/stephenafamo/expense-tracker/auth_models"
	"path/filepath"
)

type PartialProvider struct {
	server
}

func (p PartialProvider) Get(name string) (string, error) {
	return p.getTemplate("/partials/" + name)
}

func getBaseUrl() *url.URL {

	baseUrl, err := url.Parse(viper.GetString("panel_url"))
	checkError(err)

	baseUrl.Path = filepath.Join(baseUrl.Path, "/") // will be "/" if empty

	if baseUrl.Path != "/" {
		baseUrl.Path = baseUrl.Path + "/" // add trailing slash
	}

	return baseUrl
}
