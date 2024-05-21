package types



type Storage_BucketLifecycleRule struct {
	// The Lifecycle Rule's action configuration. A single block of this type is supported. Structure is documented below.
	Action Storage_BucketLifecycleRuleAction `json:"action,omitempty" yaml:"action,omitempty"`

	// The Lifecycle Rule's condition configuration. A single block of this type is supported. Structure is documented below.
	Condition Storage_BucketLifecycleRuleCondition `json:"condition,omitempty" yaml:"condition,omitempty"`
}
