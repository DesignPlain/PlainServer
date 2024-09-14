package types

type Dlm_LifecyclePolicyPolicyDetails struct {
	// A list of resource types that should be targeted by the lifecycle policy. Valid values are `VOLUME` and `INSTANCE`.
	ResourceTypes []string `json:"resourceTypes,omitempty" yaml:"resourceTypes,omitempty"`

	// See the `schedule` configuration block.
	Schedules []Dlm_LifecyclePolicyPolicyDetailsSchedule `json:"schedules,omitempty" yaml:"schedules,omitempty"`

	/*
	   A map of tag keys and their values. Any resources that match the `resource_types` and are tagged with _any_ of these tags will be targeted.

	   > Note: You cannot have overlapping lifecycle policies that share the same `target_tags`. Pulumi is unable to detect this at plan time but it will fail during apply.
	*/
	TargetTags map[string]string `json:"targetTags,omitempty" yaml:"targetTags,omitempty"`

	// The actions to be performed when the event-based policy is triggered. You can specify only one action per policy. This parameter is required for event-based policies only. If you are creating a snapshot or AMI policy, omit this parameter. See the `action` configuration block.
	Action Dlm_LifecyclePolicyPolicyDetailsAction `json:"action,omitempty" yaml:"action,omitempty"`

	// The event that triggers the event-based policy. This parameter is required for event-based policies only. If you are creating a snapshot or AMI policy, omit this parameter. See the `event_source` configuration block.
	EventSource Dlm_LifecyclePolicyPolicyDetailsEventSource `json:"eventSource,omitempty" yaml:"eventSource,omitempty"`

	//
	Parameters Dlm_LifecyclePolicyPolicyDetailsParameters `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// The valid target resource types and actions a policy can manage. Specify `EBS_SNAPSHOT_MANAGEMENT` to create a lifecycle policy that manages the lifecycle of Amazon EBS snapshots. Specify `IMAGE_MANAGEMENT` to create a lifecycle policy that manages the lifecycle of EBS-backed AMIs. Specify `EVENT_BASED_POLICY` to create an event-based policy that performs specific actions when a defined event occurs in your AWS account. Default value is `EBS_SNAPSHOT_MANAGEMENT`.
	PolicyType string `json:"policyType,omitempty" yaml:"policyType,omitempty"`

	// The location of the resources to backup. If the source resources are located in an AWS Region, specify `CLOUD`. If the source resources are located on an Outpost in your account, specify `OUTPOST`. If you specify `OUTPOST`, Amazon Data Lifecycle Manager backs up all resources of the specified type with matching target tags across all of the Outposts in your account. Valid values are `CLOUD` and `OUTPOST`.
	ResourceLocations string `json:"resourceLocations,omitempty" yaml:"resourceLocations,omitempty"`
}
