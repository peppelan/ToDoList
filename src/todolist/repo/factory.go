package repo

import (
	"todolist/spi"
	"os"
	"errors"
)

var (
	esUrl = os.Getenv("ELASTICSEARCH_URL")
	esLogin = os.Getenv("ELASTICSEARCH_USERNAME")
	esPwd = os.Getenv("ELASTICSEARCH_PASSWORD")
)

// Creates the appropriate repository depending on the application configuration:
// - set the environment variable ELASTICSEARCH_URL to back the application with ElasticSearch
// - leave the variable unset to let the application use an in-memory repository
func NewRepo() spi.Repo {
	if "" == esUrl {
		r := NewInMemoryRepo()
		r.Init()
		return r
	}

	panic(errors.New("elasticsearch repo not implemented yet"))
}