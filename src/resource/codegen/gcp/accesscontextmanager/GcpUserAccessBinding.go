package accesscontextmanager

type GcpUserAccessBinding struct {
	// Required. Access level that a user must have to be granted access. Only one access level is supported, not multiple. This repeated field must have exactly one element. Example: "accessPolicies/9522/accessLevels/device_trusted"
	AccessLevels string `json:"accessLevels,omitempty" yaml:"accessLevels,omitempty"`

	// Required. Immutable. Google Group id whose members are subject to this binding's restrictions. See "id" in the G Suite Directory API's Groups resource. If a group's email address/alias is changed, this resource will continue to point at the changed group. This field does not accept group email addresses or aliases. Example: "01d520gv4vjcrht"
	GroupKey string `json:"groupKey,omitempty" yaml:"groupKey,omitempty"`

	/*
	   Required. ID of the parent organization.


	   - - -
	*/
	OrganizationId string `json:"organizationId,omitempty" yaml:"organizationId,omitempty"`
}
