package types

type Glue_CatalogTableOptimizerConfiguration struct {
	// Indicates whether the table optimizer is enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// The ARN of the IAM role to use for the table optimizer.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`
}
