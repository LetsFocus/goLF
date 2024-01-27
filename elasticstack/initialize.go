package elasticstack

import (
	"strings"

	"github.com/elastic/go-elasticsearch/v8"

	_ "github.com/lib/pq"

	"github.com/LetsFocus/goLF/configs"
)

type config struct {
	addresses []string
	user      string
	password  string
	retry     int
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

func InitializeES(configs configs.Config, prefix string, retryCounter int) (*elasticsearch.Client, error) {
	cfg := config{
		addresses: cleanAddresses(configs.Get(prefix + "ELASTICSEARCH_ADDRESSES")),
		user:      configs.Get(prefix + "ELASTICSEARCH_USER"),
		password:  configs.Get(prefix + "ELASTICSEARCH_PASSWORD"),
		retry:     retryCounter,
	}

	es, err := elasticsearch.NewClient(
		elasticsearch.Config{
			Addresses: cfg.addresses,
			Username:  cfg.user,
			Password:  cfg.password,
		})
	if err != nil {
		configs.Log.Errorf("Failed to initialize the Elasticsearch client: %v", err)
		if cfg.retry > 0 {
			return InitializeES(configs, prefix, retryCounter-1)
		}

		return nil, err
	}

	_, err = es.Info()
	if err != nil {
		configs.Log.Errorf("Failed to initialize the Elasticsearch client: %v", err)
		if cfg.retry > 0 {
			return InitializeES(configs, prefix, retryCounter-1)
		}

		return nil, err
	}

	configs.Log.Info("Connected to elasticsearch cluster successfully")

	return es, nil
}
