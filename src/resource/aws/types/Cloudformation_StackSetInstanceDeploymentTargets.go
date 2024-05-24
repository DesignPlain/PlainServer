package types

type Cloudformation_StackSetInstanceDeploymentTargets struct {
	// The organization root ID or organizational unit (OU) IDs to which StackSets deploys.
	OrganizationalUnitIds []string `json:"organizationalUnitIds,omitempty" yaml:"organizationalUnitIds,omitempty"`
}
