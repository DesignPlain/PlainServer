package compute

type NetworkFirewallPolicyAssociation struct {
	// The target that the firewall policy is attached to.
	AttachmentTarget string `json:"attachmentTarget,omitempty" yaml:"attachmentTarget,omitempty"`

	// The firewall policy ID of the association.
	FirewallPolicy string `json:"firewallPolicy,omitempty" yaml:"firewallPolicy,omitempty"`

	/*
	   The name for an association.



	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The project for the resource
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
