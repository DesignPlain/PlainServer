package types

type Quicksight_DataSourceParameters struct {
	// Parameters for connecting to Aurora Postgresql.
	AuroraPostgresql Quicksight_DataSourceParametersAuroraPostgresql `json:"auroraPostgresql,omitempty" yaml:"auroraPostgresql,omitempty"`

	// Parameters for connecting to MariaDB.
	MariaDb Quicksight_DataSourceParametersMariaDb `json:"mariaDb,omitempty" yaml:"mariaDb,omitempty"`

	// Parameters for connecting to RDS.
	Rds Quicksight_DataSourceParametersRds `json:"rds,omitempty" yaml:"rds,omitempty"`

	// Parameters for connecting to ServiceNow.
	ServiceNow Quicksight_DataSourceParametersServiceNow `json:"serviceNow,omitempty" yaml:"serviceNow,omitempty"`

	// Parameters for connecting to SQL Server.
	SqlServer Quicksight_DataSourceParametersSqlServer `json:"sqlServer,omitempty" yaml:"sqlServer,omitempty"`

	// Parameters for connecting to Amazon Elasticsearch.
	AmazonElasticsearch Quicksight_DataSourceParametersAmazonElasticsearch `json:"amazonElasticsearch,omitempty" yaml:"amazonElasticsearch,omitempty"`

	// Parameters for connecting to Athena.
	Athena Quicksight_DataSourceParametersAthena `json:"athena,omitempty" yaml:"athena,omitempty"`

	// Parameters for connecting to AWS IOT Analytics.
	AwsIotAnalytics Quicksight_DataSourceParametersAwsIotAnalytics `json:"awsIotAnalytics,omitempty" yaml:"awsIotAnalytics,omitempty"`

	// Parameters for connecting to Aurora MySQL.
	Aurora Quicksight_DataSourceParametersAurora `json:"aurora,omitempty" yaml:"aurora,omitempty"`

	// Parameters for connecting to Postgresql.
	Postgresql Quicksight_DataSourceParametersPostgresql `json:"postgresql,omitempty" yaml:"postgresql,omitempty"`

	// Parameters for connecting to Redshift.
	Redshift Quicksight_DataSourceParametersRedshift `json:"redshift,omitempty" yaml:"redshift,omitempty"`

	// Parameters for connecting to S3.
	S3 Quicksight_DataSourceParametersS3 `json:"s3,omitempty" yaml:"s3,omitempty"`

	// Parameters for connecting to Snowflake.
	Snowflake Quicksight_DataSourceParametersSnowflake `json:"snowflake,omitempty" yaml:"snowflake,omitempty"`

	// Parameters for connecting to Spark.
	Spark Quicksight_DataSourceParametersSpark `json:"spark,omitempty" yaml:"spark,omitempty"`

	// Parameters for connecting to Twitter.
	Twitter Quicksight_DataSourceParametersTwitter `json:"twitter,omitempty" yaml:"twitter,omitempty"`

	// Parameters for connecting to Jira.
	Jira Quicksight_DataSourceParametersJira `json:"jira,omitempty" yaml:"jira,omitempty"`

	// Parameters for connecting to MySQL.
	Mysql Quicksight_DataSourceParametersMysql `json:"mysql,omitempty" yaml:"mysql,omitempty"`

	// Parameters for connecting to Oracle.
	Oracle Quicksight_DataSourceParametersOracle `json:"oracle,omitempty" yaml:"oracle,omitempty"`

	// Parameters for connecting to Presto.
	Presto Quicksight_DataSourceParametersPresto `json:"presto,omitempty" yaml:"presto,omitempty"`

	// Parameters for connecting to Teradata.
	Teradata Quicksight_DataSourceParametersTeradata `json:"teradata,omitempty" yaml:"teradata,omitempty"`
}
