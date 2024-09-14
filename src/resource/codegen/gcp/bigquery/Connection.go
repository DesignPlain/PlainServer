package bigquery

import types "libds/gcp/types"

type Connection struct {
	/*
	   Connection properties specific to Amazon Web Services.
	   Structure is documented below.
	*/
	Aws types.Bigquery_ConnectionAws `json:"aws,omitempty" yaml:"aws,omitempty"`

	/*
	   Container for connection properties for delegation of access to GCP resources.
	   Structure is documented below.
	*/
	CloudResource types.Bigquery_ConnectionCloudResource `json:"cloudResource,omitempty" yaml:"cloudResource,omitempty"`

	/*
	   Connection properties specific to Cloud Spanner
	   Structure is documented below.
	*/
	CloudSpanner types.Bigquery_ConnectionCloudSpanner `json:"cloudSpanner,omitempty" yaml:"cloudSpanner,omitempty"`

	/*
	   Connection properties specific to the Cloud SQL.
	   Structure is documented below.
	*/
	CloudSql types.Bigquery_ConnectionCloudSql `json:"cloudSql,omitempty" yaml:"cloudSql,omitempty"`

	// A descriptive name for the connection
	FriendlyName string `json:"friendlyName,omitempty" yaml:"friendlyName,omitempty"`

	/*
	   The geographic location where the connection should reside.
	   Cloud SQL instance must be in the same location as the connection
	   with following exceptions: Cloud SQL us-central1 maps to BigQuery US, Cloud SQL europe-west1 maps to BigQuery EU.
	   Examples: US, EU, asia-northeast1, us-central1, europe-west1.
	   Spanner Connections same as spanner region
	   AWS allowed regions are aws-us-east-1
	   Azure allowed regions are azure-eastus2
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   Container for connection properties specific to Azure.
	   Structure is documented below.
	*/
	Azure types.Bigquery_ConnectionAzure `json:"azure,omitempty" yaml:"azure,omitempty"`

	// Optional connection id that should be assigned to the created connection.
	ConnectionId string `json:"connectionId,omitempty" yaml:"connectionId,omitempty"`

	// A descriptive description for the connection
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Container for connection properties to execute stored procedures for Apache Spark. resources.
	   Structure is documented below.
	*/
	Spark types.Bigquery_ConnectionSpark `json:"spark,omitempty" yaml:"spark,omitempty"`
}
