package types

type Bigquery_ConnectionCloudSql struct {
	/*
	   Cloud SQL properties.
	   Structure is documented below.
	*/
	Credential Bigquery_ConnectionCloudSqlCredential `json:"credential,omitempty" yaml:"credential,omitempty"`

	// Database name.
	Database string `json:"database,omitempty" yaml:"database,omitempty"`

	// Cloud SQL instance ID in the form project:location:instance.
	InstanceId string `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`

	/*
	   (Output)
	   When the connection is used in the context of an operation in BigQuery, this service account will serve as the identity being used for connecting to the CloudSQL instance specified in this connection.
	*/
	ServiceAccountId string `json:"serviceAccountId,omitempty" yaml:"serviceAccountId,omitempty"`

	/*
	   Type of the Cloud SQL database.
	   Possible values are: `DATABASE_TYPE_UNSPECIFIED`, `POSTGRES`, `MYSQL`.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
