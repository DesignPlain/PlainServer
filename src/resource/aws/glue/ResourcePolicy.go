package glue

type ResourcePolicy struct {
	// Indicates that you are using both methods to grant cross-account. Valid values are `TRUE` and `FALSE`. Note the provider will not perform drift detetction on this field as its not return on read.
	EnableHybrid string `json:"enableHybrid,omitempty" yaml:"enableHybrid,omitempty"`

	// The policy to be applied to the aws glue data catalog.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`
}
