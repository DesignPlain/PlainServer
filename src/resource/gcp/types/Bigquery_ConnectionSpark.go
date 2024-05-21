package types

type Bigquery_ConnectionSpark struct {
	/*
	   Dataproc Metastore Service configuration for the connection.
	   Structure is documented below.
	*/
	MetastoreServiceConfig Bigquery_ConnectionSparkMetastoreServiceConfig `json:"metastoreServiceConfig,omitempty" yaml:"metastoreServiceConfig,omitempty"`

	/*
	   (Output)
	   The account ID of the service created for the purpose of this connection.
	*/
	ServiceAccountId string `json:"serviceAccountId,omitempty" yaml:"serviceAccountId,omitempty"`

	/*
	   Spark History Server configuration for the connection.
	   Structure is documented below.
	*/
	SparkHistoryServerConfig Bigquery_ConnectionSparkSparkHistoryServerConfig `json:"sparkHistoryServerConfig,omitempty" yaml:"sparkHistoryServerConfig,omitempty"`
}
