package repo

import (
	"todolist/spi"
	"gopkg.in/olivere/elastic.v5"
	"os"
	"context"
)

var (
	esUser = getEnv("ELASTICSEARCH_USERNAME", "elastic")
	esPwd = getEnv("ELASTICSEARCH_PASSWORD", "changeme")

	esIndex = getEnv("ELASTICSEARCH_INDEX", "todolist")
	esType = getEnv("ELASTICSEARCH_TYPE", "todos")
)


func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

type ElasticSearchRepo struct {
	client *elastic.Client
}

// Creates the client
func NewElasticSearchRepo(esUrl string) *ElasticSearchRepo {
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
	// Todo: retry
	return r.init()
}

func (r *ElasticSearchRepo) init() error {
	// TODO: create index & mappings
	return nil
}

func (r *ElasticSearchRepo) Find(id string) *spi.Todo {
	return nil
}

func (r *ElasticSearchRepo) FindAll() []spi.Todo {
	return nil
}

func (r *ElasticSearchRepo) Create(t spi.Todo) string {
	res, err := elastic.NewIndexService(r.client).Index(esIndex).Type(esType).BodyJson(t).Do(context.TODO())

	if (nil != err) {
		panic(err)
	}

	return res.Id
}

func (r *ElasticSearchRepo) Destroy(id string) bool {
	return false
}

func (r *ElasticSearchRepo) Update(id string, t spi.Todo) bool {
	return false
}
