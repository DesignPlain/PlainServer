package redshiftserverless

type ResourcePolicy struct {
	// The policy to create or update. For example, the following policy grants a user authorization to restore a snapshot.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`

	// The Amazon Resource Name (ARN) of the account to create or update a resource policy for.
	ResourceArn string `json:"resourceArn,omitempty" yaml:"resourceArn,omitempty"`
}
