package types

type Appflow_ConnectorProfileConnectorProfileConfigConnectorProfilePropertiesRedshift struct {
	//
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`

	//
	BucketPrefix string `json:"bucketPrefix,omitempty" yaml:"bucketPrefix,omitempty"`

	// The unique ID that's assigned to an Amazon Redshift cluster.
	ClusterIdentifier string `json:"clusterIdentifier,omitempty" yaml:"clusterIdentifier,omitempty"`

	// ARN of the IAM role that permits AppFlow to access the database through Data API.
	DataApiRoleArn string `json:"dataApiRoleArn,omitempty" yaml:"dataApiRoleArn,omitempty"`

	// The name of an Amazon Redshift database.
	DatabaseName string `json:"databaseName,omitempty" yaml:"databaseName,omitempty"`

	// The JDBC URL of the Amazon Redshift cluster.
	DatabaseUrl string `json:"databaseUrl,omitempty" yaml:"databaseUrl,omitempty"`

	// ARN of the IAM role.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`
}
