package dns

import types "DesignSphere_Server/src/resource/gcp/types"

type ResponsePolicy struct {
	// The description of the response policy, such as `My new response policy`.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The list of Google Kubernetes Engine clusters that can see this zone.
	   Structure is documented below.
	*/
	GkeClusters []types.Dns_ResponsePolicyGkeCluster `json:"gkeClusters,omitempty" yaml:"gkeClusters,omitempty"`

	/*
	   The list of network names specifying networks to which this policy is applied.
	   Structure is documented below.
	*/
	Networks []types.Dns_ResponsePolicyNetwork `json:"networks,omitempty" yaml:"networks,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The user assigned name for this Response Policy, such as `myresponsepolicy`.


	   - - -
	*/
	ResponsePolicyName string `json:"responsePolicyName,omitempty" yaml:"responsePolicyName,omitempty"`
}
