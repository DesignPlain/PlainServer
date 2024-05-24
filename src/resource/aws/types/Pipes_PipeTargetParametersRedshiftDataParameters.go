package types

type Pipes_PipeTargetParametersRedshiftDataParameters struct {
	// The name of the database. Required when authenticating using temporary credentials.
	Database string `json:"database,omitempty" yaml:"database,omitempty"`

	// The database user name. Required when authenticating using temporary credentials.
	DbUser string `json:"dbUser,omitempty" yaml:"dbUser,omitempty"`

	// The name or ARN of the secret that enables access to the database. Required when authenticating using Secrets Manager.
	SecretManagerArn string `json:"secretManagerArn,omitempty" yaml:"secretManagerArn,omitempty"`

	// List of SQL statements text to run, each of maximum length of 100,000.
	Sqls []string `json:"sqls,omitempty" yaml:"sqls,omitempty"`

	// The name of the SQL statement. You can name the SQL statement when you create it to identify the query.
	StatementName string `json:"statementName,omitempty" yaml:"statementName,omitempty"`

	// Indicates whether to send an event back to EventBridge after the SQL statement runs.
	WithEvent bool `json:"withEvent,omitempty" yaml:"withEvent,omitempty"`
}
