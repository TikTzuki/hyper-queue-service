package config

import (
	"encoding/json"
	"time"
)

type (
	// Config contains the configuration for a set of cadence services
	Config struct {
		// Persistence contains the configuration for cadence datastores
		Persistence Persistence `yaml:"persistence"`
		// Log is the logging config
		Log                 Logger                    `yaml:"log"`
		DCRedirectionPolicy *ClusterRedirectionPolicy `yaml:"dcRedirectionPolicy"`
		// Services is a map of service name to service config items
		Services map[string]Service `yaml:"services"`
		// Kafka is the config for connecting to kafka
		// Archival is the config for archival
		Archival Archival `yaml:"archival"`
		// PublicClient is config for sys worker service connecting to cadence frontend
		PublicClient PublicClient `yaml:"publicClient"`
		// DynamicConfigClient is the config for setting up the file based dynamic config client
		// Filepath would be relative to the root directory when the path wasn't absolute.
		// Included for backwards compatibility, please transition to DynamicConfig
		// If both are specified, DynamicConig will be used.
		// DomainDefaults is the default config for every domain
		DomainDefaults DomainDefaults `yaml:"domainDefaults"`
		// Blobstore is the config for setting up blobstore
		Blobstore Blobstore `yaml:"blobstore"`
		// Authorization is the config for setting up authorization
		Authorization Authorization `yaml:"authorization"`
	}

	Authorization struct {
		OAuthAuthorizer OAuthAuthorizer `yaml:"oauthAuthorizer"`
		NoopAuthorizer  NoopAuthorizer  `yaml:"noopAuthorizer"`
	}

	NoopAuthorizer struct {
		Enable bool `yaml:"enable"`
	}

	OAuthAuthorizer struct {
		Enable bool `yaml:"enable"`
		// Credentials to verify/create the JWT
		JwtCredentials JwtCredentials `yaml:"jwtCredentials"`
		// Max of TTL in the claim
		MaxJwtTTL int64 `yaml:"maxJwtTTL"`
	}

	JwtCredentials struct {
		// support: RS256 (RSA using SHA256)
		Algorithm string `yaml:"algorithm"`
		// Public Key Path for verifying JWT token passed in from external clients
		PublicKey string `yaml:"publicKey"`
	}

	// Service contains the service specific config items
	Service struct {
		// TChannel is the tchannel configuration
		RPC RPC `yaml:"rpc"`
		// PProf is the PProf configuration
		PProf PProf `yaml:"pprof"`
	}

	// PProf contains the rpc config items
	PProf struct {
		// Port is the port on which the PProf will bind to
		Port int `yaml:"port"`
	}

	// RPC contains the rpc config items
	RPC struct {
		// Port is the port  on which the Thrift TChannel will bind to
		Port uint16 `yaml:"port"`
		// GRPCPort is the port on which the grpc listener will bind to
		GRPCPort uint16 `yaml:"grpcPort"`
		// BindOnLocalHost is true if localhost is the bind address
		BindOnLocalHost bool `yaml:"bindOnLocalHost"`
		// BindOnIP can be used to bind service on specific ip (eg. `0.0.0.0`) -
		// check net.ParseIP for supported syntax, only IPv4 is supported,
		// mutually exclusive with `BindOnLocalHost` option
		BindOnIP string `yaml:"bindOnIP"`
		// DisableLogging disables all logging for rpc
		DisableLogging bool `yaml:"disableLogging"`
		// LogLevel is the desired log level
		LogLevel string `yaml:"logLevel"`
		// GRPCMaxMsgSize allows overriding default (4MB) message size for gRPC
		GRPCMaxMsgSize int `yaml:"grpcMaxMsgSize"`
	}

	// Blobstore contains the config for blobstore
	Blobstore struct {
		Filestore *FileBlobstore `yaml:"filestore"`
	}

	// FileBlobstore contains the config for a file backed blobstore
	FileBlobstore struct {
		OutputDirectory string `yaml:"outputDirectory"`
	}

	// Persistence contains the configuration for data store / persistence layer
	Persistence struct {
		// DefaultStore is the name of the default data store to use
		DefaultStore string `yaml:"defaultStore" validate:"nonzero"`
		// VisibilityStore is the name of the datastore to be used for visibility records
		// Must provide one of VisibilityStore and AdvancedVisibilityStore
		VisibilityStore string `yaml:"visibilityStore"`
		// AdvancedVisibilityStore is the name of the datastore to be used for visibility records
		// Must provide one of VisibilityStore and AdvancedVisibilityStore
		AdvancedVisibilityStore string `yaml:"advancedVisibilityStore"`
		// HistoryMaxConns is the desired number of conns to history store. Value specified
		// here overrides the MaxConns config specified as part of datastore
		HistoryMaxConns int `yaml:"historyMaxConns"`
		// EnablePersistenceLatencyHistogramMetrics is to enable latency histogram metrics for persistence layer
		EnablePersistenceLatencyHistogramMetrics bool `yaml:"enablePersistenceLatencyHistogramMetrics"`
		// NumHistoryShards is the desired number of history shards. It's for computing the historyShardID from workflowID into [0, NumHistoryShards)
		// Therefore, the value cannot be changed once set.
		// TODO This config doesn't belong here, needs refactoring
		NumHistoryShards int `yaml:"numHistoryShards" validate:"nonzero"`
		// DataStores contains the configuration for all datastores
		DataStores map[string]DataStore `yaml:"datastores"`
	}

	// DataStore is the configuration for a single datastore
	DataStore struct {
		// Cassandra contains the config for a cassandra datastore
		// Deprecated: please use NoSQL instead, the structure is backward-compatible
		Cassandra *Cassandra `yaml:"cassandra"`
		// SQL contains the config for a SQL based datastore
		SQL *SQL `yaml:"sql"`
		// NoSQL contains the config for a NoSQL based datastore
		NoSQL *NoSQL `yaml:"nosql"`
	}

	// Cassandra contains configuration to connect to Cassandra cluster
	// Deprecated: please use NoSQL instead, the structure is backward-compatible
	Cassandra = NoSQL

	// NoSQL contains configuration to connect to NoSQL Database cluster
	NoSQL struct {
		// PluginName is the name of NoSQL plugin, default is "cassandra". Supported values: cassandra
		PluginName string `yaml:"pluginName"`
		// Hosts is a csv of cassandra endpoints
		Hosts string `yaml:"hosts" validate:"nonzero"`
		// Port is the cassandra port used for connection by gocql client
		Port int `yaml:"port"`
		// User is the cassandra user used for authentication by gocql client
		User string `yaml:"user"`
		// Password is the cassandra password used for authentication by gocql client
		Password string `yaml:"password"`
		// AllowedAuthenticators informs the cassandra client to expect a custom authenticator
		AllowedAuthenticators []string `yaml:"allowedAuthenticators"`
		// Keyspace is the cassandra keyspace
		Keyspace string `yaml:"keyspace"`
		// Region is the region filter arg for cassandra
		Region string `yaml:"region"`
		// Datacenter is the data center filter arg for cassandra
		Datacenter string `yaml:"datacenter"`
		// MaxConns is the max number of connections to this datastore for a single keyspace
		MaxConns int `yaml:"maxConns"`
		// ProtoVersion
		ProtoVersion int `yaml:"protoVersion"`
		// ConnectAttributes is a set of key-value attributes as a supplement/extension to the above common fields
		// Use it ONLY when a configure is too specific to a particular NoSQL database that should not be in the common struct
		// Otherwise please add new fields to the struct for better documentation
		// If being used in any database, update this comment here to make it clear
		ConnectAttributes map[string]string `yaml:"connectAttributes"`
	}

	// SQL is the configuration for connecting to a SQL backed datastore
	SQL struct {
		// User is the username to be used for the conn
		// If useMultipleDatabases, must be empty and provide it via multipleDatabasesConfig instead
		User string `yaml:"user"`
		// Password is the password corresponding to the user name
		// If useMultipleDatabases, must be empty and provide it via multipleDatabasesConfig instead
		Password string `yaml:"password"`
		// PluginName is the name of SQL plugin
		PluginName string `yaml:"pluginName" validate:"nonzero"`
		// DatabaseName is the name of SQL database to connect to
		// If useMultipleDatabases, must be empty and provide it via multipleDatabasesConfig instead
		// Required if not useMultipleDatabases
		DatabaseName string `yaml:"databaseName"`
		// ConnectAddr is the remote addr of the database
		// If useMultipleDatabases, must be empty and provide it via multipleDatabasesConfig instead
		// Required if not useMultipleDatabases
		ConnectAddr string `yaml:"connectAddr"`
		// ConnectProtocol is the protocol that goes with the ConnectAddr ex - tcp, unix
		ConnectProtocol string `yaml:"connectProtocol" validate:"nonzero"`
		// ConnectAttributes is a set of key-value attributes to be sent as part of connect data_source_name url
		ConnectAttributes map[string]string `yaml:"connectAttributes"`
		// MaxConns the max number of connections to this datastore
		MaxConns int `yaml:"maxConns"`
		// MaxIdleConns is the max number of idle connections to this datastore
		MaxIdleConns int `yaml:"maxIdleConns"`
		// MaxConnLifetime is the maximum time a connection can be alive
		MaxConnLifetime time.Duration `yaml:"maxConnLifetime"`
		// NumShards is the number of DB shards in a sharded sql database. Default is 1 for single SQL database setup.
		// It's for computing a shardID value of [0,NumShards) to decide which shard of DB to query.
		// Relationship with NumHistoryShards, both values cannot be changed once set in the same cluster,
		// and the historyShardID value calculated from NumHistoryShards will be calculated using this NumShards to get a dbShardID
		NumShards int `yaml:"nShards"`
		// EncodingType is the configuration for the type of encoding used for sql blobs
		EncodingType string `yaml:"encodingType"`
		// DecodingTypes is the configuration for all the sql blob decoding types which need to be supported
		// DecodingTypes should not be removed unless there are no blobs in database with the encoding type
		DecodingTypes []string `yaml:"decodingTypes"`
		// UseMultipleDatabases enables using multiple databases as a sharding SQL database, default is false
		// When enabled, connection will be established using MultipleDatabasesConfig in favor of single values
		// of  User, Password, DatabaseName, ConnectAddr.
		UseMultipleDatabases bool `yaml:"useMultipleDatabases"`
		// Required when UseMultipleDatabases is true
		// the length of the list should be exactly the same as NumShards
		MultipleDatabasesConfig []MultipleDatabasesConfigEntry `yaml:"multipleDatabasesConfig"`
	}

	// MultipleDatabasesConfigEntry is an entry for MultipleDatabasesConfig to connect to a single SQL database
	MultipleDatabasesConfigEntry struct {
		// User is the username to be used for the conn
		User string `yaml:"user"`
		// Password is the password corresponding to the user name
		Password string `yaml:"password"`
		// DatabaseName is the name of SQL database to connect to
		DatabaseName string `yaml:"databaseName" validate:"nonzero"`
		// ConnectAddr is the remote addr of the database
		ConnectAddr string `yaml:"connectAddr" validate:"nonzero"`
	}

	// CustomDatastoreConfig is the configuration for connecting to a custom datastore that is not supported by cadence core
	// TODO can we remove it?
	CustomDatastoreConfig struct {
		// Name of the custom datastore
		Name string `yaml:"name"`
		// Options is a set of key-value attributes that can be used by AbstractDatastoreFactory implementation
		Options map[string]string `yaml:"options"`
	}

	// Replicator describes the configuration of replicator
	// TODO can we remove it?
	Replicator struct{}

	// Logger contains the config items for logger
	Logger struct {
		// Stdout is true then the output needs to goto standard out
		// By default this is false and output will go to standard error
		Stdout bool `yaml:"stdout"`
		// Level is the desired log level
		Level string `yaml:"level"`
		// OutputFile is the path to the log output file
		// Stdout must be false, otherwise Stdout will take precedence
		OutputFile string `yaml:"outputFile"`
		// LevelKey is the desired log level, defaults to "level"
		LevelKey string `yaml:"levelKey"`
		// Encoding decides the format, supports "console" and "json".
		// "json" will print the log in JSON format(better for machine), while "console" will print in plain-text format(more human friendly)
		// Default is "json"
		Encoding string `yaml:"encoding"`
	}

	// ClusterRedirectionPolicy contains the frontend datacenter redirection policy
	// When using XDC (global domain) feature to failover a domain from one cluster to another one, client may call the passive cluster to start /signal workflows etc.
	// To have a seamless failover experience, cluster should use this forwarding option to forward those APIs to the active cluster.
	ClusterRedirectionPolicy struct {
		// Support "noop", "selected-apis-forwarding" and "all-domain-apis-forwarding", default (when empty) is "noop"
		//
		// 1) "noop" will not do any forwarding.
		//
		// 2) "all-domain-apis-forwarding" will forward all domain specific APIs(worker and non worker) if the current active domain is
		// the same as "allDomainApisForwardingTargetCluster"( or "allDomainApisForwardingTargetCluster" is empty), otherwise it fallbacks to "selected-apis-forwarding".
		//
		// 3) "selected-apis-forwarding" will forward all non-worker APIs including
		// 1. StartWorkflowExecution
		// 2. SignalWithStartWorkflowExecution
		// 3. SignalWorkflowExecution
		// 4. RequestCancelWorkflowExecution
		// 5. TerminateWorkflowExecution
		// 6. QueryWorkflow
		// 7. ResetWorkflow
		//
		// 4) "selected-apis-forwarding-v2" will forward all of "selected-apis-forwarding", and also activity responses
		// and heartbeats, but not other worker APIs.
		//
		// "selected-apis-forwarding(-v2)" and "all-domain-apis-forwarding" can work with EnableDomainNotActiveAutoForwarding dynamicconfig to select certain domains using the policy.
		//
		// Usage recommendation: when enabling XDC(global domain) feature, either "all-domain-apis-forwarding" or "selected-apis-forwarding(-v2)" should be used to ensure seamless domain failover(high availability)
		// Depending on the cost of cross cluster calls:
		//
		// 1) If the network communication overhead is high(e.g., clusters are in remote datacenters of different region), then should use "selected-apis-forwarding(-v2)".
		// But you must ensure a different set of workers with the same workflow & activity code are connected to each Cadence cluster.
		//
		// 2) If the network communication overhead is low (e.g. in the same datacenter, mostly for cluster migration usage), then you can use "all-domain-apis-forwarding". Then only one set of
		// workflow & activity worker connected of one of the Cadence cluster is enough as all domain APIs are forwarded. See more details in documentation of cluster migration section.
		// Usually "allDomainApisForwardingTargetCluster" should be empty(default value) except for very rare cases: you have more than two clusters and some are in a remote region but some are in local region.
		Policy string `yaml:"policy"`
		// A supplement for "all-domain-apis-forwarding" policy. It decides how the policy fallback to  "selected-apis-forwarding" policy.
		// If this is not empty, and current domain is not active in the value of allDomainApisForwardingTargetCluster, then the policy will fallback to "selected-apis-forwarding" policy.
		// Default is empty, meaning that all requests will not fallback.
		AllDomainApisForwardingTargetCluster string `yaml:"allDomainApisForwardingTargetCluster"`
		// Not being used, but we have to keep it so that config loading is not broken
		ToDC string `yaml:"toDC"`
	}

	// Statsd contains the config items for statsd metrics reporter
	Statsd struct {
		// The host and port of the statsd server
		HostPort string `yaml:"hostPort" validate:"nonzero"`
		// The prefix to use in reporting to statsd
		Prefix string `yaml:"prefix" validate:"nonzero"`
		// FlushInterval is the maximum interval for sending packets.
		// If it is not specified, it defaults to 1 second.
		FlushInterval time.Duration `yaml:"flushInterval"`
		// FlushBytes specifies the maximum udp packet size you wish to send.
		// If FlushBytes is unspecified, it defaults  to 1432 bytes, which is
		// considered safe for local traffic.
		FlushBytes int `yaml:"flushBytes"`
	}

	// Archival contains the config for archival
	Archival struct {
		// History is the config for the history archival
		History HistoryArchival `yaml:"history"`
		// Visibility is the config for visibility archival
		Visibility VisibilityArchival `yaml:"visibility"`
	}

	// HistoryArchival contains the config for history archival
	HistoryArchival struct {
		// Status is the status of history archival either: enabled, disabled, or paused
		Status string `yaml:"status"`
		// EnableRead whether history can be read from archival
		EnableRead bool `yaml:"enableRead"`
		// Provider contains the config for all history archivers
		Provider *HistoryArchiverProvider `yaml:"provider"`
	}

	// HistoryArchiverProvider contains the config for all history archivers
	HistoryArchiverProvider struct {
		Filestore *FilestoreArchiver `yaml:"filestore"`
		Gstorage  *GstorageArchiver  `yaml:"gstorage"`
		S3store   *S3Archiver        `yaml:"s3store"`
	}

	// VisibilityArchival contains the config for visibility archival
	VisibilityArchival struct {
		// Status is the status of visibility archival either: enabled, disabled, or paused
		Status string `yaml:"status"`
		// EnableRead whether visibility can be read from archival
		EnableRead bool `yaml:"enableRead"`
		// Provider contains the config for all visibility archivers
		Provider *VisibilityArchiverProvider `yaml:"provider"`
	}

	// VisibilityArchiverProvider contains the config for all visibility archivers
	VisibilityArchiverProvider struct {
		Filestore *FilestoreArchiver `yaml:"filestore"`
		S3store   *S3Archiver        `yaml:"s3store"`
		Gstorage  *GstorageArchiver  `yaml:"gstorage"`
	}

	// FilestoreArchiver contain the config for filestore archiver
	FilestoreArchiver struct {
		FileMode string `yaml:"fileMode"`
		DirMode  string `yaml:"dirMode"`
	}

	// GstorageArchiver contain the config for google storage archiver
	GstorageArchiver struct {
		CredentialsPath string `yaml:"credentialsPath"`
	}

	// S3Archiver contains the config for S3 archiver
	S3Archiver struct {
		Region           string  `yaml:"region"`
		Endpoint         *string `yaml:"endpoint"`
		S3ForcePathStyle bool    `yaml:"s3ForcePathStyle"`
	}

	// PublicClient is config for connecting to cadence frontend
	PublicClient struct {
		// HostPort is the host port to connect on. Host can be DNS name
		// Default to currentCluster's RPCAddress in ClusterInformation
		HostPort string `yaml:"hostPort"`
		// Transport is the tranport to use when communicating using the SDK client.
		// Defaults to:
		// - currentCluster's RPCTransport in ClusterInformation (if HostPort is not provided)
		// - grpc (if HostPort is provided)
		Transport string `yaml:"transport"`
		// interval to refresh DNS. Default to 10s
		RefreshInterval time.Duration `yaml:"RefreshInterval"`
	}

	// DomainDefaults is the default config for each domain
	DomainDefaults struct {
		// Archival is the default archival config for each domain
		Archival ArchivalDomainDefaults `yaml:"archival"`
	}

	// ArchivalDomainDefaults is the default archival config for each domain
	ArchivalDomainDefaults struct {
		// History is the domain default history archival config for each domain
		History HistoryArchivalDomainDefaults `yaml:"history"`
		// Visibility is the domain default visibility archival config for each domain
		Visibility VisibilityArchivalDomainDefaults `yaml:"visibility"`
	}

	// HistoryArchivalDomainDefaults is the default history archival config for each domain
	HistoryArchivalDomainDefaults struct {
		// Status is the domain default status of history archival: enabled or disabled
		Status string `yaml:"status"`
		// URI is the domain default URI for history archiver
		URI string `yaml:"URI"`
	}

	// VisibilityArchivalDomainDefaults is the default visibility archival config for each domain
	VisibilityArchivalDomainDefaults struct {
		// Status is the domain default status of visibility archival: enabled or disabled
		Status string `yaml:"status"`
		// URI is the domain default URI for visibility archiver
		URI string `yaml:"URI"`
	}
)

// String converts the config object into a string
func (c *Config) String() string {
	out, _ := json.MarshalIndent(c, "", "    ")
	return string(out)
}
