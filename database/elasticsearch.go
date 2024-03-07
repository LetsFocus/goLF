package database

import (
	"github.com/LetsFocus/goLF/types"
	"net/http"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"

	"github.com/LetsFocus/goLF/logger"
)

type ESConfig struct {
	Addresses             []string
	Username              string
	Password              string
	MaxRetries            int
	RetryDuration         int
	MonitoringEnable      bool
	MaxOpenConns          int
	MaxIdleConns          int
	IdleConnectionTimeout int
}

func (e ESConfig) GetHost() string {
	return strings.Join(e.Addresses, ",")
}

func (e ESConfig) GetDBName() string {
	return ElasticSearch
}
func (e ESConfig) GetMaxRetries() int {
	return e.MaxRetries
}
func (e ESConfig) GetMaxRetryDuration() int {
	return e.RetryDuration
}

func InitializeES(log *logger.CustomLogger, c *ESConfig) (Es, error) {
	transport := &http.Transport{
		MaxIdleConns:    c.MaxIdleConns,
		MaxConnsPerHost: c.MaxOpenConns,
		IdleConnTimeout: time.Minute * time.Duration(c.IdleConnectionTimeout),
	}

	cfg := elasticsearch.Config{
		Addresses: c.Addresses,
		Username:  c.Username,
		Password:  c.Password,
		Transport: transport,
	}

	esClient, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Errorf("Failed to initialize the Elasticsearch client, Error:%v", err)
		return Es{}, err
	}

	_, err = esClient.Info()
	if err != nil {
		log.Errorf("Failed to ping the Elasticsearch cluster, Error:%v", err)
		return Es{}, err
	}

	log.Info("Elasticsearch client is connected successfully")
	return Es{Client: esClient}, nil
}

func (e *Es) HealthCheckES() types.Health {
	if e == nil {
		return types.Health{Status: Down, Name: ElasticSearch}
	}

	if _, err := e.Info(); err != nil {
		return types.Health{Status: Down, Name: ElasticSearch}
	}

	return types.Health{Status: Up, Name: ElasticSearch}
}
