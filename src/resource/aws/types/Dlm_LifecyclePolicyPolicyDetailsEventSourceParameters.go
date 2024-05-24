package types

type Dlm_LifecyclePolicyPolicyDetailsEventSourceParameters struct {
	// The snapshot description that can trigger the policy. The description pattern is specified using a regular expression. The policy runs only if a snapshot with a description that matches the specified pattern is shared with your account.
	DescriptionRegex string `json:"descriptionRegex,omitempty" yaml:"descriptionRegex,omitempty"`

	// The type of event. Currently, only `shareSnapshot` events are supported.
	EventType string `json:"eventType,omitempty" yaml:"eventType,omitempty"`

	// The IDs of the AWS accounts that can trigger policy by sharing snapshots with your account. The policy only runs if one of the specified AWS accounts shares a snapshot with your account.
	SnapshotOwners []string `json:"snapshotOwners,omitempty" yaml:"snapshotOwners,omitempty"`
}
