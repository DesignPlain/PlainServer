package redshift

type ResourcePolicy struct {
	// The Amazon Resource Name (ARN) of the account to create or update a resource policy for.
	ResourceArn string `json:"resourceArn,omitempty" yaml:"resourceArn,omitempty"`

	// The content of the resource policy being updated.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`
}
