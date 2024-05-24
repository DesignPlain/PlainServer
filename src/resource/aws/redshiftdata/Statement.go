package redshiftdata

import types "DesignSphere_Server/src/resource/aws/types"

type Statement struct {
	// The database user name.
	DbUser string `json:"dbUser,omitempty" yaml:"dbUser,omitempty"`

	//
	Parameters []types.Redshiftdata_StatementParameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// The name or ARN of the secret that enables access to the database.
	SecretArn string `json:"secretArn,omitempty" yaml:"secretArn,omitempty"`

	/*
	   The SQL statement text to run.

	   The following arguments are optional:
	*/
	Sql string `json:"sql,omitempty" yaml:"sql,omitempty"`

	// A value that indicates whether to send an event to the Amazon EventBridge event bus after the SQL statement runs.
	WithEvent bool `json:"withEvent,omitempty" yaml:"withEvent,omitempty"`

	// The serverless workgroup name. This parameter is required when connecting to a serverless workgroup and authenticating using either Secrets Manager or temporary credentials.
	WorkgroupName string `json:"workgroupName,omitempty" yaml:"workgroupName,omitempty"`

	// The name of the database.
	Database string `json:"database,omitempty" yaml:"database,omitempty"`

	// The name of the SQL statement. You can name the SQL statement when you create it to identify the query.
	StatementName string `json:"statementName,omitempty" yaml:"statementName,omitempty"`

	// The cluster identifier. This parameter is required when connecting to a cluster and authenticating using either Secrets Manager or temporary credentials.
	ClusterIdentifier string `json:"clusterIdentifier,omitempty" yaml:"clusterIdentifier,omitempty"`
}
