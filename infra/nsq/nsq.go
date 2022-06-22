package nsq

import (
	"github.com/nsqio/go-nsq"
	"log"
)

type NSQProducer struct {
	Producer *nsq.Producer
}

func newConn() (*nsq.Producer, error) {
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Printf("[NSQ Producer] Gagal membuat koneksi ke nsq, err:%v", err)
		return nil, err
	}
	return producer, err
}

func NewNSQProducer() *NSQProducer {
	producer, err := newConn()
	if err != nil {
		log.Fatalln("[NSQ Producer] Gagal membuat NewNSQProducer")
	}
	nsqProducer := NSQProducer{
		Producer: producer,
	}
	return &nsqProducer
}
