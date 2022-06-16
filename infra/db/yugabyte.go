package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

type YugabyteClient struct {
	DB *gorm.DB
}

func newConn() (*gorm.DB, error) {
	dbUri := "host=localhost user=yugabyte password=yugabyte dbname=testdb port=5433 sslmode=disable TimeZone=Asia/Shanghai"
	dbConn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		log.Printf("[Yugabyte] Gagal menginisia koneksi ke database, err: %v", err)
		return nil, err
	}
	return dbConn, nil
}

func NewYugabyteClient() *YugabyteClient {
	db, err := newConn()
	if err != nil {
		log.Fatalln("[YugabyteClient] Gagal membuat Yugabyte Client")
	}
	client := YugabyteClient{
		DB: db,
	}
	fmt.Println("berhasail menginisiasi yugabyte")
	return &client
}
