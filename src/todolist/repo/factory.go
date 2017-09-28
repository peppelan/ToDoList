package repo

import (
	"todolist/spi"
	"os"
)

var (
	esUrl = os.Getenv("ELASTICSEARCH_URL")
)

// Creates the appropriate repository depending on the application configuration:
// - set the environment variable ELASTICSEARCH_URL to back the application with ElasticSearch
// - leave the variable unset to let the application use an in-memory repository
func NewRepo() spi.Repo {

	var r spi.Repo

	if "" == esUrl {
		r = NewInMemoryRepo()
	} else {
		r = NewElasticSearchRepo(esUrl)
	}

	r.Init()
	return r
}
