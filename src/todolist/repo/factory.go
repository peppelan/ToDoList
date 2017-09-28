package repo

import (
	"todolist/spi"
	"os"
)

var (
	esUrl = os.Getenv("ELASTICSEARCH_URL")
	esUser = os.Getenv("ELASTICSEARCH_USERNAME")
	esPwd = os.Getenv("ELASTICSEARCH_PASSWORD")
)

// Creates the appropriate repository depending on the application configuration:
// - set the environment variable ELASTICSEARCH_URL to back the application with ElasticSearch
// - leave the variable unset to let the application use an in-memory repository
func NewRepo() spi.Repo {

	var r spi.Repo

	if "" == esUrl {
		r = NewInMemoryRepo()
	} else {
		r = NewElasticSearchRepo(esUrl, esUser, esPwd)
	}

	r.Init()
	return r
}