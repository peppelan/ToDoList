package repo

import (
	"todolist/spi"
	"gopkg.in/olivere/elastic.v5"
	"os"
	"context"
	"encoding/json"
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
	if !r.exists(id) {
		return nil
	}

	res, err := r.client.Get().Index(esIndex).Type(esType).Id(id).Do(context.TODO())

	if nil != err {
		panic(err)
	}

	ret := new(spi.Todo)

	err = json.Unmarshal(*res.Source, ret)
	if err != nil {
		panic(err)
	}
	return ret
}

func (r *ElasticSearchRepo) FindAll() map[string] *spi.Todo {
	return nil
}

func (r *ElasticSearchRepo) Create(t spi.Todo) string {
	res, err := r.client.Index().Index(esIndex).Type(esType).BodyJson(t).Do(context.TODO())

	if nil != err {
		panic(err)
	}

	return res.Id
}

func (r *ElasticSearchRepo) Destroy(id string) bool {
	// FIXME: is there a better API in Elasticsearch for failing the destroy if the document does not exist?
	if !r.exists(id) {
		return false
	}

	resp, err := r.client.Delete().Index(esIndex).Type(esType).Id(id).Do(context.TODO())

	if nil != err {
		panic(err)
	}

	return resp.Found
}

func (r *ElasticSearchRepo) Update(id string, t spi.Todo) bool {
	// FIXME: is there a better API in Elasticsearch for failing the update if the document does not exist?
	if !r.exists(id) {
		return false
	}

	_, err := r.client.Update().Index(esIndex).Type(esType).Id(id).Doc(t).Do(context.TODO())

	if nil != err {
		panic(err)
	}

	return true
}

func (r *ElasticSearchRepo) exists(id string) bool {
	res, err := r.client.Exists().Index(esIndex).Type(esType).Id(id).Do(context.TODO())

	if nil != err {
		panic(err)
	}

	return res
}
