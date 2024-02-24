package database

import (
	"context"
	"fmt"
	"github.com/LetsFocus/goLF/goLF/model"

	"strconv"
	"time"

	"github.com/gocql/gocql"

	"github.com/LetsFocus/goLF/logger"
)

type cassandraConfig struct {
	addresses         []string
	db                string
	password          string
	maxRetries        int
	retryDuration     int
	monitoringEnable  bool
	maxOpenConns      int
	connectionTimeout int
}

// InitializeCassandra creates a new CassandraClient instance
func InitializeCassandra(golf *model.GoLF, prefix string) (*gocql.Session, error) {
	var (
		maxConnections, connectionTimeout, maxRetries, retryDuration int
		monitoring                                                   bool
		err                                                          error
	)

	maxRetries, err = strconv.Atoi(golf.Config.Get(prefix + "CASSANDRA_RETRY_COUNT"))
	if err != nil {
		maxRetries = 5
	}

	retryDuration, err = strconv.Atoi(golf.Config.Get(prefix + "CASSANDRA_RETRY_DURATION"))
	if err != nil {
		retryDuration = 5
	}

	monitoring, err = strconv.ParseBool(golf.Config.Get(prefix + "CASSANDRA_MONITORING"))
	if err != nil {
		monitoring = false
	}

	maxConnections, err = strconv.Atoi(golf.Config.Get(prefix + "CASSANDRA_MAX_CONNECTIONS"))
	if err != nil {
		maxConnections = 20
	}

	connectionTimeout, err = strconv.Atoi(golf.Config.Get(prefix + "CASSANDRA_TIMEOUT"))
	if err != nil {
		connectionTimeout = 100
	}

	addressesString := golf.Config.Get(prefix + "CASSANDRA_ADDRESSES")

	c := cassandraConfig{
		addresses:         cleanAddresses(addressesString),
		password:          golf.Config.Get(prefix + "CASSANDRA_PASSWORD"),
		maxRetries:        maxRetries,
		db:                golf.Config.Get(prefix + "CASSANDRA_DB"),
		retryDuration:     retryDuration,
		monitoringEnable:  monitoring,
		maxOpenConns:      maxConnections,
		connectionTimeout: connectionTimeout,
	}

	if addressesString != "" {
		session, err := establishCassandraConnection(golf.Logger, c)
		if err == nil {
			golf.Cql.Session = session
			go monitoringCassandra(golf, c)
		}
	}

	return
}

func establishCassandraConnection(log *logger.CustomLogger, c cassandraConfig) (*gocql.Session, error) {
	cluster := gocql.NewCluster(c.addresses...)
	cluster.Keyspace = c.db

	// number of connections per host
	cluster.NumConns = c.maxOpenConns

	// maximum time to wait for a single Cassandra operation to complete.
	cluster.Timeout = time.Duration(c.connectionTimeout) * time.Second

	cluster.RetryPolicy = &gocql.SimpleRetryPolicy{NumRetries: 0} // Disable Gocql's built-in retry mechanism

	cluster.PoolConfig = gocql.PoolConfig{
		HostSelectionPolicy: gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy()),
	}

	cluster.Consistency = gocql.Quorum

	// Create the initial session
	session, err := cluster.CreateSession()
	if err != nil {
		log.Errorf("Failed to create the cassandra session, Error:%v", err)
		return nil, err
	}

	return session, err
}

func monitoringCassandra(golf *model.GoLF, c cassandraConfig) {
	ticker := time.NewTicker(time.Second)

	var (
		client       *gocql.Session
		err          error
		retryCounter int
	)

monitoringLoop:
	for range ticker.C {
		if err := golf.Cql.Query("SELECT * FROM system.local").Exec(); err != nil {
			if retryCounter < c.maxRetries {
				for i := 0; i < c.maxRetries; i++ {
					client, err = establishCassandraConnection(golf.Logger, c)
					if err == nil {
						golf.Cql.Session = client
						retryCounter = 0
						break
					}

					retryCounter++
					time.Sleep(time.Second * time.Duration(c.retryDuration))
					golf.Logger.Errorf("Cassandra Retry %d failed: %v", i+1, err)
				}
			} else {
				break monitoringLoop
			}
		} else {
			retryCounter = 0
		}
	}

	ticker.Stop()
	golf.Logger.Errorf("Cassandra monitoring stopped after reaching maximum retries. Error for cassandra breakdown is %v", err)
}

// Close : closes the Cassandra session
func (c *CassandraClient) Close() {
	c.Session.Close()
}

// ExecuteQuery : executes a CQL query and returns the iterator
func (c *CassandraClient) ExecuteQuery(query string, values ...interface{}) *gocql.Iter {
	return c.Session.Query(query, values...).Iter()
}

// ExecuteBatch : executes a batch operation
func (c *CassandraClient) ExecuteBatch(batch *gocql.Batch) error {
	return c.Session.ExecuteBatch(batch)
}

// AsyncQuery : ExecuteQueryAsync executes a CQL query asynchronously and returns a channel to receive the result
func (c *CassandraClient) AsyncQuery(ctx context.Context, query string, values ...interface{}) <-chan *gocql.Query {
	resultChan := make(chan *gocql.Query, 1)
	go func() {
		resultChan <- c.Session.Query(query, values...)
	}()
	return resultChan
}

// QueryWithContext : executes a CQL query with context and returns the iterator
func (c *CassandraClient) QueryWithContext(ctx context.Context, query string, values ...interface{}) (*gocql.Iter, error) {
	return c.Session.Query(query, values...).WithContext(ctx).Iter(), nil
}

// CreateKeyspace : creates a new keyspace in Cassandra
func (c *CassandraClient) CreateKeyspace(keyspace string, replication map[string]interface{}) error {
	query := fmt.Sprintf("CREATE KEYSPACE IF NOT EXISTS %s WITH REPLICATION = %v", keyspace, replication)
	return c.Session.Query(query).Exec()
}

// CreateTable : creates a new table in a keyspace
func (c *CassandraClient) CreateTable(keyspace, tableName, schema string) error {
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s.%s (%s)", keyspace, tableName, schema)
	return c.Session.Query(query).Exec()
}

// DropKeyspace : drops a keyspace from Cassandra
func (c *CassandraClient) DropKeyspace(keyspace string) error {
	query := fmt.Sprintf("DROP KEYSPACE IF EXISTS %s", keyspace)
	return c.Session.Query(query).Exec()
}

// DropTable : drops a table from a keyspace
func (c *CassandraClient) DropTable(keyspace, tableName string) error {
	query := fmt.Sprintf("DROP TABLE IF EXISTS %s.%s", keyspace, tableName)
	return c.Session.Query(query).Exec()
}
