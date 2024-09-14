package types

type Cloudwatch_EventTargetRedshiftTarget struct {
	// The name of the database.
	Database string `json:"database,omitempty" yaml:"database,omitempty"`

	// The database user name.
	DbUser string `json:"dbUser,omitempty" yaml:"dbUser,omitempty"`

	// The name or ARN of the secret that enables access to the database.
	SecretsManagerArn string `json:"secretsManagerArn,omitempty" yaml:"secretsManagerArn,omitempty"`

	// The SQL statement text to run.
	Sql string `json:"sql,omitempty" yaml:"sql,omitempty"`

	// The name of the SQL statement.
	StatementName string `json:"statementName,omitempty" yaml:"statementName,omitempty"`

	// Indicates whether to send an event back to EventBridge after the SQL statement runs.
	WithEvent bool `json:"withEvent,omitempty" yaml:"withEvent,omitempty"`
}
