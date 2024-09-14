package types

type Storage_getBucketLifecycleRule struct {
	// The Lifecycle Rule's action configuration. A single block of this type is supported.
	Actions []Storage_getBucketLifecycleRuleAction `json:"actions,omitempty" yaml:"actions,omitempty"`

	// The Lifecycle Rule's condition configuration.
	Conditions []Storage_getBucketLifecycleRuleCondition `json:"conditions,omitempty" yaml:"conditions,omitempty"`
}
