package database

// import (
// 	"strings"
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/assert"

// 	"github.com/LetsFocus/goLF/configs"
// 	"github.com/LetsFocus/goLF/goLF/model"
// 	"github.com/LetsFocus/goLF/logger"
// )

// func Test_establishESConnection(t *testing.T) {
// 	log := logger.NewCustomLogger()

// 	testcases := []struct {
// 		desc     string
// 		esConfig esConfig
// 		err      error
// 	}{
// 		{
// 			desc:     "successfully established elastic search connection",
// 			esConfig: esConfig{addresses: []string{"http://localhost:9200"}, username: "", password: ""},
// 		},
// 	}

// 	for i, tc := range testcases {
// 		_, err := establishESConnection(log, tc.esConfig)

// 		assert.Equalf(t, tc.err, err, "Test[%d] FAILED, Could not connect to ES, got error: %v\n", i, err)
// 	}
// }

// func Test_InitializeES(t *testing.T) {
// 	t.Setenv("ES_ADDRESSES", "http://localhost:9200")
// 	t.Setenv("ES_USERNAME", "")
// 	t.Setenv("ES_PASSWORD", "")

// 	testcases := []struct {
// 		input *model.GoLF
// 	}{
// 		{
// 			input: &model.GoLF{
// 				Config: configs.Config{Log: logger.NewCustomLogger()},
// 				Logger: logger.NewCustomLogger(),
// 			},
// 		},
// 	}

// 	for _, tc := range testcases {
// 		InitializeES(tc.input, "")
// 	}
// }

// func Test_InitializeDBWithMonitoring(t *testing.T) {
// 	t.Setenv("ES_ADDRESSES", "http://localhost:9200")
// 	t.Setenv("ES_USERNAME", "")
// 	t.Setenv("ES_PASSWORD", "")
// 	t.Setenv("ES_MONITORING", "TRUE")

// 	testcases := []struct {
// 		input *model.GoLF
// 	}{
// 		{
// 			input: &model.GoLF{
// 				Config: configs.Config{Log: logger.NewCustomLogger()},
// 				Logger: logger.NewCustomLogger(),
// 			},
// 		},
// 	}

// 	for _, tc := range testcases {
// 		InitializeES(tc.input, "")
// 	}
// }

// func Test_monitoringES(t *testing.T) {
// 	es1, _ := establishESConnection(logger.NewCustomLogger(), esConfig{addresses: []string{"http://localhost:9200"}, username: "", password: ""})
// 	es2, _ := establishESConnection(logger.NewCustomLogger(), esConfig{addresses: []string{"http://localhost:9100"}, username: "", password: ""})

// 	testcases := []struct {
// 		desc      string
// 		input     *model.GoLF
// 		esConfig  esConfig
// 		retry     int
// 		retryTime int
// 		log       string
// 	}{
// 		{
// 			desc: "successfully monitored the es",
// 			input: &model.GoLF{
// 				Database: model.Database{Elasticsearch: es1},
// 				Config:   configs.Config{Log: logger.NewCustomLogger()},
// 				Logger:   logger.NewCustomLogger(),
// 			},
// 			esConfig: esConfig{addresses: []string{"http://localhost:9200"}, username: "", password: ""},
// 			log:      "Elasticsearch client is connected successfully",
// 		},

// 		{
// 			desc: "retry is less than retryCount",
// 			input: &model.GoLF{
// 				Database: model.Database{Elasticsearch: es2},
// 				Config:   configs.Config{Log: logger.NewCustomLogger()},
// 				Logger:   logger.NewCustomLogger(),
// 			},
// 			esConfig: esConfig{addresses: []string{"http://localhost:9100"}, username: "", password: "", maxRetries: 1, retryDuration: 1},
// 			log:      "Elasticsearch monitoring stopped after reaching maximum retries.",
// 		},
// 	}

// 	for i, tc := range testcases {
// 		go monitoringES(tc.input, tc.esConfig)

// 		if strings.Contains(tc.input.Logger.GetLog(), tc.log) {
// 			t.Errorf("Testcase Failed[%v], Required Log: %v, Got: %v", i+1, tc.input.Logger.GetLog(), tc.log)
// 		}
// 		time.Sleep(time.Second * 3)
// 	}
// }
