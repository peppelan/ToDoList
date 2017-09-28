package repo

import (
	"todolist/spi"
	"gopkg.in/olivere/elastic.v5"
)

type ElasticSearchRepo struct {
	client *elastic.Client
}

// Creates the client
func NewElasticSearchRepo(esUrl string, esUser string, esPwd string) *ElasticSearchRepo {
	r := new (ElasticSearchRepo)
	client, err := elastic.NewSimpleClient(
		elastic.SetURL(esUrl),
	    elastic.SetMaxRetries(2),
	    elastic.SetBasicAuth(esUser, esPwd))

	if nil != err {
		panic(err)
	}

	r.client = client
	return r
}

func (r *ElasticSearchRepo) Init() error {
	// Nothing to do here for now
	return nil
}

func (r *ElasticSearchRepo) Find(id string) *spi.Todo {
	return nil
}

func (r *ElasticSearchRepo) FindAll() spi.Todos {
	return nil
}

func (r *ElasticSearchRepo) Create(t spi.Todo) string {
	return ""
}

func (r *ElasticSearchRepo) Destroy(id string) bool {
	return false
}

func (r *ElasticSearchRepo) Update(t spi.Todo) bool {
	return false
}
