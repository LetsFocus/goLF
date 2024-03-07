package database

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gocql/gocql"

	"github.com/LetsFocus/goLF/logger"
	"github.com/LetsFocus/goLF/types"
)

type CassandraConfig struct {
	Addresses         []string
	DB                string
	Password          string
	MaxRetries        int
	RetryDuration     int
	MonitoringEnable  bool
	MaxOpenConns      int
	ConnectionTimeout int
}

func (c CassandraConfig) GetHost() string {
	return strings.Join(c.Addresses, ",")
}

func (c CassandraConfig) GetDBName() string {
	return CassandraDB
}
func (c CassandraConfig) GetMaxRetries() int {
	return c.MaxRetries
}
func (c CassandraConfig) GetMaxRetryDuration() int {
	return c.RetryDuration
}

// InitializeCassandra creates a new CassandraClient instance
func InitializeCassandra(log *logger.CustomLogger, c *CassandraConfig) (Cassandra, error) {
	if c.GetHost() != "" {
		session, err := establishCassandraConnection(log, c)
		if err != nil {
			return Cassandra{}, err
		}

		return Cassandra{Session: session}, nil
	}

	return Cassandra{}, nil
}

func establishCassandraConnection(log *logger.CustomLogger, c *CassandraConfig) (*gocql.Session, error) {
	cluster := gocql.NewCluster(c.Addresses...)
	cluster.Keyspace = c.DB

	// number of connections per host
	cluster.NumConns = c.MaxOpenConns

	// maximum time to wait for a single Cassandra operation to complete.
	cluster.Timeout = time.Duration(c.ConnectionTimeout) * time.Second

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

func (c *Cassandra) HealthCheckCassandra() types.Health {
	if c == nil {
		return types.Health{Status: Down, Name: SQL}
	}

	if err := c.Query("SELECT * FROM system.local").Exec(); err != nil {
		return types.Health{Status: Down, Name: SQL}
	}

	return types.Health{Status: Up, Name: SQL}
}

// Close : closes the Cassandra session
func (c *Cassandra) Close() {
	c.Session.Close()
}

// ExecuteQuery : executes a CQL query and returns the iterator
func (c *Cassandra) ExecuteQuery(query string, values ...interface{}) *gocql.Iter {
	return c.Session.Query(query, values...).Iter()
}

// ExecuteBatch : executes a batch operation
func (c *Cassandra) ExecuteBatch(batch *gocql.Batch) error {
	return c.Session.ExecuteBatch(batch)
}

// AsyncQuery : ExecuteQueryAsync executes a CQL query asynchronously and returns a channel to receive the result
func (c *Cassandra) AsyncQuery(ctx context.Context, query string, values ...interface{}) <-chan *gocql.Query {
	resultChan := make(chan *gocql.Query, 1)
	go func() {
		resultChan <- c.Session.Query(query, values...)
	}()
	return resultChan
}

// QueryWithContext : executes a CQL query with context and returns the iterator
func (c *Cassandra) QueryWithContext(ctx context.Context, query string, values ...interface{}) (*gocql.Iter, error) {
	return c.Session.Query(query, values...).WithContext(ctx).Iter(), nil
}

// CreateKeyspace : creates a new keyspace in Cassandra
func (c *Cassandra) CreateKeyspace(keyspace string, replication map[string]interface{}) error {
	query := fmt.Sprintf("CREATE KEYSPACE IF NOT EXISTS %s WITH REPLICATION = %v", keyspace, replication)
	return c.Session.Query(query).Exec()
}

// CreateTable : creates a new table in a keyspace
func (c *Cassandra) CreateTable(keyspace, tableName, schema string) error {
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s.%s (%s)", keyspace, tableName, schema)
	return c.Session.Query(query).Exec()
}

// DropKeyspace : drops a keyspace from Cassandra
func (c *Cassandra) DropKeyspace(keyspace string) error {
	query := fmt.Sprintf("DROP KEYSPACE IF EXISTS %s", keyspace)
	return c.Session.Query(query).Exec()
}

// DropTable : drops a table from a keyspace
func (c *Cassandra) DropTable(keyspace, tableName string) error {
	query := fmt.Sprintf("DROP TABLE IF EXISTS %s.%s", keyspace, tableName)
	return c.Session.Query(query).Exec()
}
