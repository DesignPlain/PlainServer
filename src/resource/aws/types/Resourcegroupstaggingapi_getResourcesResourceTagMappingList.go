package types

type Resourcegroupstaggingapi_getResourcesResourceTagMappingList struct {
	// List of objects with information that shows whether a resource is compliant with the effective tag policy, including details on any noncompliant tag keys.
	ComplianceDetails []Resourcegroupstaggingapi_getResourcesResourceTagMappingListComplianceDetail `json:"complianceDetails,omitempty" yaml:"complianceDetails,omitempty"`

	// ARN of the resource.
	ResourceArn string `json:"resourceArn,omitempty" yaml:"resourceArn,omitempty"`

	// Map of tags assigned to the resource.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
