package types

type Devicefarm_DevicePoolRule struct {
	// The rule's stringified attribute. Valid values are: `APPIUM_VERSION`, `ARN`, `AVAILABILITY`, `FLEET_TYPE`, `FORM_FACTOR`, `INSTANCE_ARN`, `INSTANCE_LABELS`, `MANUFACTURER`, `MODEL`, `OS_VERSION`, `PLATFORM`, `REMOTE_ACCESS_ENABLED`, `REMOTE_DEBUG_ENABLED`.
	Attribute string `json:"attribute,omitempty" yaml:"attribute,omitempty"`

	// Specifies how Device Farm compares the rule's attribute to the value. For the operators that are supported by each attribute. Valid values are: `EQUALS`, `NOT_IN`, `IN`, `GREATER_THAN`, `GREATER_THAN_OR_EQUALS`, `LESS_THAN`, `LESS_THAN_OR_EQUALS`, `CONTAINS`.
	Operator string `json:"operator,omitempty" yaml:"operator,omitempty"`

	// The rule's value.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
