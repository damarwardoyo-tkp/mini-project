package es

import (
	"github.com/olivere/elastic/v7"
	"log"
)

type ElasticsearchClient struct {
	ES *elastic.Client
}

func newConn() (*elastic.Client, error) {
	es, err := elastic.NewClient()
	if err != nil {
		log.Printf("[Elasticsearch] Gagal, err %v", err)
	}
	return es, err
}

func NewElasticsearchClient() *ElasticsearchClient {
	es, err := newConn()
	if err != nil {
		log.Fatalln("[ElasticsearchClient] Gagal membuat Elasticsearch Client")
	}
	client := ElasticsearchClient{
		ES: es,
	}
	return &client
}
