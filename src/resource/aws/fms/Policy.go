package fms

import types "DesignSphere_Server/src/resource/aws/types"

type Policy struct {
	// A map of resource tags, that if present will filter protections on resources based on the exclude_resource_tags.
	ResourceTags map[string]string `json:"resourceTags,omitempty" yaml:"resourceTags,omitempty"`

	// A resource type to protect. Conflicts with `resource_type_list`. See the [FMS API Reference](https://docs.aws.amazon.com/fms/2018-01-01/APIReference/API_Policy.html#fms-Type-Policy-ResourceType) for more information about supported values.
	ResourceType string `json:"resourceType,omitempty" yaml:"resourceType,omitempty"`

	// The objects to include in Security Service Policy Data. Documented below.
	SecurityServicePolicyData types.Fms_PolicySecurityServicePolicyData `json:"securityServicePolicyData,omitempty" yaml:"securityServicePolicyData,omitempty"`

	// If true, the request will also perform a clean-up process. Defaults to `true`. More information can be found here [AWS Firewall Manager delete policy](https://docs.aws.amazon.com/fms/2018-01-01/APIReference/API_DeletePolicy.html)
	DeleteAllPolicyResources bool `json:"deleteAllPolicyResources,omitempty" yaml:"deleteAllPolicyResources,omitempty"`

	// If true, Firewall Manager will automatically remove protections from resources that leave the policy scope. Defaults to `false`. More information can be found here [AWS Firewall Manager policy contents](https://docs.aws.amazon.com/fms/2018-01-01/APIReference/API_Policy.html)
	DeleteUnusedFmManagedResources bool `json:"deleteUnusedFmManagedResources,omitempty" yaml:"deleteUnusedFmManagedResources,omitempty"`

	// A boolean value, if true the tags that are specified in the `resource_tags` are not protected by this policy. If set to false and resource_tags are populated, resources that contain tags will be protected by this policy.
	ExcludeResourceTags bool `json:"excludeResourceTags,omitempty" yaml:"excludeResourceTags,omitempty"`

	// The friendly name of the AWS Firewall Manager Policy.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A boolean value, indicates if the policy should automatically applied to resources that already exist in the account.
	RemediationEnabled bool `json:"remediationEnabled,omitempty" yaml:"remediationEnabled,omitempty"`

	// The description of the AWS Network Firewall firewall policy.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// A map of lists of accounts and OU's to exclude from the policy.
	ExcludeMap types.Fms_PolicyExcludeMap `json:"excludeMap,omitempty" yaml:"excludeMap,omitempty"`

	// A map of lists of accounts and OU's to include in the policy.
	IncludeMap types.Fms_PolicyIncludeMap `json:"includeMap,omitempty" yaml:"includeMap,omitempty"`

	// A list of resource types to protect. Conflicts with `resource_type`. See the [FMS API Reference](https://docs.aws.amazon.com/fms/2018-01-01/APIReference/API_Policy.html#fms-Type-Policy-ResourceType) for more information about supported values. Lists with only one element are not supported, instead use `resource_type`.
	ResourceTypeLists []string `json:"resourceTypeLists,omitempty" yaml:"resourceTypeLists,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
