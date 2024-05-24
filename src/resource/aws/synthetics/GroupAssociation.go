package synthetics

type GroupAssociation struct {
	// ARN of the canary.
	CanaryArn string `json:"canaryArn,omitempty" yaml:"canaryArn,omitempty"`

	// Name of the group that the canary will be associated with.
	GroupName string `json:"groupName,omitempty" yaml:"groupName,omitempty"`
}
