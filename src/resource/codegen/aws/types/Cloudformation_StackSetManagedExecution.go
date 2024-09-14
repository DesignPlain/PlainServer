package types

type Cloudformation_StackSetManagedExecution struct {
	// When set to true, StackSets performs non-conflicting operations concurrently and queues conflicting operations. After conflicting operations finish, StackSets starts queued operations in request order. Default is false.
	Active bool `json:"active,omitempty" yaml:"active,omitempty"`
}
