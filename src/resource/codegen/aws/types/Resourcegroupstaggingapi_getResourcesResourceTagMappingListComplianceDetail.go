package types

type Resourcegroupstaggingapi_getResourcesResourceTagMappingListComplianceDetail struct {
	/*
	   Whether the resource is compliant.
	   - `keys_with_noncompliant_values ` - Set of tag keys with non-compliant tag values.
	   - `non_compliant_keys ` - Set of non-compliant tag keys.
	*/
	ComplianceStatus bool `json:"complianceStatus,omitempty" yaml:"complianceStatus,omitempty"`

	//
	KeysWithNoncompliantValues []string `json:"keysWithNoncompliantValues,omitempty" yaml:"keysWithNoncompliantValues,omitempty"`

	//
	NonCompliantKeys []string `json:"nonCompliantKeys,omitempty" yaml:"nonCompliantKeys,omitempty"`
}
