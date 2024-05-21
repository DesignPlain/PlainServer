package types

type Bigquery_JobQueryScriptOptions struct {
	// Limit on the number of bytes billed per statement. Exceeding this budget results in an error.
	StatementByteBudget string `json:"statementByteBudget,omitempty" yaml:"statementByteBudget,omitempty"`

	// Timeout period for each statement in a script.
	StatementTimeoutMs string `json:"statementTimeoutMs,omitempty" yaml:"statementTimeoutMs,omitempty"`

	/*
	   Determines which statement in the script represents the "key result",
	   used to populate the schema and query results of the script job.
	   Possible values are: `LAST`, `FIRST_SELECT`.
	*/
	KeyResultStatement string `json:"keyResultStatement,omitempty" yaml:"keyResultStatement,omitempty"`
}
