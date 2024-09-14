package types

type Datastream_StreamSourceConfig struct {
	/*
	   MySQL data source configuration.
	   Structure is documented below.
	*/
	MysqlSourceConfig Datastream_StreamSourceConfigMysqlSourceConfig `json:"mysqlSourceConfig,omitempty" yaml:"mysqlSourceConfig,omitempty"`

	/*
	   MySQL data source configuration.
	   Structure is documented below.
	*/
	OracleSourceConfig Datastream_StreamSourceConfigOracleSourceConfig `json:"oracleSourceConfig,omitempty" yaml:"oracleSourceConfig,omitempty"`

	/*
	   PostgreSQL data source configuration.
	   Structure is documented below.
	*/
	PostgresqlSourceConfig Datastream_StreamSourceConfigPostgresqlSourceConfig `json:"postgresqlSourceConfig,omitempty" yaml:"postgresqlSourceConfig,omitempty"`

	// Source connection profile resource. Format: projects/{project}/locations/{location}/connectionProfiles/{name}
	SourceConnectionProfile string `json:"sourceConnectionProfile,omitempty" yaml:"sourceConnectionProfile,omitempty"`
}
