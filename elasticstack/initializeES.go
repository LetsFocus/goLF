package elasticstack

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"

	"github.com/LetsFocus/goLF/goLF/model"
	"github.com/LetsFocus/goLF/logger"
)

type esConfig struct {
	addresses             []string
	username              string
	password              string
	maxRetries            int
	retryDuration         int
	monitoringEnable      bool
	maxOpenConns          int
	maxIdleConns          int
	idleConnectionTimeout int
}

func cleanAddresses(addressesString string) []string {
	addressesList := strings.Split(addressesString, ",")

	var addressesSlice []string

	for _, address := range addressesList {
		trimmedAddress := strings.TrimSpace(address)
		addressesSlice = append(addressesSlice, trimmedAddress)
	}

	return addressesSlice
}

func InitializeES(golf *model.GoLF, prefix string) {
	var (
		maxConnections, maxIdleConnections, idleConnectionTimeout, maxRetries, retryDuration int
		monitoring                                                                           bool
		err                                                                                  error
	)

	maxRetries, err = strconv.Atoi(golf.Config.Get(prefix + "ES_MAX_RETRIES"))
	if err != nil {
		maxRetries = 5
	}

	retryDuration, err = strconv.Atoi(golf.Config.Get(prefix + "ES_RETRY_DURATION"))
	if err != nil {
		retryDuration = 5
	}

	monitoring, err = strconv.ParseBool(golf.Config.Get(prefix + "ES_MONITORING"))
	if err != nil {
		monitoring = false
	}

	maxIdleConnections, err = strconv.Atoi(golf.Config.Get(prefix + "ES_MAX_IDLE_CONNECTIONS"))
	if err != nil {
		maxIdleConnections = 5
	}

	maxConnections, err = strconv.Atoi(golf.Config.Get(prefix + "ES_MAX_CONNECTIONS"))
	if err != nil {
		maxConnections = 20
	}

	idleConnectionTimeout, err = strconv.Atoi(golf.Config.Get(prefix + "ES_IDLE_CONNECTION_TIMEOUT"))
	if err != nil {
		idleConnectionTimeout = 10
	}

	c := esConfig{
		addresses:             cleanAddresses(golf.Config.Get(prefix + "ES_ADDRESSES")),
		username:              golf.Config.Get(prefix + "ES_USERNAME"),
		password:              golf.Config.Get(prefix + "ES_PASSWORD"),
		maxRetries:            maxRetries,
		retryDuration:         retryDuration,
		monitoringEnable:      monitoring,
		maxOpenConns:          maxConnections,
		maxIdleConns:          maxIdleConnections,
		idleConnectionTimeout: idleConnectionTimeout,
	}

	if len(c.addresses) > 0 {
		client, err := establishESConnection(golf.Logger, c)
		if err == nil {
			golf.Elasticsearch = client
			go monitoringES(golf, c)
		}
	}
}

func monitoringES(golf *model.GoLF, c esConfig) {
	ticker := time.NewTicker(time.Second)

	var (
		client       *elasticsearch.Client
		err          error
		retryCounter int
	)

	monitoringLoop:
	for range ticker.C{
		
		if _, err = golf.Elasticsearch.Info(); err != nil {
			if retryCounter < c.maxRetries {
				for i := 0; i < c.maxRetries; i++ {
					client, err = establishESConnection(golf.Logger, c)
					if err == nil {
						golf.Elasticsearch = client
						retryCounter = 0
						break
					}

					retryCounter++
					time.Sleep(time.Second * time.Duration(c.retryDuration))
					golf.Logger.Errorf("ES Retry %d failed: %v", i+1, err)
				}
			} else {
				break monitoringLoop
			}
		} else {
			retryCounter = 0
		}
	}

	ticker.Stop()
	golf.Logger.Errorf("Elasticsearch monitoring stopped after reaching maximum retries. Error for Elasticsearch breakdown is %v", err)
}

func establishESConnection(log *logger.CustomLogger, c esConfig) (*elasticsearch.Client, error) {
	transport := &http.Transport{
		MaxIdleConns:    c.maxIdleConns,
		MaxConnsPerHost: c.maxOpenConns,
		IdleConnTimeout: time.Minute * time.Duration(c.idleConnectionTimeout),
	}

	cfg := elasticsearch.Config{
		Addresses: c.addresses,
		Username:  c.username,
		Password:  c.password,
		Transport: transport,
	}

	esClient, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Errorf("Failed to initialize the Elasticsearch client, Error:%v", err)
		return nil, err
	}

	_, err = esClient.Info()
	if err != nil {
		log.Errorf("Failed to ping the Elasticsearch cluster, Error:%v", err)
		return esClient, err
	}

	log.Info("Elasticsearch client is connected successfully")
	return esClient, nil
}
