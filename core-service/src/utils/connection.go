package utils

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-redis/redis/v8"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
)

type Conn struct {
	Mysql   *sql.DB
	Elastic *elasticsearch.Client
	AWS     *session.Session
	Redis   *redis.Client
}

func NewDBConn(cfg *config.Config) *Conn {
	return &Conn{
		Mysql:   initMysql(cfg),
		Elastic: initELastClient(cfg),
		AWS:     initAWSClient(cfg),
		Redis:   initRedisClient(cfg),
	}
}

// InitDB a new connection to the database
func initMysql(cfg *config.Config) *sql.DB {
	db, err := sql.Open("mysql", cfg.DB.DSN)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}

// initRedisClient ...
func initRedisClient(cfg *config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
	})

	_, err := rdb.Ping(rdb.Context()).Result()
	if err != nil {
		log.Fatal(err)
	}

	return rdb
}

func initELastClient(cfg *config.Config) *elasticsearch.Client {
	es, err := elasticsearch.NewClient(*cfg.ELastic.ElasticConfig)
	if err != nil {
		log.Fatal(err)
	}

	// 1. Get cluster info
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	return es
}

func initAWSClient(cfg *config.Config) *session.Session {
	s, err := session.NewSession(&aws.Config{
		Region: aws.String(cfg.AWS.Region),
		Credentials: credentials.NewStaticCredentials(
			cfg.AWS.AccessKey,
			cfg.AWS.SecretAccessKey,
			"",
		),
	})

	if err != nil {
		log.Fatal(err)
	}

	return s
}
