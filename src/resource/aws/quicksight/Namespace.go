package quicksight

import types "DesignSphere_Server/src/resource/aws/types"

type Namespace struct {
	// AWS account ID.
	AwsAccountId string `json:"awsAccountId,omitempty" yaml:"awsAccountId,omitempty"`

	// User identity directory type. Defaults to `QUICKSIGHT`, the only current valid value.
	IdentityStore string `json:"identityStore,omitempty" yaml:"identityStore,omitempty"`

	/*
	   Name of the namespace.

	   The following arguments are optional:
	*/
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	//
	Timeouts types.Quicksight_NamespaceTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`
}
